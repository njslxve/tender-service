package storage

import (
	"context"
	"fmt"

	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725721237-team-77964/zadanie-6105/internal/entity"
	sq "github.com/Masterminds/squirrel"
)

func (s *Storage) CreateTender(tender entity.Tender) (entity.Tender, error) {
	const op = "storage.CreateTender"

	var status string

	if tender.Status == "" {
		status = "Created"
	} else {
		status = tender.Status
	}

	if tender.Version == 0 {
		tender.Version = 1
	}

	tx, err := s.db.Begin(context.Background())
	if err != nil {
		return entity.Tender{}, fmt.Errorf("%s: %w", op, err)
	}

	var newTenderID string

	querry := qb.Insert("tenders").
		Columns("latest_version").
		Values(tender.Version).
		Suffix("RETURNING id")

	sql, args, err := querry.ToSql()
	if err != nil {
		return entity.Tender{}, fmt.Errorf("%s: %w", op, err)
	}

	row := tx.QueryRow(context.Background(), sql, args...)

	err = row.Scan(&newTenderID)
	if err != nil {
		tx.Rollback(context.Background())
		return entity.Tender{}, fmt.Errorf("%s: %w", op, err)
	}

	querry = qb.Insert("tenders_versions").
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
			newTenderID,
			tender.Name,
			tender.Description,
			tender.ServiceType,
			status,
			tender.CreatorUsername,
			tender.OrganizationID,
			tender.Version,
		)

	sql, args, err = querry.ToSql()
	if err != nil {
		return entity.Tender{}, fmt.Errorf("%s: %w", op, err)
	}

	_, err = tx.Exec(context.Background(), sql, args...)
	if err != nil {
		tx.Rollback(context.Background())
		return entity.Tender{}, fmt.Errorf("%s: %w", op, err)
	}

	q := qb.Select(
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
		Where(sq.Eq{"t.id": newTenderID})

	sql, args, err = q.ToSql()
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
