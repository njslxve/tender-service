package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/njslxve/tender-service/internal/dto"
	"github.com/njslxve/tender-service/internal/usecase"
	"github.com/njslxve/tender-service/internal/validate"
)

func SubmitBidDecision(logger *slog.Logger, ucase *usecase.Usecase) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handler.SubmitBidDecision"

		bidID := chi.URLParam(r, "bidId")
		username := r.URL.Query().Get("username")
		decision := r.URL.Query().Get("decision")

		validData := make(map[string]string)
		validData["bidID"] = bidID
		validData["decision"] = decision

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

		bid, err := ucase.SubmitBidDecision(bidID, username, decision)
		if err != nil {
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
