package handler_test

import (
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725721237-team-77964/zadanie-6105/internal/dto"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725721237-team-77964/zadanie-6105/internal/server/handler"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725721237-team-77964/zadanie-6105/internal/usecase"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725721237-team-77964/zadanie-6105/pkg/mocks"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetTenderStatus(t *testing.T) {
	logger := slog.Default()
	mockucase := new(mocks.UsecaseMock)

	h := handler.GetTenderStatus(logger, mockucase)

	r := chi.NewRouter()
	r.Get("/api/tenders/{tenderId}/status", h)

	mockucase.On("TenderStatus", mock.Anything, mock.Anything).Return("Created", nil)

	httpReq := httptest.NewRequest(http.MethodGet, "/api/tenders/d976cd81-3c1f-4d75-9841-1003af7d1e40/status?username=test_user", nil)
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, httpReq)

	res := rr.Result()

	assert.Equal(t, http.StatusOK, res.StatusCode)

	status, _ := io.ReadAll(res.Body)
	defer res.Body.Close()

	assert.Equal(t, "Created", strings.Trim(string(status), "\"\n"))
}

func TestGetTenderStatusValidationError(t *testing.T) {
	logger := slog.Default()
	mockucase := new(mocks.UsecaseMock)

	h := handler.GetTenderStatus(logger, mockucase)

	r := chi.NewRouter()
	r.Get("/api/tenders/{tenderId}/status", h)

	httpReq := httptest.NewRequest(http.MethodGet, "/api/tenders/aboba/status?username=test_user", nil)
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, httpReq)

	res := rr.Result()

	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	mockucase.AssertExpectations(t)
}

func TestGetTenderStatusUnatorized(t *testing.T) {
	logger := slog.Default()
	mockucase := new(mocks.UsecaseMock)

	h := handler.GetTenderStatus(logger, mockucase)

	r := chi.NewRouter()
	r.Get("/api/tenders/{tenderId}/status", h)

	mockucase.On("TenderStatus", mock.Anything, mock.Anything).Return("", usecase.ErrUserNotFound)

	httpReq := httptest.NewRequest(http.MethodGet, "/api/tenders/d976cd81-3c1f-4d75-9841-1003af7d1e40/status", nil)
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, httpReq)

	res := rr.Result()

	assert.Equal(t, http.StatusUnauthorized, res.StatusCode)
}

func TestGetTenderStatusNotForbidden(t *testing.T) {
	logger := slog.Default()
	mockucase := new(mocks.UsecaseMock)

	h := handler.GetTenderStatus(logger, mockucase)

	r := chi.NewRouter()
	r.Get("/api/tenders/{tenderId}/status", h)

	mockucase.On("TenderStatus", mock.Anything, mock.Anything).Return("", usecase.ErrNotPermissions)

	httpReq := httptest.NewRequest(http.MethodGet, "/api/tenders/d976cd81-3c1f-4d75-9841-1003af7d1e40/status?username=test_user", nil)
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, httpReq)

	res := rr.Result()

	assert.Equal(t, http.StatusForbidden, res.StatusCode)
}

func TestGetTenderStatusNotFound(t *testing.T) {
	logger := slog.Default()
	mockucase := new(mocks.UsecaseMock)

	h := handler.GetTenderStatus(logger, mockucase)

	r := chi.NewRouter()
	r.Get("/api/tenders/{tenderId}/status", h)

	mockucase.On("TenderStatus", mock.Anything, mock.Anything).Return("", usecase.ErrTenderNotFound)

	httpReq := httptest.NewRequest(http.MethodGet, "/api/tenders/d976cd81-3c1f-4d75-9841-1003af7d1e40/status?username=test_user", nil)
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, httpReq)

	res := rr.Result()

	assert.Equal(t, http.StatusNotFound, res.StatusCode)
}

func TestUpdateTenderStatus(t *testing.T) {
	logger := slog.Default()
	mockucase := new(mocks.UsecaseMock)

	h := handler.UpdateTenderStatus(logger, mockucase)

	r := chi.NewRouter()
	r.Put("/api/tenders/{tenderId}/status", h)

	mockucase.On("UpdateTenderStatus", mock.Anything, mock.Anything, mock.Anything).Return(dto.TenderResponse{}, nil)

	httpReq := httptest.NewRequest(http.MethodPut, "/api/tenders/d976cd81-3c1f-4d75-9841-1003af7d1e40/status?username=test_user&status=Published", nil)
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, httpReq)

	res := rr.Result()

	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestUpdateTenderStatusValidationError(t *testing.T) {
	logger := slog.Default()
	mockucase := new(mocks.UsecaseMock)

	h := handler.UpdateTenderStatus(logger, mockucase)

	r := chi.NewRouter()
	r.Put("/api/tenders/{tenderId}/status", h)

	httpReq := httptest.NewRequest(http.MethodPut, "/api/tenders/d976cd81-3c1f-4d75-9841-1003af7d1e40/status?username=test_user&status=Removed", nil)
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, httpReq)

	res := rr.Result()

	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	mockucase.AssertExpectations(t)
}

func TestUpdateTenderStatusUnauthorized(t *testing.T) {
	logger := slog.Default()
	mockucase := new(mocks.UsecaseMock)

	h := handler.UpdateTenderStatus(logger, mockucase)

	r := chi.NewRouter()
	r.Put("/api/tenders/{tenderId}/status", h)

	mockucase.On("UpdateTenderStatus", mock.Anything, mock.Anything, mock.Anything).Return(dto.TenderResponse{}, usecase.ErrUserNotFound)

	httpReq := httptest.NewRequest(http.MethodPut, "/api/tenders/d976cd81-3c1f-4d75-9841-1003af7d1e40/status?username=test_user&status=Published", nil)

	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, httpReq)

	res := rr.Result()

	assert.Equal(t, http.StatusUnauthorized, res.StatusCode)
}

func TestUpdateTenderStatusForbidden(t *testing.T) {
	logger := slog.Default()
	mockucase := new(mocks.UsecaseMock)

	h := handler.UpdateTenderStatus(logger, mockucase)

	r := chi.NewRouter()
	r.Put("/api/tenders/{tenderId}/status", h)

	mockucase.On("UpdateTenderStatus", mock.Anything, mock.Anything, mock.Anything).Return(dto.TenderResponse{}, usecase.ErrNotPermissions)

	httpReq := httptest.NewRequest(http.MethodPut, "/api/tenders/d976cd81-3c1f-4d75-9841-1003af7d1e40/status?username=test_user&status=Published", nil)

	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, httpReq)

	res := rr.Result()

	assert.Equal(t, http.StatusForbidden, res.StatusCode)
}

func TestUpdateTenderStatusNotFound(t *testing.T) {
	logger := slog.Default()
	mockucase := new(mocks.UsecaseMock)

	h := handler.UpdateTenderStatus(logger, mockucase)

	r := chi.NewRouter()
	r.Put("/api/tenders/{tenderId}/status", h)

	mockucase.On("UpdateTenderStatus", mock.Anything, mock.Anything, mock.Anything).Return(dto.TenderResponse{}, usecase.ErrTenderNotFound)

	httpReq := httptest.NewRequest(http.MethodPut, "/api/tenders/d976cd81-3c1f-4d75-9841-1003af7d1e40/status?username=test_user&status=Published", nil)

	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, httpReq)

	res := rr.Result()

	assert.Equal(t, http.StatusNotFound, res.StatusCode)
	mockucase.AssertExpectations(t)
}
