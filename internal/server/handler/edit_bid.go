package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725721237-team-77964/zadanie-6105/internal/dto"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725721237-team-77964/zadanie-6105/internal/usecase"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725721237-team-77964/zadanie-6105/internal/validate"
	"github.com/go-chi/chi/v5"
)

func EditBid(logger *slog.Logger, ucase *usecase.Usecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handler.EditBid"

		bidID := chi.URLParam(r, "bidId")
		username := r.URL.Query().Get("username")

		validData := make(map[string]string)
		validData["bidID"] = bidID

		err := validate.ValidateParams(validData)
		if err != nil {
			e := dto.Error{
				Reason: err.Error(), // TODO: add error message
			}

			logger.Error(op, slog.String("error", err.Error()))

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(e)

			return
		}

		var req dto.BidRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			e := dto.Error{
				Reason: "Bad request", // TODO: add error message
			}

			logger.Error(op, slog.String("error", err.Error()))

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(e)

			return
		}

		bid, err := ucase.EditBid(bidID, username, req)
		if err != nil {
			e := dto.Error{
				Reason: err.Error(), // TODO: add error message
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
