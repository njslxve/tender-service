package handler

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725721237-team-77964/zadanie-6105/internal/dto"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725721237-team-77964/zadanie-6105/internal/usecase"
)

type UserBidsInterface interface {
	GetUserBids(username string, limit string, offset string) ([]dto.BidResponse, error)
}

func GetUserBids(logger *slog.Logger, ucase UserBidsInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handler.GetUserBids"

		limit := r.URL.Query().Get("limit")
		offset := r.URL.Query().Get("offset")
		username := r.URL.Query().Get("username")

		bids, err := ucase.GetUserBids(username, limit, offset)
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
			json.NewEncoder(w).Encode(bids)
		}
	}
}
