package storage

import (
	"context"
	"fmt"

	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725721237-team-77964/zadanie-6105/internal/entity"
	sq "github.com/Masterminds/squirrel"
)

func (s *Storage) EditTender(tender entity.Tender) (entity.Tender, error) {
	const op = "storage.EditTender"

	tx, err := s.db.Begin(context.Background())
	if err != nil {
		return entity.Tender{}, fmt.Errorf("%s: %w", op, err)
	}

	querry := qb.Insert("tenders_versions").
		Columns(
			"tender_id",
			"name",
			"description",
			"service_type",
			"status",
			"creator_username",
			"organization_id",
			"version",
		).
		Values(
			tender.ID,
			tender.Name,
			tender.Description,
			tender.ServiceType,
			tender.Status,
			tender.CreatorUsername,
			tender.OrganizationID,
			tender.Version,
		).
		Suffix("RETURNING version")

	var vers int32

	sql, args, err := querry.ToSql()
	if err != nil {
		return entity.Tender{}, fmt.Errorf("%s: %w", op, err)
	}

	row := tx.QueryRow(context.Background(), sql, args...)

	err = row.Scan(&vers)
	if err != nil {
		tx.Rollback(context.Background())
		return entity.Tender{}, fmt.Errorf("%s: %w", op, err)
	}

	q := qb.Update("tenders").
		Set("latest_version", vers).
		Where(sq.Eq{"id": tender.ID})

	sql, args, err = q.ToSql()
	if err != nil {
		return entity.Tender{}, fmt.Errorf("%s: %w", op, err)
	}

	_, err = tx.Exec(context.Background(), sql, args...)
	if err != nil {
		tx.Rollback(context.Background())
		return entity.Tender{}, fmt.Errorf("%s: %w", op, err)
	}

	qry := qb.Select(
		"t.id as tender_id",
		"tv.name",
		"tv.description",
		"tv.service_type",
		"tv.status",
		"tv.creator_username",
		"tv.organization_id",
		"tv.version",
		"tv.created_at",
	).
		From("tenders t").
		Join("tenders_versions tv ON tv.tender_id = t.id").
		Where("tv.version = t.latest_version").
		Where(sq.Eq{"t.id": tender.ID})

	sql, args, err = qry.ToSql()
	if err != nil {
		return entity.Tender{}, fmt.Errorf("%s: %w", op, err)
	}

	row = tx.QueryRow(context.Background(), sql, args...)

	var newTender entity.Tender

	err = row.Scan(&newTender.ID,
		&newTender.Name,
		&newTender.Description,
		&newTender.ServiceType,
		&newTender.Status,
		&newTender.CreatorUsername,
		&newTender.OrganizationID,
		&newTender.Version,
		&newTender.CreatedAt,
	)
	if err != nil {
		tx.Rollback(context.Background())
		return entity.Tender{}, fmt.Errorf("%s: %w", op, err)
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return entity.Tender{}, fmt.Errorf("%s: %w", op, err)
	}

	return newTender, nil
}
