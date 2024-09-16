package usecase

import (
	"fmt"
	"strconv"
	"time"

	"github.com/njslxve/tender-service/internal/dto"
	"github.com/njslxve/tender-service/internal/entity"
)

func (u *Usecase) RollbackBid(bidId string, version string, username string) (dto.BidResponse, error) {
	const op = "usecase.RollbackBid"

	ver, err := strconv.Atoi(version)
	if err != nil {
		return dto.BidResponse{}, fmt.Errorf("%s: %w", op, err)
	}

	old, err := u.db.GetBidByVersion(bidId, int32(ver))
	if err != nil {
		return dto.BidResponse{}, fmt.Errorf("%s: %w", op, err)
	}

	var new entity.Bid = old

	v, err := u.db.GetBidLastVersion(bidId)
	if err != nil {
		return dto.BidResponse{}, fmt.Errorf("%s: %w", op, err)
	}

	new.Version = v + 1

	new.CreatedAt = time.Now()

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
