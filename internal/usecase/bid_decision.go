package usecase

import (
	"fmt"
	"time"

	"github.com/njslxve/tender-service/internal/dto"
	"github.com/njslxve/tender-service/internal/entity"
)

func (u *Usecase) SubmitBidDecision(bidId string, username string, decision string) (dto.BidResponse, error) {
	const op = "usecase.SubmitBidDecision"

	tenderID, err := u.db.GetTenderIdByBidId(bidId)
	if err != nil {
		return dto.BidResponse{}, fmt.Errorf("%s: %w", op, err)
	}

	var bidDecision entity.BidDecision

	if decision == "Approved" {
		bidDecision.ApprovedCount++
	} else {
		bidDecision.RejectedCount++
	}

	bidDecision.BidID = bidId
	bidDecision.TenderID = tenderID
	bidDecision.Decision = decision

	fromDB, err := u.db.SubmitBidDecision(bidDecision)
	if err != nil {
		return dto.BidResponse{}, fmt.Errorf("%s: %w", op, err)
	}

	bid := dto.BidResponse{
		ID:         fromDB.ID,
		Name:       fromDB.Name,
		Status:     fromDB.Status,
		AuthorType: fromDB.AuthorType,
		AuthorID:   fromDB.AuthorID,
		Version:    fromDB.Version,
		CreatedAt:  fromDB.CreatedAt.Format(time.RFC3339),
	}

	return bid, nil
}
