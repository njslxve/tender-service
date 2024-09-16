package usecase

import (
	"fmt"
	"time"

	"github.com/njslxve/tender-service/internal/dto"
)

func (u *Usecase) GetUserBids(username string, limit string, offset string) ([]dto.BidResponse, error) {
	const op = "usecase.GetUserBids"

	bids := make([]dto.BidResponse, 0)

	fromDB, err := u.db.GetUserBids(username, limit, offset)
	if err != nil {
		return bids, fmt.Errorf("%s: %w", op, err)
	}

	for _, b := range fromDB {
		bids = append(bids, dto.BidResponse{
			ID:         b.ID,
			Name:       b.Name,
			Status:     b.Status,
			AuthorType: b.AuthorType,
			AuthorID:   b.AuthorID,
			Version:    b.Version,
			CreatedAt:  b.CreatedAt.Format(time.RFC3339),
		})
	}

	if len(bids) > 50 {
		return bids[:50], nil
	} else {
		return bids, nil
	}
}
