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
