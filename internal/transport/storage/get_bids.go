package storage

import (
	"context"
	"fmt"
	"strconv"

	sq "github.com/Masterminds/squirrel"
	"github.com/njslxve/tender-service/internal/entity"
)

func (s *Storage) GetBidsForTender(tenderId string, username string, limit string, offset string) ([]entity.Bid, error) {
	const op = "storage.GetBidsForTender"

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
		Where(sq.Eq{"bv.tender_id": tenderId}).
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

	bids := make([]entity.Bid, 0)

	for rows.Next() {
		var bid entity.Bid

		err = rows.Scan(
			&bid.ID,
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

func (s *Storage) GetBid(bidId string) (entity.Bid, error) {
	const op = "storage.GetBid"

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
		Where(sq.Eq{"b.id": bidId})

	sql, args, err := querry.ToSql()
	if err != nil {
		return entity.Bid{}, fmt.Errorf("%s: %w", op, err)
	}

	row := s.db.QueryRow(context.Background(), sql, args...)

	var bid entity.Bid

	err = row.Scan(
		&bid.ID,
		&bid.Name,
		&bid.Status,
		&bid.AuthorType,
		&bid.AuthorID,
		&bid.Version,
		&bid.CreatedAt,
	)
	if err != nil {
		return entity.Bid{}, fmt.Errorf("%s: %w", op, err)
	}

	return bid, nil
}
