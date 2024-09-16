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
	"github.com/stretchr/testify/assert"
)

func TestCreateHandler(t *testing.T) {
	logger := slog.Default()

	mockucase := new(mocks.UsecaseMock)

	h := handler.CreateTender(logger, mockucase)

	req := dto.TenderRequest{
		Name:            "test",
		Description:     "test",
		ServiceType:     "Delivery",
		OrganizationId:  "2599da85-8a05-4c2f-bd4a-755c21cd788e",
		CreatorUsername: "test_user",
	}

	mockucase.On("CreateTender", req).Return(dto.TenderResponse{}, nil)

	reqBody, _ := json.Marshal(req)
	httpReq := httptest.NewRequest(http.MethodPost, "/api/tenders/new", bytes.NewBuffer(reqBody))
	rr := httptest.NewRecorder()

	h(rr, httpReq)

	res := rr.Result()

	assert.Equal(t, http.StatusOK, res.StatusCode)

	mockucase.AssertExpectations(t)
}

func TestCreateTenderBadRequest(t *testing.T) {
	logger := slog.Default()
	mockucase := new(mocks.UsecaseMock)

	h := handler.CreateTender(logger, mockucase)

	httpReq := httptest.NewRequest("POST", "/api/tenders/new", nil)
	rr := httptest.NewRecorder()

	h(rr, httpReq)

	res := rr.Result()
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func TestCreateTenderValidationError(t *testing.T) {
	logger := slog.Default()
	mockucase := new(mocks.UsecaseMock)

	h := handler.CreateTender(logger, mockucase)

	req := dto.TenderRequest{
		Name:            "test",
		Description:     "test",
		ServiceType:     "Aboba",
		OrganizationId:  "2599da85-8a05-4c2f-bd4a-755c21cd788e",
		CreatorUsername: "test_user",
	}

	reqBody, _ := json.Marshal(req)
	httpReq := httptest.NewRequest(http.MethodPost, "/api/tenders/new", bytes.NewBuffer(reqBody))
	rr := httptest.NewRecorder()

	h(rr, httpReq)

	res := rr.Result()

	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func TestCreateTenderUserNotFound(t *testing.T) {
	logger := slog.Default()
	mockucase := new(mocks.UsecaseMock)

	h := handler.CreateTender(logger, mockucase)

	req := dto.TenderRequest{
		Name:            "test",
		Description:     "test",
		ServiceType:     "Delivery",
		OrganizationId:  "2599da85-8a05-4c2f-bd4a-755c21cd788e",
		CreatorUsername: "test_user",
	}

	reqBody, _ := json.Marshal(req)
	httpReq := httptest.NewRequest(http.MethodPost, "/api/tenders/new", bytes.NewBuffer(reqBody))
	rr := httptest.NewRecorder()

	mockucase.On("CreateTender", req).Return(dto.TenderResponse{}, usecase.ErrUserNotFound)

	h(rr, httpReq)

	res := rr.Result()

	assert.Equal(t, http.StatusUnauthorized, res.StatusCode)
}

func TestCreateTenderNotPermissions(t *testing.T) {
	logger := slog.Default()
	mockucase := new(mocks.UsecaseMock)

	h := handler.CreateTender(logger, mockucase)

	req := dto.TenderRequest{
		Name:            "test",
		Description:     "test",
		ServiceType:     "Delivery",
		OrganizationId:  "2599da85-8a05-4c2f-bd4a-755c21cd788e",
		CreatorUsername: "test_user",
	}

	reqBody, _ := json.Marshal(req)
	httpReq := httptest.NewRequest(http.MethodPost, "/api/tenders/new", bytes.NewBuffer(reqBody))
	rr := httptest.NewRecorder()

	mockucase.On("CreateTender", req).Return(dto.TenderResponse{}, usecase.ErrNotPermissions)

	h(rr, httpReq)

	res := rr.Result()

	assert.Equal(t, http.StatusForbidden, res.StatusCode)
}
