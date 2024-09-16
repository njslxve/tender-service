package handler

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/njslxve/tender-service/internal/dto"
	"github.com/njslxve/tender-service/internal/usecase"
	"github.com/njslxve/tender-service/internal/validate"
)

type RollbackBidInterface interface {
	RollbackBid(bidID, username, version string) (dto.BidResponse, error)
}

func RollbackBid(logger *slog.Logger, ucase RollbackBidInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handler.RollbackBid"

		bidID := chi.URLParam(r, "bidId")
		username := r.URL.Query().Get("username")
		version := r.URL.Query().Get("version")

		validData := make(map[string]string)
		validData["bidId"] = bidID

		err := validate.ValidateParams(validData)
		if err != nil {
			e := dto.Error{
				Reason: ErrInternal,
			}

			logger.Error(op, slog.String("error", err.Error()))

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(e)

			return
		}

		bid, err := ucase.RollbackBid(bidID, version, username)
		if err != nil {
			if errors.Is(err, usecase.ErrBidNotFound) {
				e := dto.Error{
					Reason: ErrBidNotFound,
				}

				logger.Error(op, slog.String("error", err.Error()))

				w.Header().Add("Content-Type", "application/json")
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode(e)

				return
			}

			if errors.Is(err, usecase.ErrNotPermissions) {
				e := dto.Error{
					Reason: ErrNotPermissions,
				}

				logger.Error(op, slog.String("error", err.Error()))

				w.Header().Add("Content-Type", "application/json")
				w.WriteHeader(http.StatusForbidden)
				json.NewEncoder(w).Encode(e)

				return
			}

			if errors.Is(err, usecase.ErrUserNotFound) {
				e := dto.Error{
					Reason: ErrUserNotFound,
				}

				logger.Error(op, slog.String("error", err.Error()))

				w.Header().Add("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(e)

				return
			}

			e := dto.Error{
				Reason: ErrInternal,
			}

			logger.Error(op, slog.String("error", err.Error()))

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(e)

			return
		} else {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(bid)
		}
	}
}
