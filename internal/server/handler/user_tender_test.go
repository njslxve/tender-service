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

func TestGetUserTenders(t *testing.T) {
	logger := slog.Default()
	mockucase := new(mocks.UsecaseMock)

	h := handler.GetUserTenders(logger, mockucase)

	mockucase.On("GetUserTenders", "test_user", mock.Anything, mock.Anything).Return([]dto.TenderResponse{}, nil)

	httpReq := httptest.NewRequest(http.MethodGet, "/api/tenders?username=test_user", nil)
	rr := httptest.NewRecorder()

	h(rr, httpReq)

	res := rr.Result()

	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestGetUserTendersUserNotFound(t *testing.T) {
	logger := slog.Default()
	mockucase := new(mocks.UsecaseMock)

	h := handler.GetUserTenders(logger, mockucase)

	mockucase.On("GetUserTenders", "test_user", mock.Anything, mock.Anything).Return([]dto.TenderResponse{}, usecase.ErrUserNotFound)

	httpReq := httptest.NewRequest(http.MethodGet, "/api/tenders?username=test_user", nil)
	rr := httptest.NewRecorder()

	h(rr, httpReq)

	res := rr.Result()

	assert.Equal(t, http.StatusUnauthorized, res.StatusCode)
}
