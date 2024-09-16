package storage

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/njslxve/tender-service/internal/entity"
)

func (s *Storage) SubmitBidFeedback(feedback entity.BidFeedback) (entity.Bid, error) {
	const op = "storage.SubmitBidFeedback"

	tx, err := s.db.Begin(context.Background())
	if err != nil {
		return entity.Bid{}, fmt.Errorf("%s: %w", op, err)
	}

	querry := qb.Insert("bid_feedback").
		Columns(
			"bid_id",
			"description",
			"author_id",
		).
		Values(
			feedback.BidID,
			feedback.Description,
			feedback.AuthorID,
		)

	sql, args, err := querry.ToSql()
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
		Where(sq.Eq{"b.id": feedback.BidID})

	sql, args, err = q.ToSql()
	if err != nil {
		return entity.Bid{}, fmt.Errorf("%s: %w", op, err)
	}

	row := tx.QueryRow(context.Background(), sql, args...)

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
		tx.Rollback(context.Background())
		return entity.Bid{}, fmt.Errorf("%s: %w", op, err)
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return entity.Bid{}, fmt.Errorf("%s: %w", op, err)
	}

	return bid, nil
}
