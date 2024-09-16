package storage

import (
	"context"
	"fmt"

	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725721237-team-77964/zadanie-6105/internal/entity"
	sq "github.com/Masterminds/squirrel"
)

func (s *Storage) SubmitBidDecision(bidDecision entity.BidDecision) (entity.Bid, error) {
	const op = "storage.SubmitBidDecision"

	querry := qb.Insert("bid_decisions").
		Columns("bid_id", "tender_id", "decision").
		Values(bidDecision.BidID, bidDecision.TenderID, bidDecision.Decision)

	sql, args, err := querry.ToSql()
	if err != nil {
		return entity.Bid{}, fmt.Errorf("%s: %w", op, err)
	}

	tx, err := s.db.Begin(context.Background())
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
		Where(sq.Eq{"b.id": bidDecision.BidID})

	sql, args, err = q.ToSql()
	if err != nil {
		return entity.Bid{}, fmt.Errorf("%s: %w", op, err)
	}

	var bid entity.Bid

	row := tx.QueryRow(context.Background(), sql, args...)

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

func (s *Storage) GetTenderIdByBidId(bidId string) (string, error) {
	const op = "storage.GetTenderIdByBidId"

	querry := qb.Select("tender_id").
		From("bids").
		Where(sq.Eq{"id": bidId})

	sql, args, err := querry.ToSql()
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	var tenderId string

	err = s.db.QueryRow(context.Background(), sql, args...).Scan(&tenderId)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return tenderId, nil
}
