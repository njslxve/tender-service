package handler_test

import (
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725721237-team-77964/zadanie-6105/internal/dto"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725721237-team-77964/zadanie-6105/internal/server/handler"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725721237-team-77964/zadanie-6105/internal/usecase"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725721237-team-77964/zadanie-6105/pkg/mocks"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRollbackBid(t *testing.T) {
	logger := slog.Default()
	mockucase := new(mocks.UsecaseMock)

	h := handler.RollbackBid(logger, mockucase)

	r := chi.NewRouter()
	r.Put("/api/bids/{bidId}/rollback/{version}", h)

	mockucase.On("RollbackBid", mock.Anything, mock.Anything, mock.Anything).Return(dto.BidResponse{}, nil)

	httpReq := httptest.NewRequest(http.MethodPut, "/api/bids/d976cd81-3c1f-4d75-9841-1003af7d1e40/rollback/1?username=test_user", nil)
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, httpReq)

	res := rr.Result()

	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestRollbackBidValidationError(t *testing.T) {
	logger := slog.Default()
	mockucase := new(mocks.UsecaseMock)

	h := handler.RollbackBid(logger, mockucase)

	r := chi.NewRouter()
	r.Put("/api/bids/{bidId}/rollback/{version}", h)

	httpReq := httptest.NewRequest(http.MethodPut, "/api/bids/d976cd81/rollback/1?username=test_user", nil)

	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, httpReq)

	res := rr.Result()

	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func TestRollbackBidUnauthorized(t *testing.T) {
	logger := slog.Default()
	mockucase := new(mocks.UsecaseMock)

	h := handler.RollbackBid(logger, mockucase)

	r := chi.NewRouter()
	r.Put("/api/bids/{bidId}/rollback/{version}", h)

	mockucase.On("RollbackBid", mock.Anything, mock.Anything, mock.Anything).Return(dto.BidResponse{}, usecase.ErrNotPermissions)

	httpReq := httptest.NewRequest(http.MethodPut, "/api/bids/d976cd81-3c1f-4d75-9841-1003af7d1e40/rollback/1?username=test_user", nil)

	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, httpReq)

	res := rr.Result()

	assert.Equal(t, http.StatusForbidden, res.StatusCode)
}

func TestRollbackBidForbidden(t *testing.T) {
	logger := slog.Default()
	mockucase := new(mocks.UsecaseMock)

	h := handler.RollbackBid(logger, mockucase)

	r := chi.NewRouter()
	r.Put("/api/bids/{bidId}/rollback/{version}", h)

	mockucase.On("RollbackBid", mock.Anything, mock.Anything, mock.Anything).Return(dto.BidResponse{}, usecase.ErrNotPermissions)

	httpReq := httptest.NewRequest(http.MethodPut, "/api/bids/d976cd81-3c1f-4d75-9841-1003af7d1e40/rollback/1?username=test_user", nil)

	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, httpReq)

	res := rr.Result()

	assert.Equal(t, http.StatusForbidden, res.StatusCode)
}

func TestRollbackBidNotFound(t *testing.T) {
	logger := slog.Default()
	mockucase := new(mocks.UsecaseMock)

	h := handler.RollbackBid(logger, mockucase)

	r := chi.NewRouter()
	r.Put("/api/bids/{tenderId}/rollback/{version}", h)

	mockucase.On("RollbackBid", mock.Anything, mock.Anything, mock.Anything).Return(dto.BidResponse{}, usecase.ErrBidNotFound)

	httpReq := httptest.NewRequest(http.MethodPut, "/api/tenders/d976cd81-3c1f-4d75-9841-1003af7d1e40/rollback/1?username=test_user", nil)

	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, httpReq)

	res := rr.Result()

	assert.Equal(t, http.StatusNotFound, res.StatusCode)
}
