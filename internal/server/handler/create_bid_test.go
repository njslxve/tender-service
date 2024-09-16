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

func TestCreateBid(t *testing.T) {
	logger := slog.Default()
	mockucase := new(mocks.UsecaseMock)

	h := handler.CreateBid(logger, mockucase)

	req := dto.BidRequest{
		Name:        "test",
		Description: "test",
		TendterID:   "2599da85-8a05-4c2f-bd4a-755c21cd788e",
		AuthorType:  "User",
		AuthorID:    "2599da85-8a05-4c2f-bd4a-755c21cd788e",
	}

	mockucase.On("CreateBid", req).Return(dto.BidResponse{}, nil)

	reqBody, _ := json.Marshal(req)
	httpReq := httptest.NewRequest(http.MethodPost, "/api/bids/new", bytes.NewBuffer(reqBody))
	rr := httptest.NewRecorder()

	h(rr, httpReq)

	res := rr.Result()

	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestCreateBidBadRequest(t *testing.T) {
	logger := slog.Default()
	mockucase := new(mocks.UsecaseMock)

	h := handler.CreateBid(logger, mockucase)

	httpReq := httptest.NewRequest("POST", "/api/bids/new", nil)
	rr := httptest.NewRecorder()

	h(rr, httpReq)

	res := rr.Result()
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func TestCreateBidUnauthorized(t *testing.T) {
	logger := slog.Default()
	mockucase := new(mocks.UsecaseMock)

	h := handler.CreateBid(logger, mockucase)

	req := dto.BidRequest{
		Name:        "test",
		Description: "test",
		TendterID:   "2599da85-8a05-4c2f-bd4a-755c21cd788e",
		AuthorType:  "User",
		AuthorID:    "2599da85-8a05-4c2f-bd4a-755c21cd788e",
	}

	mockucase.On("CreateBid", req).Return(dto.BidResponse{}, usecase.ErrUserNotFound)

	reqBody, _ := json.Marshal(req)

	httpReq := httptest.NewRequest("POST", "/api/bids/new", bytes.NewBuffer(reqBody))
	rr := httptest.NewRecorder()

	h(rr, httpReq)

	res := rr.Result()
	assert.Equal(t, http.StatusUnauthorized, res.StatusCode)
}

func TestCreateBidForbidden(t *testing.T) {
	logger := slog.Default()
	mockucase := new(mocks.UsecaseMock)

	h := handler.CreateBid(logger, mockucase)

	req := dto.BidRequest{
		Name:        "test",
		Description: "test",
		TendterID:   "2599da85-8a05-4c2f-bd4a-755c21cd788e",
		AuthorType:  "User",
		AuthorID:    "2599da85-8a05-4c2f-bd4a-755c21cd788e",
	}

	mockucase.On("CreateBid", req).Return(dto.BidResponse{}, usecase.ErrNotPermissions)

	reqBody, _ := json.Marshal(req)

	httpReq := httptest.NewRequest("POST", "/api/bids/new", bytes.NewBuffer(reqBody))
	rr := httptest.NewRecorder()

	h(rr, httpReq)

	res := rr.Result()
	assert.Equal(t, http.StatusForbidden, res.StatusCode)
}

func TestCreateBidValidationError(t *testing.T) {
	logger := slog.Default()
	mockucase := new(mocks.UsecaseMock)

	h := handler.CreateBid(logger, mockucase)

	req := dto.BidRequest{
		Name:        "test",
		Description: "test",
		TendterID:   "qwertyuiop",
		AuthorType:  "User",
		AuthorID:    "2599da85-8a05-4c2f-bd4a-755c21cd788e",
	}

	reqBody, _ := json.Marshal(req)

	httpReq := httptest.NewRequest("POST", "/api/bids/new", bytes.NewBuffer(reqBody))
	rr := httptest.NewRecorder()

	h(rr, httpReq)

	res := rr.Result()
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func TestCreateBidTenderNotFound(t *testing.T) {
	logger := slog.Default()
	mockucase := new(mocks.UsecaseMock)

	h := handler.CreateBid(logger, mockucase)

	req := dto.BidRequest{
		Name:        "test",
		Description: "test",
		TendterID:   "2599da85-8a05-4c2f-bd4a-755c21cd788e",
		AuthorType:  "User",
		AuthorID:    "2599da85-8a05-4c2f-bd4a-755c21cd788e",
	}

	mockucase.On("CreateBid", req).Return(dto.BidResponse{}, usecase.ErrTenderNotFound)

	reqBody, _ := json.Marshal(req)

	httpReq := httptest.NewRequest("POST", "/api/bids/new", bytes.NewBuffer(reqBody))
	rr := httptest.NewRecorder()

	h(rr, httpReq)

	res := rr.Result()
	assert.Equal(t, http.StatusNotFound, res.StatusCode)
}
