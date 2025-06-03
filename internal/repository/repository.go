package repository

import (
	"API_Service"
	"API_Service/internal/repository/postgres"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user API_Service.User) (int, error)
}

type Article interface {
}

type Repository struct {
	Authorization
	Article
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: postgres.NewAuthPostgres(db),
	}
}
