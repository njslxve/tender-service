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

func TestGetBidStatus(t *testing.T) {
	logger := slog.Default()
	mockucase := new(mocks.UsecaseMock)

	h := handler.GetBidStatus(logger, mockucase)

	mockucase.On("BidStatus", mock.Anything, mock.Anything).Return("Created", nil)

	r := chi.NewRouter()
	r.Get("/api/bids/{bidId}/status", h)

	httpReq := httptest.NewRequest(http.MethodGet, "/api/bids/d976cd81-3c1f-4d75-9841-1003af7d1e40/status?username=test_user", nil)
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, httpReq)

	res := rr.Result()

	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestGetBidStatusValidationError(t *testing.T) {
	logger := slog.Default()
	mockucase := new(mocks.UsecaseMock)

	h := handler.GetBidStatus(logger, mockucase)

	r := chi.NewRouter()
	r.Get("/api/bids/{bidId}/status", h)

	httpReq := httptest.NewRequest(http.MethodGet, "/api/bids/aboba/status?username=test_user", nil)
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, httpReq)

	res := rr.Result()

	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func TestGetBidStatusUnauthorized(t *testing.T) {
	logger := slog.Default()
	mockucase := new(mocks.UsecaseMock)

	h := handler.GetBidStatus(logger, mockucase)

	mockucase.On("BidStatus", mock.Anything, mock.Anything).Return("", usecase.ErrUserNotFound)

	r := chi.NewRouter()
	r.Get("/api/bids/{bidId}/status", h)

	httpReq := httptest.NewRequest(http.MethodGet, "/api/bids/d976cd81-3c1f-4d75-9841-1003af7d1e40/status?username=test_user", nil)
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, httpReq)

	res := rr.Result()

	assert.Equal(t, http.StatusUnauthorized, res.StatusCode)
}

func TestGetBidStatusForbidden(t *testing.T) {
	logger := slog.Default()
	mockucase := new(mocks.UsecaseMock)

	h := handler.GetBidStatus(logger, mockucase)

	mockucase.On("BidStatus", mock.Anything, mock.Anything).Return("", usecase.ErrNotPermissions)

	r := chi.NewRouter()
	r.Get("/api/bids/{bidId}/status", h)

	httpReq := httptest.NewRequest(http.MethodGet, "/api/bids/d976cd81-3c1f-4d75-9841-1003af7d1e40/status?username=test_user", nil)
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, httpReq)

	res := rr.Result()

	assert.Equal(t, http.StatusForbidden, res.StatusCode)
}

func TestGetBidStatusBidNotFound(t *testing.T) {
	logger := slog.Default()
	mockucase := new(mocks.UsecaseMock)

	h := handler.GetBidStatus(logger, mockucase)

	mockucase.On("BidStatus", mock.Anything, mock.Anything).Return("", usecase.ErrBidNotFound)

	r := chi.NewRouter()
	r.Get("/api/bids/{bidId}/status", h)

	httpReq := httptest.NewRequest(http.MethodGet, "/api/bids/d976cd81-3c1f-4d75-9841-1003af7d1e40/status?username=test_user", nil)
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, httpReq)

	res := rr.Result()

	assert.Equal(t, http.StatusNotFound, res.StatusCode)
}

func TestUpdateBidStatus(t *testing.T) {
	logger := slog.Default()
	mockucase := new(mocks.UsecaseMock)

	h := handler.UpdateBidStatus(logger, mockucase)

	mockucase.On("UpdateBidStatus", mock.Anything, mock.Anything, mock.Anything).Return(dto.BidResponse{}, nil)

	r := chi.NewRouter()
	r.Put("/api/bids/{bidId}/status", h)

	httpReq := httptest.NewRequest(http.MethodPut, "/api/bids/d976cd81-3c1f-4d75-9841-1003af7d1e40/status?username=test_user&status=Created", nil)
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, httpReq)

	res := rr.Result()

	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestUpdateBidStatusValidationError(t *testing.T) {
	logger := slog.Default()
	mockucase := new(mocks.UsecaseMock)

	h := handler.UpdateBidStatus(logger, mockucase)

	r := chi.NewRouter()
	r.Put("/api/bids/{bidId}/status", h)

	httpReq := httptest.NewRequest(http.MethodPut, "/api/bids/aboba/status?username=test_user&status=Created", nil)
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, httpReq)

	res := rr.Result()

	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func TestUpdateBidStatusUnauthorized(t *testing.T) {
	logger := slog.Default()
	mockucase := new(mocks.UsecaseMock)

	h := handler.UpdateBidStatus(logger, mockucase)

	mockucase.On("UpdateBidStatus", mock.Anything, mock.Anything, mock.Anything).Return(dto.BidResponse{}, usecase.ErrUserNotFound)

	r := chi.NewRouter()
	r.Put("/api/bids/{bidId}/status", h)

	httpReq := httptest.NewRequest(http.MethodPut, "/api/bids/d976cd81-3c1f-4d75-9841-1003af7d1e40/status?username=test_user&status=Created", nil)
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, httpReq)

	res := rr.Result()

	assert.Equal(t, http.StatusUnauthorized, res.StatusCode)
}

func TestUpdateBidStatusForbidden(t *testing.T) {
	logger := slog.Default()
	mockucase := new(mocks.UsecaseMock)

	h := handler.UpdateBidStatus(logger, mockucase)

	mockucase.On("UpdateBidStatus", mock.Anything, mock.Anything, mock.Anything).Return(dto.BidResponse{}, usecase.ErrNotPermissions)

	r := chi.NewRouter()
	r.Put("/api/bids/{bidId}/status", h)

	httpReq := httptest.NewRequest(http.MethodPut, "/api/bids/d976cd81-3c1f-4d75-9841-1003af7d1e40/status?username=test_user&status=Created", nil)
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, httpReq)

	res := rr.Result()

	assert.Equal(t, http.StatusForbidden, res.StatusCode)
}

func TestUpdateBidStatusBidNotFound(t *testing.T) {
	logger := slog.Default()
	mockucase := new(mocks.UsecaseMock)

	h := handler.UpdateBidStatus(logger, mockucase)

	mockucase.On("UpdateBidStatus", mock.Anything, mock.Anything, mock.Anything).Return(dto.BidResponse{}, usecase.ErrBidNotFound)

	r := chi.NewRouter()
	r.Put("/api/bids/{bidId}/status", h)

	httpReq := httptest.NewRequest(http.MethodPut, "/api/bids/d976cd81-3c1f-4d75-9841-1003af7d1e40/status?username=test_user&status=Created", nil)
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, httpReq)

	res := rr.Result()

	assert.Equal(t, http.StatusNotFound, res.StatusCode)
}
