package usecase

import (
	"fmt"
	"time"

	"github.com/njslxve/tender-service/internal/dto"
	"github.com/njslxve/tender-service/internal/entity"
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
