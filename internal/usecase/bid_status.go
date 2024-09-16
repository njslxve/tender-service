package usecase

import (
	"fmt"
	"time"

	"github.com/njslxve/tender-service/internal/dto"
)

func (u *Usecase) BidStatus(bidId string, username string) (string, error) {
	const op = "usecase.GetBidStatus"

	status, err := u.db.BidStatus(bidId, username)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return status, nil
}

func (u *Usecase) UpdateBidStatus(bidId string, username string, status string) (dto.BidResponse, error) {
	const op = "usecase.UpdateBidStatus"

	b, err := u.db.UpdateBidStatus(bidId, username, status)
	if err != nil {
		return dto.BidResponse{}, fmt.Errorf("%s: %w", op, err)
	}

	bid := dto.BidResponse{
		ID:         b.ID,
		Name:       b.Name,
		Status:     b.Status,
		AuthorType: b.AuthorType,
		AuthorID:   b.AuthorID,
		Version:    b.Version,
		CreatedAt:  b.CreatedAt.Format(time.RFC3339),
	}

	return bid, nil
}
