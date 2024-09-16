package handler

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/njslxve/tender-service/internal/dto"
	"github.com/njslxve/tender-service/internal/usecase"
)

type GetUserTendersInterface interface {
	GetUserTenders(username string, limit string, offset string) ([]dto.TenderResponse, error)
}

func GetUserTenders(logger *slog.Logger, ucase GetUserTendersInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handler.GetTenders"

		limit := r.URL.Query().Get("limit")
		offset := r.URL.Query().Get("offset")
		username := r.URL.Query().Get("username")

		tenders, err := ucase.GetUserTenders(username, limit, offset)
		if err != nil {
			if errors.Is(err, usecase.ErrUserNotFound) {
				e := dto.Error{
					Reason: ErrUserNotFound,
				}

				logger.Error("User not found",
					slog.String("operation", op),
					slog.String("error", err.Error()),
				)

				w.Header().Add("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(e)

				return
			}

			e := dto.Error{
				Reason: ErrBadRequest,
			}

			logger.Error(op, slog.String("error", err.Error()))

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(e)

			return
		} else {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(tenders)
		}
	}
}
