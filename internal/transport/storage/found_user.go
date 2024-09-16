package storage

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

func (s *Storage) FoundUser(username string) error {
	const op = "storage.FoundUser"

	querry := qb.Select("id").
		From("employee").
		Where(sq.Eq{"username": username})

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

	if userID == "" {
		return fmt.Errorf("%s: %w", op, fmt.Errorf("user not found"))
	}

	return nil
}
