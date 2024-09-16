package storage

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

func (s *Storage) IsResponsible(user string, org string) error {
	const op = "storage.IsResponsible"

	querry := qb.Select("id").
		From("employee").
		Where(sq.Eq{"username": user})

	sql, args, err := querry.ToSql()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	var userID string

	err = s.db.QueryRow(context.Background(), sql, args...).Scan(&userID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return fmt.Errorf("%s: %w", op, err)
		}

		return fmt.Errorf("%s: %w", op, err)
	}

	querry = qb.Select("id").
		From("organization_responsible").
		Where(sq.Eq{"user_id": userID}).
		Where(sq.Eq{"organization_id": org})

	sql, args, err = querry.ToSql()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	var id string

	err = s.db.QueryRow(context.Background(), sql, args...).Scan(&id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if id == "" {
		return fmt.Errorf("%s: %w", op, fmt.Errorf("not responsible")) // TODO: add error message
	}

	return nil
}
