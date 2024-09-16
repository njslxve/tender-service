package main

import (
	"log/slog"
	"os"

	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725721237-team-77964/zadanie-6105/internal/config"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725721237-team-77964/zadanie-6105/internal/server"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725721237-team-77964/zadanie-6105/internal/transport/storage"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725721237-team-77964/zadanie-6105/internal/usecase"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725721237-team-77964/zadanie-6105/pkg/client/postgres"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725721237-team-77964/zadanie-6105/pkg/logger"
	_ "github.com/jackc/pgx/v5/stdlib"
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

	// goose.SetBaseFS(migrations.EmbedFS)
	// db, err := sql.Open("pgx", cfg.PostgresConn)
	// if err != nil {
	// 	log.Error(err.Error())
	// 	os.Exit(1)
	// }

	// err = goose.Up(db, ".")
	// if err != nil {
	// 	log.Error(err.Error())
	// 	os.Exit(1)
	// }
	// db.Close()

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
