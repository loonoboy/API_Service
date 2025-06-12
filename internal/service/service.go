package service

import (
	"API_Service/internal/dto"
	"API_Service/internal/repository"
)

type Authorization interface {
	CreateUser(user dto.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Article interface {
	CreateArticle(userId int, article dto.Article) (int, error)
	GetAllById(userId int) ([]dto.Article, error)
	GetArticleById(userId, articleId int) (dto.Article, error)
	DeleteArticleById(userId, articleId int) error
	UpdateArticleById(userId, articleID int, input dto.UpdateArticle) error
	WarmupRecentArticles() error
	RefreshRecentArticles() error
	GetAllArticles() ([]dto.Article, error)
}

type Service struct {
	Authorization
	Article
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		Article:       NewArticleService(repo.Article),
	}
}
