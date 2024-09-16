package storage

import (
	"context"
	"fmt"

	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725721237-team-77964/zadanie-6105/internal/entity"
	sq "github.com/Masterminds/squirrel"
)

func (s *Storage) BidStatus(bidId string, username string) (string, error) {
	const op = "storage.BidStatus"

	var status string

	querry := qb.Select("bv.status").
		From("bids_versions bv").
		Join("bids b ON bv.bid_id = b.id").
		Where(sq.Eq{"bv.bid_id": bidId}).
		Where(sq.Eq{"bv.creator_username": username}).
		Where("bv.version = b.latest_version")

	sql, args, err := querry.ToSql()
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	row := s.db.QueryRow(context.Background(), sql, args...)

	err = row.Scan(&status)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return status, nil
}

func (s *Storage) UpdateBidStatus(bidId string, username string, status string) (entity.Bid, error) {
	const op = "storage.UpdateBidStatus"

	querry := qb.Update("bids_versions").
		Set("status", status).
		Where(sq.Eq{"bid_id": bidId}).
		Where(sq.Eq{"creator_username": username}).
		Where("version = (SELECT latest_version FROM bids WHERE id = $4)", bidId).
		Suffix("RETURNING id, name, status, author_type, author_id, version, created_at")

	sql, args, err := querry.ToSql()
	if err != nil {
		fmt.Println("1", err)
		return entity.Bid{}, fmt.Errorf("%s: %w", op, err)
	}

	row := s.db.QueryRow(context.Background(), sql, args...)

	var bid entity.Bid

	err = row.Scan(&bid.ID,
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
