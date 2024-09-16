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

func GetBidsForTender(logger *slog.Logger, ucase *usecase.Usecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handler.GetBidsForTender"

		tendterID := chi.URLParam(r, "tenderId")
		username := r.URL.Query().Get("username")
		limit := r.URL.Query().Get("limit")
		offset := r.URL.Query().Get("offset")

		validData := make(map[string]string)
		validData["tenderID"] = tendterID

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

		bids, err := ucase.GetBidsForTender(tendterID, username, limit, offset)
		if err != nil {
			e := dto.Error{
				Reason: ErrInternal,
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
