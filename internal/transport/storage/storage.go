package storage

import (
	"log/slog"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

type Storage struct {
	logger *slog.Logger
	db     *pgx.Conn
}

func New(logger *slog.Logger, db *pgx.Conn) *Storage {
	return &Storage{
		logger: logger,
		db:     db,
	}
}

var (
	qb = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
)
