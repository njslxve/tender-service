package usecase

import (
	"fmt"
	"time"

	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725721237-team-77964/zadanie-6105/internal/dto"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725721237-team-77964/zadanie-6105/internal/entity"
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
