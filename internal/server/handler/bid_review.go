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

func GetBidReviews(logger *slog.Logger, ucase *usecase.Usecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handler.GetBidReviews"

		tenderID := chi.URLParam(r, "tenderId")
		author := r.URL.Query().Get("authorUsername")
		requester := r.URL.Query().Get("requesterUsername")
		limit := r.URL.Query().Get("limit")
		offset := r.URL.Query().Get("offset")

		validData := make(map[string]string)
		validData["tenderID"] = tenderID

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

		review, err := ucase.GetBidReviews(tenderID, author, requester, limit, offset)
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
			json.NewEncoder(w).Encode(review)
		}
	}
}
