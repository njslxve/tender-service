package usecase

import (
	"fmt"
	"time"

	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725721237-team-77964/zadanie-6105/internal/dto"
)

func (u *Usecase) GetBidsForTender(tenderId string, username string, limit string, offset string) ([]dto.BidResponse, error) {
	const op = "usecase.GetBidsForTender"

	bids := make([]dto.BidResponse, 0)

	// err := u.Permissions(username, tenderId)
	// if err != nil {
	// 	return bids, fmt.Errorf("%s: %w", op, err)
	// }

	fromDB, err := u.db.GetBidsForTender(tenderId, username, limit, offset)
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
