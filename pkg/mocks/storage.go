package mocks

import (
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725721237-team-77964/zadanie-6105/internal/entity"
	"github.com/stretchr/testify/mock"
)

type StorageMock struct {
	mock.Mock
}

func (m *StorageMock) CreateTender(tender entity.Tender) (entity.Tender, error) {
	args := m.Called(tender)

	return args.Get(0).(entity.Tender), args.Error(1)
}
func (m *StorageMock) GetTenders(serviceType string, limit string, offset string) ([]entity.Tender, error) {
	args := m.Called(serviceType, limit, offset)

	return args.Get(0).([]entity.Tender), args.Error(1)
}
func (m *StorageMock) GetUserTenders(username string, limit string, offset string) ([]entity.Tender, error) {
	args := m.Called(username, limit, offset)

	return args.Get(0).([]entity.Tender), args.Error(1)
}
func (m *StorageMock) TenderStatus(tenderId string, username string) (string, error) {
	args := m.Called(tenderId, username)

	return args.String(0), args.Error(1)
}
func (m *StorageMock) UpdateTenderStatus(tenderId string, username string, status string) (entity.Tender, error) {
	args := m.Called(tenderId, username, status)

	return args.Get(0).(entity.Tender), args.Error(1)
}
func (m *StorageMock) GetTenderByID(tenderId string) (entity.Tender, error) {
	args := m.Called(tenderId)

	return args.Get(0).(entity.Tender), args.Error(1)
}
func (m *StorageMock) FoundUser(username string) error {
	args := m.Called(username)

	return args.Error(0)
}
func (m *StorageMock) EditTender(tender entity.Tender) (entity.Tender, error) {
	args := m.Called(tender)

	return args.Get(0).(entity.Tender), args.Error(1)
}
func (m *StorageMock) CreateBid(bid entity.Bid) (entity.Bid, error) {
	args := m.Called(bid)

	return args.Get(0).(entity.Bid), args.Error(1)
}
func (m *StorageMock) GetUserBids(username string, limit string, offset string) ([]entity.Bid, error) {
	args := m.Called(username, limit, offset)

	return args.Get(0).([]entity.Bid), args.Error(1)
}
func (m *StorageMock) IsResponsible(user string, org string) error {
	args := m.Called(user, org)

	return args.Error(0)
}
func (m *StorageMock) GetBidsForTender(tenderId string, username string, limit string, offset string) ([]entity.Bid, error) {
	args := m.Called(tenderId, username, limit, offset)

	return args.Get(0).([]entity.Bid), args.Error(1)
}
func (m *StorageMock) BidStatus(bidId string, username string) (string, error) {
	args := m.Called(bidId, username)

	return args.String(0), args.Error(1)
}
func (m *StorageMock) UpdateBidStatus(bidId string, username string, status string) (entity.Bid, error) {
	args := m.Called(bidId, username, status)

	return args.Get(0).(entity.Bid), args.Error(1)
}
func (m *StorageMock) EditBid(bid entity.Bid) (entity.Bid, error) {
	args := m.Called(bid)

	return args.Get(0).(entity.Bid), args.Error(1)
}
func (m *StorageMock) GetBid(bidId string) (entity.Bid, error) {
	args := m.Called(bidId)

	return args.Get(0).(entity.Bid), args.Error(1)
}
func (m *StorageMock) SubmitBidFeedback(feedback entity.BidFeedback) (entity.Bid, error) {
	args := m.Called(feedback)

	return args.Get(0).(entity.Bid), args.Error(1)
}
func (m *StorageMock) GetTenderIdByBidId(bidId string) (string, error) {
	args := m.Called(bidId)

	return args.String(0), args.Error(1)
}
func (m *StorageMock) SubmitBidDecision(bidDecision entity.BidDecision) (entity.Bid, error) {
	args := m.Called(bidDecision)

	return args.Get(0).(entity.Bid), args.Error(1)
}
func (m *StorageMock) GetBidReviews(authorUsername string, limit string, offset string) ([]entity.BidFeedback, error) {
	args := m.Called(authorUsername, limit, offset)

	return args.Get(0).([]entity.BidFeedback), args.Error(1)
}
func (m *StorageMock) GetTenderLastVersion(tenderId string) (int32, error) {
	args := m.Called(tenderId)

	return args.Get(0).(int32), args.Error(1)
}
func (m *StorageMock) GetTenderByVersion(tenderId string, version int32) (entity.Tender, error) {
	args := m.Called(tenderId, version)

	return args.Get(0).(entity.Tender), args.Error(1)
}
func (m *StorageMock) GetBidLastVersion(bidId string) (int32, error) {
	args := m.Called(bidId)

	return args.Get(0).(int32), args.Error(1)
}
func (m *StorageMock) GetBidByVersion(bidId string, version int32) (entity.Bid, error) {
	args := m.Called(bidId, version)

	return args.Get(0).(entity.Bid), args.Error(1)
}
