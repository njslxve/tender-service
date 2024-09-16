package usecase

import (
	"fmt"
	"time"

	"github.com/njslxve/tender-service/internal/dto"
	"github.com/njslxve/tender-service/internal/entity"
)

func (u *Usecase) CreateBid(req dto.BidRequest) (dto.BidResponse, error) {
	const op = "usecase.CreateBid"

	new := entity.Bid{
		Name:        req.Name,
		Description: req.Description,
		TenderID:    req.TendterID,
		AuthorType:  req.AuthorType,
		AuthorID:    req.AuthorID,
	}

	fromDB, err := u.db.CreateBid(new)
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
