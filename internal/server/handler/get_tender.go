package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/njslxve/tender-service/internal/dto"
	"github.com/njslxve/tender-service/internal/validate"
)

type GetTendersInterface interface {
	GetTenders(serviceType string, limit string, offset string) ([]dto.TenderResponse, error)
}

func GetTenders(logger *slog.Logger, ucase GetTendersInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handler.GetTenders"

		serviceType := r.URL.Query().Get("service_type")
		limit := r.URL.Query().Get("limit")
		offset := r.URL.Query().Get("offset")

		validData := make(map[string]string)
		if serviceType != "" {
			validData["serviceType"] = serviceType
		}

		err := validate.ValidateParams(validData)
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

		tenders, err := ucase.GetTenders(serviceType, limit, offset)
		if err != nil {
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

			return
		} else {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(tenders)
		}
	}
}
