package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725721237-team-77964/zadanie-6105/internal/dto"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725721237-team-77964/zadanie-6105/internal/usecase"
	"github.com/go-chi/chi/v5"
)

func GetBidsForTender(logger *slog.Logger, ucase *usecase.Usecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handler.GetBidsForTender"

		tendterID := chi.URLParam(r, "tenderId")
		username := r.URL.Query().Get("username")
		limit := r.URL.Query().Get("limit")
		offset := r.URL.Query().Get("offset")

		// err := validate.Params(tendtedID, username)
		// if err != nil {
		// 	e := dto.Error{
		// 		Reason: err.Error(), // TODO: add error message
		// 	}

		// 	logger.Error(op, slog.String("error", err.Error()))

		// 	w.Header().Add("Content-Type", "application/json")
		// 	w.WriteHeader(http.StatusBadRequest)
		// 	json.NewEncoder(w).Encode(e)

		// 	return
		// }

		bids, err := ucase.GetBidsForTender(tendterID, username, limit, offset)
		if err != nil {
			e := dto.Error{
				Reason: err.Error(), // TODO: add error message
			}

			logger.Error(op, slog.String("error", err.Error()))

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(e)
		} else {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(bids)
		}
	}
}
