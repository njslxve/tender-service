package usecase

import (
	"fmt"
	"time"

	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725721237-team-77964/zadanie-6105/internal/dto"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725721237-team-77964/zadanie-6105/internal/entity"
)

func (u *Usecase) SubmitBidFeedback(bidId string, feedback string, username string) (dto.BidResponse, error) {
	const op = "usecase.SubmitBidFeedback"

	new := entity.BidFeedback{
		BidID:       bidId,
		Description: feedback,
		AuthorID:    username,
	}

	fromDB, err := u.db.SubmitBidFeedback(new)
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
