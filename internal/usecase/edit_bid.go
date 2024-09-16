package usecase

import (
	"fmt"
	"time"

	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725721237-team-77964/zadanie-6105/internal/dto"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725721237-team-77964/zadanie-6105/internal/entity"
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
