package handler_test

import (
	"bytes"
	"encoding/json"
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

func TestEditTender(t *testing.T) {
	logger := slog.Default()
	mockucase := new(mocks.UsecaseMock)

	h := handler.EditTender(logger, mockucase)

	r := chi.NewRouter()
	r.Patch("/api/tenders/{tenderId}/edit", h)

	mockucase.On("EditTender", mock.Anything, mock.Anything, mock.Anything).Return(dto.TenderResponse{}, nil)

	req := dto.TenderRequest{
		Name:        "test_user",
		Description: "test_description",
	}

	reqBody, _ := json.Marshal(req)

	httpReq := httptest.NewRequest(http.MethodPatch, "/api/tenders/d976cd81-3c1f-4d75-9841-1003af7d1e40/edit?username=test_user", bytes.NewBuffer(reqBody))

	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, httpReq)

	res := rr.Result()

	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestEditTenderValidateParams(t *testing.T) {
	logger := slog.Default()
	mockucase := new(mocks.UsecaseMock)

	h := handler.EditTender(logger, mockucase)

	r := chi.NewRouter()
	r.Patch("/api/tenders/{tenderId}/edit", h)

	req := dto.TenderRequest{
		Name:        "test_user",
		Description: "test_description",
	}

	reqBody, _ := json.Marshal(req)

	httpReq := httptest.NewRequest(http.MethodPatch, "/api/tenders/d976cd81-3c1f0/edit?username=test_user", bytes.NewBuffer(reqBody))

	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, httpReq)

	res := rr.Result()

	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func TestEditTenderValidateBody(t *testing.T) {
	logger := slog.Default()
	mockucase := new(mocks.UsecaseMock)

	h := handler.EditTender(logger, mockucase)

	r := chi.NewRouter()
	r.Patch("/api/tenders/{tenderId}/edit", h)

	mockucase.On("EditTender", mock.Anything, mock.Anything, mock.Anything).Return(dto.TenderResponse{}, nil)

	req := dto.TenderRequest{
		Name:        "test_user",
		Description: "test_description",
		ServiceType: "Aboba",
	}

	reqBody, _ := json.Marshal(req)

	httpReq := httptest.NewRequest(http.MethodPatch, "/api/tenders/d976cd81-3c1f-4d75-9841-1003af7d1e40/edit?username=test_user", bytes.NewBuffer(reqBody))

	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, httpReq)

	res := rr.Result()

	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func TestEditTenderBadRequest(t *testing.T) {
	logger := slog.Default()
	mockucase := new(mocks.UsecaseMock)

	h := handler.EditTender(logger, mockucase)

	r := chi.NewRouter()
	r.Patch("/api/tenders/{tenderId}/edit", h)

	httpReq := httptest.NewRequest(http.MethodPatch, "/api/tenders/d976cd81-3c1f-4d75-9841-1003af7d1e40/edit?username=test_user", nil)
	rr := httptest.NewRecorder()

	h(rr, httpReq)

	res := rr.Result()

	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func TestEditTenderUnauthorized(t *testing.T) {
	logger := slog.Default()
	mockucase := new(mocks.UsecaseMock)

	h := handler.EditTender(logger, mockucase)

	r := chi.NewRouter()
	r.Patch("/api/tenders/{tenderId}/edit", h)

	mockucase.On("EditTender", mock.Anything, mock.Anything, mock.Anything).Return(dto.TenderResponse{}, usecase.ErrUserNotFound)

	req := dto.TenderRequest{
		Name:        "test_user",
		Description: "test_description",
	}

	reqBody, _ := json.Marshal(req)

	httpReq := httptest.NewRequest(http.MethodPatch, "/api/tenders/d976cd81-3c1f-4d75-9841-1003af7d1e40/edit?username=test_user", bytes.NewBuffer(reqBody))

	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, httpReq)

	res := rr.Result()

	assert.Equal(t, http.StatusUnauthorized, res.StatusCode)
}

func TestEditTenderForbidden(t *testing.T) {
	logger := slog.Default()
	mockucase := new(mocks.UsecaseMock)

	h := handler.EditTender(logger, mockucase)

	r := chi.NewRouter()
	r.Patch("/api/tenders/{tenderId}/edit", h)

	mockucase.On("EditTender", mock.Anything, mock.Anything, mock.Anything).Return(dto.TenderResponse{}, usecase.ErrNotPermissions)

	req := dto.TenderRequest{
		Name:        "test_user",
		Description: "test_description",
	}

	reqBody, _ := json.Marshal(req)

	httpReq := httptest.NewRequest(http.MethodPatch, "/api/tenders/d976cd81-3c1f-4d75-9841-1003af7d1e40/edit?username=test_user", bytes.NewBuffer(reqBody))

	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, httpReq)

	res := rr.Result()

	assert.Equal(t, http.StatusForbidden, res.StatusCode)
}

func TestEditTenderTenderNotFound(t *testing.T) {
	logger := slog.Default()
	mockucase := new(mocks.UsecaseMock)

	h := handler.EditTender(logger, mockucase)

	r := chi.NewRouter()
	r.Patch("/api/tenders/{tenderId}/edit", h)

	mockucase.On("EditTender", mock.Anything, mock.Anything, mock.Anything).Return(dto.TenderResponse{}, usecase.ErrTenderNotFound)

	req := dto.TenderRequest{
		Name:        "test_user",
		Description: "test_description",
	}

	reqBody, _ := json.Marshal(req)

	httpReq := httptest.NewRequest(http.MethodPatch, "/api/tenders/d976cd81-3c1f-4d75-9841-1003af7d1e40/edit?username=test_user", bytes.NewBuffer(reqBody))

	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, httpReq)

	res := rr.Result()

	assert.Equal(t, http.StatusNotFound, res.StatusCode)
}
