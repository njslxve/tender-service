package main

import (
	"database/sql"
	"log/slog"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/njslxve/tender-service/internal/config"
	"github.com/njslxve/tender-service/internal/server"
	"github.com/njslxve/tender-service/internal/transport/storage"
	"github.com/njslxve/tender-service/internal/usecase"
	"github.com/njslxve/tender-service/migrations"
	"github.com/njslxve/tender-service/pkg/client/postgres"
	"github.com/njslxve/tender-service/pkg/logger"
	"github.com/pressly/goose/v3"
)

func main() {
	log := logger.New()
	slog.SetDefault(log)

	log.Info("Starting tender service")

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Error("Failed to load config")
		return
	}

	goose.SetBaseFS(migrations.EmbedFS)
	db, err := sql.Open("pgx", cfg.PostgresConn)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	err = goose.Up(db, ".")
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
	db.Close()

	client, err := postgres.NewClient(cfg)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	storage := storage.New(log, client)

	ucase := usecase.New(cfg, log, storage)

	server := server.New(cfg, log, ucase)

	server.Start()
}
