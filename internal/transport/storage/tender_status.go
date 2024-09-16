package storage

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/njslxve/tender-service/internal/entity"
)

func (s *Storage) TenderStatus(tenderId string, username string) (string, error) {
	const op = "storage.TenderStatus"

	var status string

	fmt.Println(tenderId, username)

	querry := qb.Select("tv.status").
		From("tenders_versions tv").
		Join("tenders t ON tv.tender_id = t.id").
		Where(sq.Eq{"tv.tender_id": tenderId}).
		Where(sq.Eq{"tv.creator_username": username}).
		Where("tv.version = t.latest_version")

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

func (s *Storage) UpdateTenderStatus(tenderId string, username string, status string) (entity.Tender, error) {
	const op = "storage.UpdateTenderStatus"

	querry := qb.Update("tenders_versions").
		Set("status", status).
		Where(sq.Eq{"tender_id": tenderId}).
		Where(sq.Eq{"creator_username": username}).
		Where("version = (SELECT latest_version FROM tenders WHERE id = $4)", tenderId).
		Suffix("RETURNING tender_id, name, description, service_type, status, creator_username, organization_id, version, created_at")

	sql, args, err := querry.ToSql()
	if err != nil {
		fmt.Println("1", err)
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
