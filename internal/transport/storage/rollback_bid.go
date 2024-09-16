package storage

import (
	"context"
	"fmt"

	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725721237-team-77964/zadanie-6105/internal/entity"
	sq "github.com/Masterminds/squirrel"
)

func (s *Storage) GetBidLastVersion(bidId string) (int32, error) {
	const op = "storage.GetBidLastVersion"

	querry := qb.Select("latest_version").
		From("bids").
		Where(sq.Eq{"id": bidId})

	sql, args, err := querry.ToSql()
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	var version int32

	err = s.db.QueryRow(context.Background(), sql, args...).Scan(&version)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return version, nil
}

func (s *Storage) GetBidByVersion(bidId string, version int32) (entity.Bid, error) {
	const op = "storage.GetBidByVersion"

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
		Where(sq.Eq{"bv.version": version}).
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
