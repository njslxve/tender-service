package storage

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/njslxve/tender-service/internal/entity"
)

func (s *Storage) CreateBid(bid entity.Bid) (entity.Bid, error) {
	const op = "storage.CreateBid"

	var status string

	if bid.Status == "" {
		status = "Created"
	} else {
		status = bid.Status
	}

	if bid.Version == 0 {
		bid.Version = 1
	}

	tx, err := s.db.Begin(context.Background())
	if err != nil {
		return entity.Bid{}, fmt.Errorf("%s: %w", op, err)
	}

	var newBidID string

	querry := qb.Insert("bids").
		Columns("latest_version").
		Values(bid.Version).
		Suffix("RETURNING id")

	sql, args, err := querry.ToSql()
	if err != nil {
		return entity.Bid{}, fmt.Errorf("%s: %w", op, err)
	}

	row := tx.QueryRow(context.Background(), sql, args...)

	err = row.Scan(&newBidID)
	if err != nil {
		tx.Rollback(context.Background())
		return entity.Bid{}, fmt.Errorf("%s: %w", op, err)
	}

	querry = qb.Insert("bids_versions").
		Columns(
			"bid_id",
			"name",
			"description",
			"tender_id",
			"status",
			"author_type",
			"author_id",
			"version",
		).
		Values(
			newBidID,
			bid.Name,
			bid.Description,
			bid.TenderID,
			status,
			bid.AuthorType,
			bid.AuthorID,
			bid.Version,
		)

	sql, args, err = querry.ToSql()
	if err != nil {
		return entity.Bid{}, fmt.Errorf("%s: %w", op, err)
	}

	_, err = tx.Exec(context.Background(), sql, args...)
	if err != nil {
		tx.Rollback(context.Background())
		return entity.Bid{}, fmt.Errorf("%s: %w", op, err)
	}

	q := qb.Select(
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
		Where(sq.Eq{"b.id": newBidID})

	sql, args, err = q.ToSql()
	if err != nil {
		return entity.Bid{}, fmt.Errorf("%s: %w", op, err)
	}

	row = tx.QueryRow(context.Background(), sql, args...)

	var newBid entity.Bid
	err = row.Scan(
		&newBid.ID,
		&newBid.Name,
		&newBid.Status,
		&newBid.AuthorType,
		&newBid.AuthorID,
		&newBid.Version,
		&newBid.CreatedAt,
	)
	if err != nil {
		tx.Rollback(context.Background())
		return entity.Bid{}, fmt.Errorf("%s: %w", op, err)
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return entity.Bid{}, fmt.Errorf("%s: %w", op, err)
	}

	return newBid, nil
}
