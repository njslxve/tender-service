package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/njslxve/tender-service/internal/config"
)

func NewClient(cfg *config.Config) (*pgx.Conn, error) {
	db, err := pgx.Connect(context.Background(), cfg.PostgresConn)
	if err != nil {
		return nil, err
	}

	err = db.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	return db, nil
}
