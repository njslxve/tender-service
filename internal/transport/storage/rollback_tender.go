package storage

import (
	"context"
	"fmt"

	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725721237-team-77964/zadanie-6105/internal/entity"
	sq "github.com/Masterminds/squirrel"
)

func (s *Storage) GetTenderLastVersion(tenderId string) (int32, error) {
	const op = "storage.GetTenderLastVersion"

	querry := qb.Select("latest_version").
		From("tenders").
		Where(sq.Eq{"id": tenderId})

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

func (s *Storage) GetTenderByVersion(tenderId string, version int32) (entity.Tender, error) {
	const op = "storage.GetTenderByVersion"

	querry := qb.Select(
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
		Join("tenders_versions tv ON t.id = tv.tender_id").
		Where(sq.Eq{"tv.version": version}).
		Where(sq.Eq{"t.id": tenderId})

	sql, args, err := querry.ToSql()
	if err != nil {
		return entity.Tender{}, fmt.Errorf("%s: %w", op, err)
	}

	row := s.db.QueryRow(context.Background(), sql, args...)

	var tender entity.Tender

	err = row.Scan(&tender.ID,
		&tender.Name,
		&tender.Description,
		&tender.ServiceType,
		&tender.Status,
		&tender.CreatorUsername,
		&tender.OrganizationID,
		&tender.Version,
		&tender.CreatedAt,
	)
	if err != nil {
		return entity.Tender{}, fmt.Errorf("%s: %w", op, err)
	}

	return tender, nil
}
