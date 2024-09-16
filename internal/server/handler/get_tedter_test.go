package handler_test

import (
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725721237-team-77964/zadanie-6105/internal/dto"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725721237-team-77964/zadanie-6105/internal/server/handler"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725721237-team-77964/zadanie-6105/pkg/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetTenders(t *testing.T) {
	logger := slog.Default()
	mockucase := new(mocks.UsecaseMock)

	h := handler.GetTenders(logger, mockucase)

	mockucase.On("GetTenders", mock.Anything, mock.Anything, mock.Anything).Return([]dto.TenderResponse{}, nil)

	httpReq := httptest.NewRequest(http.MethodGet, "/api/tenders", nil)
	rr := httptest.NewRecorder()

	h(rr, httpReq)

	res := rr.Result()

	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestGetTendersBadRequest(t *testing.T) {
	logger := slog.Default()
	mockucase := new(mocks.UsecaseMock)

	h := handler.GetTenders(logger, mockucase)

	httpReq := httptest.NewRequest("GET", "/api/tenders?service_type=aboba", nil)
	rr := httptest.NewRecorder()

	h(rr, httpReq)

	res := rr.Result()

	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
}
