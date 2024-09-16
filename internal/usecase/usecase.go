package usecase

import (
	"log/slog"

	"github.com/njslxve/tender-service/internal/config"
	"github.com/njslxve/tender-service/internal/entity"
)

type StorageInterface interface {
	CreateTender(tender entity.Tender) (entity.Tender, error)
	GetTenders(serviceType string, limit string, offset string) ([]entity.Tender, error)
	GetUserTenders(username string, limit string, offset string) ([]entity.Tender, error)
	TenderStatus(tenderId string, username string) (string, error)
	UpdateTenderStatus(tenderId string, username string, status string) (entity.Tender, error)
	GetTenderByID(tenderId string) (entity.Tender, error)
	FoundUser(username string) error
	EditTender(tender entity.Tender) (entity.Tender, error)
	CreateBid(bid entity.Bid) (entity.Bid, error)
	GetUserBids(username string, limit string, offset string) ([]entity.Bid, error)
	IsResponsible(user string, org string) error
	GetBidsForTender(tenderId string, username string, limit string, offset string) ([]entity.Bid, error)
	BidStatus(bidId string, username string) (string, error)
	UpdateBidStatus(bidId string, username string, status string) (entity.Bid, error)
	EditBid(bid entity.Bid) (entity.Bid, error)
	GetBid(bidId string) (entity.Bid, error)
	SubmitBidFeedback(feedback entity.BidFeedback) (entity.Bid, error)
	GetTenderIdByBidId(bidId string) (string, error)
	SubmitBidDecision(bidDecision entity.BidDecision) (entity.Bid, error)
	GetBidReviews(authorUsername string, limit string, offset string) ([]entity.BidFeedback, error)
	GetTenderLastVersion(tenderId string) (int32, error)
	GetTenderByVersion(tenderId string, version int32) (entity.Tender, error)
	GetBidLastVersion(bidId string) (int32, error)
	GetBidByVersion(bidId string, version int32) (entity.Bid, error)
}

type Usecase struct {
	cfg    *config.Config
	logger *slog.Logger
	db     StorageInterface
}

func New(cfg *config.Config, logger *slog.Logger, db StorageInterface) *Usecase {
	return &Usecase{
		cfg:    cfg,
		logger: logger,
		db:     db,
	}
}
