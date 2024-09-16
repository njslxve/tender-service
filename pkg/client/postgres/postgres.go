package postgres

import (
	"context"

	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725721237-team-77964/zadanie-6105/internal/config"
	"github.com/jackc/pgx/v5"
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
