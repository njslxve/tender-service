package server

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/njslxve/tender-service/internal/config"
	"github.com/njslxve/tender-service/internal/server/handler"
	"github.com/njslxve/tender-service/internal/usecase"
)

type Server struct {
	cfg    *config.Config
	logger *slog.Logger
	ucase  *usecase.Usecase
}

func New(cfg *config.Config, logger *slog.Logger, ucase *usecase.Usecase) *Server {
	return &Server{
		cfg:    cfg,
		logger: logger,
		ucase:  ucase,
	}
}

func (s *Server) Start() {
	r := chi.NewRouter()

	r.Use(middleware.Timeout(30 * time.Second))

	r.Route("/api", func(r chi.Router) {
		r.Get("/ping", handler.Ping(s.logger, s.ucase))

		r.Route("/tenders", func(r chi.Router) {
			r.Get("/", handler.GetTenders(s.logger, s.ucase))
			r.Post("/new", handler.CreateTender(s.logger, s.ucase))
			r.Get("/my", handler.GetUserTenders(s.logger, s.ucase))

			r.Route("/{tenderId}", func(r chi.Router) {
				r.Get("/status", handler.GetTenderStatus(s.logger, s.ucase))
				r.Put("/status", handler.UpdateTenderStatus(s.logger, s.ucase))
				r.Patch("/edit", handler.EditTender(s.logger, s.ucase))
				r.Put("/rollback/{version}", handler.RollbackTender(s.logger, s.ucase))
			})
		})

		r.Route("/bids", func(r chi.Router) {
			r.Post("/new", handler.CreateBid(s.logger, s.ucase))
			r.Get("/my", handler.GetUserBids(s.logger, s.ucase))

			r.Get("/{tenderId}/list", handler.GetBidsForTender(s.logger, s.ucase))
			r.Get("/{tenderId}/reviews", handler.GetBidReviews(s.logger, s.ucase))

			r.Route("/{bidId}", func(r chi.Router) {
				r.Get("/status", handler.GetBidStatus(s.logger, s.ucase))
				r.Put("/status", handler.UpdateBidStatus(s.logger, s.ucase))
				r.Patch("/edit", handler.EditBid(s.logger, s.ucase))
				r.Put("/submit_decision", handler.SubmitBidDecision(s.logger, s.ucase))
				r.Put("/feedback", handler.SubmitBidFeedback(s.logger, s.ucase))
				r.Put("/rollback/{version}", handler.RollbackBid(s.logger, s.ucase))
			})
		})
	})

	s.logger.Info("Starting server",
		slog.String("address", s.cfg.ServerAddr),
	)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	srv := &http.Server{
		Addr:         s.cfg.ServerAddr,
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.logger.Debug("server error",
				slog.String("error", err.Error()),
			)
		}
	}()

	s.logger.Info("server started")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	<-done

	s.logger.Info("server shutting down")

	if err := srv.Shutdown(ctx); err != nil {
		s.logger.Error("failed to shutdown server")
	}

	s.logger.Info("server stopped")
}
