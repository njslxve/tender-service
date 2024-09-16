package handler_test

import (
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/njslxve/tender-service/internal/dto"
	"github.com/njslxve/tender-service/internal/server/handler"
	"github.com/njslxve/tender-service/internal/usecase"
	"github.com/njslxve/tender-service/pkg/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetUserBids(t *testing.T) {
	logger := slog.Default()
	mockucase := new(mocks.UsecaseMock)

	h := handler.GetUserBids(logger, mockucase)

	mockucase.On("GetUserBids", "test_user", mock.Anything, mock.Anything).Return([]dto.BidResponse{}, nil)

	httpReq := httptest.NewRequest(http.MethodGet, "/api/bids?username=test_user", nil)
	rr := httptest.NewRecorder()

	h(rr, httpReq)

	res := rr.Result()

	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestGetUserBidsUserNotFound(t *testing.T) {
	logger := slog.Default()
	mockucase := new(mocks.UsecaseMock)

	h := handler.GetUserBids(logger, mockucase)

	mockucase.On("GetUserBids", "test_user", mock.Anything, mock.Anything).Return([]dto.BidResponse{}, usecase.ErrUserNotFound)

	httpReq := httptest.NewRequest(http.MethodGet, "/api/bids?username=test_user", nil)
	rr := httptest.NewRecorder()

	h(rr, httpReq)

	res := rr.Result()

	assert.Equal(t, http.StatusUnauthorized, res.StatusCode)
}
