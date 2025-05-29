package postgres

import (
	"API_Service/internal/config"
	"fmt"
	"github.com/jmoiron/sqlx"
)

func NewPostgresDB(cfg config.DB) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode))
	if err != nil {
		return nil, err
	}
	return db, nil
}
