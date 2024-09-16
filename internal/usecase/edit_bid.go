package usecase

import (
	"fmt"
	"time"

	"github.com/njslxve/tender-service/internal/dto"
	"github.com/njslxve/tender-service/internal/entity"
)

func (u *Usecase) EditBid(bidId string, username string, req dto.BidRequest) (dto.BidResponse, error) {
	const op = "usecase.EditBid"

	old, err := u.db.GetBid(bidId)
	if err != nil {
		return dto.BidResponse{}, fmt.Errorf("%s: %w", op, err)
	}

	var new entity.Bid = old

	if req.Name != "" {
		new.Name = req.Name
	}

	if req.Description != "" {
		new.Description = req.Description
	}

	new.Version++

	new.CreatedAt = time.Now()

	b, err := u.db.EditBid(new)
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
