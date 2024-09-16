package storage

import (
	"context"
	"fmt"

	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725721237-team-77964/zadanie-6105/internal/entity"
	sq "github.com/Masterminds/squirrel"
)

func (t *Storage) EditBid(bid entity.Bid) (entity.Bid, error) {
	const op = "storage.EditBid"

	tx, err := t.db.Begin(context.Background())
	if err != nil {
		return entity.Bid{}, fmt.Errorf("%s: %w", op, err)
	}

	querry := qb.Insert("bids_versions").
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
			bid.ID,
			bid.Name,
			bid.Description,
			bid.TenderID,
			bid.Status,
			bid.AuthorType,
			bid.AuthorID,
			bid.Version,
		).
		Suffix("RETURNING version")

	var vers int32

	sql, args, err := querry.ToSql()
	if err != nil {
		return entity.Bid{}, fmt.Errorf("%s: %w", op, err)
	}

	row := tx.QueryRow(context.Background(), sql, args...)

	err = row.Scan(&vers)
	if err != nil {
		tx.Rollback(context.Background())
		return entity.Bid{}, fmt.Errorf("%s: %w", op, err)
	}

	q := qb.Update("bids").
		Set("latest_version", vers).
		Where(sq.Eq{"id": bid.ID})

	sql, args, err = q.ToSql()
	if err != nil {
		return entity.Bid{}, fmt.Errorf("%s: %w", op, err)
	}

	_, err = tx.Exec(context.Background(), sql, args...)
	if err != nil {
		tx.Rollback(context.Background())
		return entity.Bid{}, fmt.Errorf("%s: %w", op, err)
	}

	qry := qb.Select(
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
		Where(sq.Eq{"b.id": bid.ID})

	sql, args, err = qry.ToSql()
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
