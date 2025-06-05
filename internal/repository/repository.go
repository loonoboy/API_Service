package repository

import (
	"API_Service"
	"API_Service/internal/repository/postgres"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user API_Service.User) (int, error)
	GetUser(username, password string) (API_Service.User, error)
}

type Article interface {
	CreateArticle(userId int, article API_Service.Article) (int, error)
	GetAll(userId int) ([]API_Service.Article, error)
	GetArticleById(userId, articleId int) (API_Service.Article, error)
	DeleteArticleById(userId, articleId int) error
	UpdateArticleById(userId, articleID int, input API_Service.UpdateArticle) error
}

type Repository struct {
	Authorization
	Article
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: postgres.NewAuthPostgres(db),
		Article:       postgres.NewArticlePostgres(db),
	}
}
