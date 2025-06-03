package postgres

import (
	"API_Service/internal/config"
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	usersTable    = "users"
	articlesTable = "articles"
)

func NewPostgresDB(cfg config.DB) (*sqlx.DB, error) {
	const op = "repository.postgres.NewPostgresDB"
	db, err := sqlx.Connect("postgres", fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode))
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return db, nil
}
