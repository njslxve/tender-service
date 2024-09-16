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

type RollbackTenderInterface interface {
	RollbackTender(tenderId string, version string, username string) (dto.TenderResponse, error)
}

func RollbackTender(logger *slog.Logger, ucase RollbackTenderInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handler.RollbackTender"

		tenderId := chi.URLParam(r, "tenderId")
		version := chi.URLParam(r, "version")
		username := r.URL.Query().Get("username")

		validData := make(map[string]string)
		validData["tenderId"] = tenderId

		err := validate.ValidateParams(validData)
		if err != nil {
			e := dto.Error{
				Reason: ErrBadRequest,
			}

			logger.Error(op, slog.String("error", err.Error()))

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(e)

			return
		}

		tender, err := ucase.RollbackTender(tenderId, version, username)
		if err != nil {
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

			if errors.Is(err, usecase.ErrTenderNotFound) {
				e := dto.Error{
					Reason: ErrTenderNotFound,
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
			json.NewEncoder(w).Encode(tender)
		}
	}
}
