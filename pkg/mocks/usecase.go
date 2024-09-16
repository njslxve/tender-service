package mocks

import (
	"github.com/njslxve/tender-service/internal/dto"
	"github.com/stretchr/testify/mock"
)

type UsecaseMock struct {
	mock.Mock
}

func (m *UsecaseMock) CreateTender(req dto.TenderRequest) (dto.TenderResponse, error) {
	args := m.Called(req)

	return args.Get(0).(dto.TenderResponse), args.Error(1)
}

func (m *UsecaseMock) GetTenders(serviceType string, limit string, offset string) ([]dto.TenderResponse, error) {
	args := m.Called(serviceType, limit, offset)

	return args.Get(0).([]dto.TenderResponse), args.Error(1)
}

func (m *UsecaseMock) GetUserTenders(username string, limit string, offset string) ([]dto.TenderResponse, error) {
	args := m.Called(username, limit, offset)

	return args.Get(0).([]dto.TenderResponse), args.Error(1)
}

func (m *UsecaseMock) TenderStatus(tenderId string, username string) (string, error) {
	args := m.Called(tenderId, username)

	return args.String(0), args.Error(1)
}

func (m *UsecaseMock) UpdateTenderStatus(tenderId string, username string, status string) (dto.TenderResponse, error) {
	args := m.Called(tenderId, username, status)

	return args.Get(0).(dto.TenderResponse), args.Error(1)
}

func (m *UsecaseMock) EditTender(tenderId string, username string, req dto.TenderRequest) (dto.TenderResponse, error) {
	args := m.Called(tenderId, username, req)

	return args.Get(0).(dto.TenderResponse), args.Error(1)
}

func (m *UsecaseMock) RollbackTender(tenderId string, version string, username string) (dto.TenderResponse, error) {
	args := m.Called(tenderId, version)

	return args.Get(0).(dto.TenderResponse), args.Error(1)
}

func (m *UsecaseMock) CreateBid(req dto.BidRequest) (dto.BidResponse, error) {
	args := m.Called(req)

	return args.Get(0).(dto.BidResponse), args.Error(1)
}

func (m *UsecaseMock) GetUserBids(username string, limit string, offset string) ([]dto.BidResponse, error) {
	args := m.Called(username, limit, offset)

	return args.Get(0).([]dto.BidResponse), args.Error(1)
}

func (m *UsecaseMock) BidStatus(bidId string, username string) (string, error) {
	args := m.Called(bidId, username)

	return args.String(0), args.Error(1)
}
func (m *UsecaseMock) UpdateBidStatus(bidId string, username string, status string) (dto.BidResponse, error) {
	args := m.Called(bidId, username, status)

	return args.Get(0).(dto.BidResponse), args.Error(1)
}

func (m *UsecaseMock) RollbackBid(bidId string, version string, username string) (dto.BidResponse, error) {
	args := m.Called(bidId, version)

	return args.Get(0).(dto.BidResponse), args.Error(1)
}
