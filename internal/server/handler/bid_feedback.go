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

func SubmitBidFeedback(logger *slog.Logger, ucase *usecase.Usecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handler.SubmitBidFeedback"

		bidID := chi.URLParam(r, "bidId")
		feedback := r.URL.Query().Get("bidFeedback")
		username := r.URL.Query().Get("username")

		validData := make(map[string]string)
		validData["bidID"] = bidID
		validData["feedback"] = feedback

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

		bid, err := ucase.SubmitBidFeedback(bidID, feedback, username)
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
