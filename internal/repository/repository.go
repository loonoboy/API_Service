package repository

import (
	"API_Service/internal/dto"
	"API_Service/internal/repository/postgres"
	"API_Service/internal/repository/redisDB"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type Authorization interface {
	CreateUser(user dto.User) (int, error)
	GetUser(username, password string) (dto.User, error)
}

type ArticleDB interface {
	CreateArticle(userId int, article dto.Article) (int, error)
	GetAllById(userId int) ([]dto.Article, error)
	GetArticleById(userId, articleId int) (dto.Article, error)
	DeleteArticleById(userId, articleId int) error
	UpdateArticleById(userId, articleID int, input dto.UpdateArticle) error
	GetLastArticles(count int) ([]dto.Article, error)
}

type ArticleRedis interface {
	GetArticles() ([]dto.Article, error)
	SetRecentArticles(articles []dto.Article) error
}

type Article struct {
	ArticleDB
	ArticleRedis
}

type Repository struct {
	Authorization
	Article
}

func NewRepository(db *sqlx.DB, rdb *redis.Client) *Repository {
	return &Repository{
		Authorization: postgres.NewAuthPostgres(db),
		Article: Article{
			ArticleDB:    postgres.NewArticlePostgres(db),
			ArticleRedis: redisDB.NewArticleRedis(rdb),
		},
	}
}
