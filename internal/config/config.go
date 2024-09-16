package config

import (
	"fmt"
	"log"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	ServerAddr   string `env:"SERVER_ADDRESS" env-default:"0.0.0.0:8080"`
	PostgresConn string `env:"POSTGRES_CONN"`
	PostgresJDBC string `env:"POSTGRES_JDBC_URL"`
	PostgresUser string `env:"POSTGRES_USERNAME"`
	PostgresPass string `env:"POSTGRES_PASSWORD"`
	PostgresHost string `env:"POSTGRES_HOST"`
	PostgresPort string `env:"POSTGRES_PORT"`
	PostgresDB   string `env:"POSTGRES_DATABASE"`
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load("./.env-example")
	if err != nil {
		log.Fatal(err)
	}
	var cfg Config

	err = cleanenv.ReadEnv(&cfg)
	if err != nil {
		return nil, err
	}

	if cfg.PostgresJDBC == "" {
		cfg.PostgresJDBC = fmt.Sprintf("jdbc:postgresql://%s:%s/%s", cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresDB)
	}

	if cfg.PostgresConn == "" {
		cfg.PostgresConn = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", cfg.PostgresUser, cfg.PostgresPass, cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresDB)
	}

	return &cfg, nil
}
