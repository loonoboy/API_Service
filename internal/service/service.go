package service

import (
	"API_Service"
	"API_Service/internal/repository"
)

type Authorization interface {
	CreateUser(user API_Service.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Article interface {
	CreateArticle(userId int, article API_Service.Article) (int, error)
	GetAll(userId int) ([]API_Service.Article, error)
	GetArticleById(userId, articleId int) (API_Service.Article, error)
	DeleteArticleById(userId, articleId int) error
	UpdateArticleById(userId, articleID int, input API_Service.UpdateArticle) error
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
