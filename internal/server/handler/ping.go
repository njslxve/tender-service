package handler

import (
	"log/slog"
	"net/http"

	"github.com/njslxve/tender-service/internal/usecase"
)

func Ping(logger *slog.Logger, ucase *usecase.Usecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		_, _ = w.Write([]byte("ok"))
	}
}
