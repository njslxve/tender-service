package storage

import (
	"context"
	"fmt"
	"strconv"

	sq "github.com/Masterminds/squirrel"
	"github.com/njslxve/tender-service/internal/entity"
)

func (s *Storage) GetUserBids(username string, limit string, offset string) ([]entity.Bid, error) {
	const op = "storage.GetUserBids"

	var lim, off int

	if limit != "" {
		l, err := strconv.Atoi(limit)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		lim = l
	}

	if offset != "" {
		o, err := strconv.Atoi(offset)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		off = o
	}

	querry := qb.Select(
		"b.id as bid_id",
		"bv.name",
		"bv.status",
		"bv.author_type",
		"bv.author_id",
		"bv.version",
		"bv.created_at",
	).
		From("bids b").
		Join("bids_versions bv ON b.id = bv.bid_id").
		Where("bv.version = b.latest_version").
		Where(sq.Eq{"bv.author_id": username}).
		OrderBy("bv.name")

	if limit != "" {
		querry = querry.Limit(uint64(lim))
	}

	if offset != "" {
		querry = querry.Offset(uint64(off))
	}

	sql, args, err := querry.ToSql()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	rows, err := s.db.Query(context.Background(), sql, args...)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var bids []entity.Bid

	for rows.Next() {
		var bid entity.Bid

		err = rows.Scan(&bid.ID,
			&bid.Name,
			&bid.Status,
			&bid.AuthorType,
			&bid.AuthorID,
			&bid.Version,
			&bid.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		bids = append(bids, bid)
	}

	return bids, nil
}
