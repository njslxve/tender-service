package main

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
	"github.com/njslxve/tender-service/internal/config"
	"github.com/njslxve/tender-service/migrations"
	"github.com/pressly/goose/v3"
)

func main() {
	err := godotenv.Load("./.env-example")
	if err != nil {
		log.Fatal(err)
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	conn, err := sql.Open("pgx", cfg.PostgresConn)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	goose.SetBaseFS(migrations.EmbedFS)

	err = goose.Up(conn, ".")
	if err != nil {
		log.Fatal(err)
	}
}
