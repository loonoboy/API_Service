package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {
}

type Article interface {
}

type Repository struct {
	Authorization
	Article
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
