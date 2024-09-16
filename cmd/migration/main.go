package main

import (
	"database/sql"
	"log"

	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725721237-team-77964/zadanie-6105/internal/config"
	"git.codenrock.com/avito-testirovanie-na-backend-1270/cnrprod1725721237-team-77964/zadanie-6105/migrations"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
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
