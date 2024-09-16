package storage

import (
	"context"
	"fmt"
	"strconv"

	sq "github.com/Masterminds/squirrel"
	"github.com/njslxve/tender-service/internal/entity"
)

func (s *Storage) GetUserTenders(username string, limit string, offset string) ([]entity.Tender, error) {
	const op = "storage.GetUserTenders"

	var lim, off int

	if limit != "" {
		l, err := strconv.Atoi(limit)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		lim = l
	}

	if offset != "" {
		o, err := strconv.Atoi(offset)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		off = o
	}

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
		Where("tv.version = t.latest_version").
		Where(sq.Eq{"tv.creator_username": username}).
		OrderBy("tv.name")

	if limit != "" {
		querry = querry.Limit(uint64(lim))
	}

	if offset != "" {
		querry = querry.Offset(uint64(off))
	}

	sql, args, err := querry.ToSql()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	rows, err := s.db.Query(context.Background(), sql, args...)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	tenders := make([]entity.Tender, 0)

	for rows.Next() {
		var t entity.Tender

		err = rows.Scan(&t.ID,
			&t.Name,
			&t.Description,
			&t.ServiceType,
			&t.Status,
			&t.CreatorUsername,
			&t.OrganizationID,
			&t.Version,
			&t.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		tenders = append(tenders, t)
	}

	return tenders, nil
}
