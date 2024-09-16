package handler

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725721237-team-77964/zadanie-6105/internal/dto"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725721237-team-77964/zadanie-6105/internal/usecase"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725721237-team-77964/zadanie-6105/internal/validate"
)

type CreateTenderInterface interface {
	CreateTender(req dto.TenderRequest) (dto.TenderResponse, error)
}

func CreateTender(logger *slog.Logger, ucase CreateTenderInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handler.CreateTender"

		var req dto.TenderRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			e := dto.Error{
				Reason: ErrBadRequest,
			}

			logger.Error("Bad request",
				slog.String("operation", op),
				slog.String("error", err.Error()),
			)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(e)

			return
		}

		err := validate.ValidateTender(req)
		if err != nil {
			e := dto.Error{
				Reason: ErrBadRequest,
			}

			logger.Error("Bad request",
				slog.String("operation", op),
				slog.String("error", err.Error()),
			)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(e)

			return
		}

		tender, err := ucase.CreateTender(req)
		if err != nil {
			if errors.Is(err, usecase.ErrNotPermissions) {
				e := dto.Error{
					Reason: ErrNotPermissions,
				}

				logger.Error("Not permissions",
					slog.String("operation", op),
					slog.String("error", err.Error()),
				)

				w.Header().Add("Content-Type", "application/json")
				w.WriteHeader(http.StatusForbidden)
				json.NewEncoder(w).Encode(e)

				return
			}

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
				Reason: ErrInternal,
			}

			logger.Error("Server error",
				slog.String("operation", op),
				slog.String("error", err.Error()),
			)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(e)
		} else {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(tender)
		}
	}
}
