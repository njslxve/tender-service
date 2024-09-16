package handler

import (
	"log/slog"
	"net/http"

	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725721237-team-77964/zadanie-6105/internal/usecase"
)

func Ping(logger *slog.Logger, ucase *usecase.Usecase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		_, _ = w.Write([]byte("ok"))
	}
}
