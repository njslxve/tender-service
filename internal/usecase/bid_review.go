package usecase

import (
	"fmt"
	"time"

	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725721237-team-77964/zadanie-6105/internal/dto"
)

func (u *Usecase) GetBidReviews(
	tenderId string,
	authorUsername string,
	requesterUsername string,
	limit string,
	offset string,
) ([]dto.ReviewResponse, error) {
	const op = "usecase.BidReviews"

	reviews := make([]dto.ReviewResponse, 0)

	fromDB, err := u.db.GetBidReviews(authorUsername, limit, offset)
	if err != nil {
		return reviews, fmt.Errorf("%s: %w", op, err)
	}

	for _, r := range fromDB {
		reviews = append(reviews, dto.ReviewResponse{
			ID:          r.ID,
			Description: r.Description,
			CreatedAt:   r.CreatedAt.Format(time.RFC3339),
		})
	}

	return reviews, nil
}
