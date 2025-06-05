package service

import (
	"API_Service"
	"API_Service/internal/repository"
)

type ArticleService struct {
	repo repository.Article
}

func NewArticleService(repo repository.Article) *ArticleService {
	return &ArticleService{
		repo: repo,
	}
}

func (s *ArticleService) CreateArticle(userId int, article API_Service.Article) (int, error) {
	return s.repo.CreateArticle(userId, article)
}

func (s *ArticleService) GetAll(userId int) ([]API_Service.Article, error) {
	return s.repo.GetAll(userId)
}

func (s *ArticleService) GetArticleById(userId, articleId int) (API_Service.Article, error) {
	return s.repo.GetArticleById(userId, articleId)
}

func (s *ArticleService) DeleteArticleById(userId, articleId int) error {
	return s.repo.DeleteArticleById(userId, articleId)
}

func (s *ArticleService) UpdateArticleById(userId, articleId int, input API_Service.UpdateArticle) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.UpdateArticleById(userId, articleId, input)
}
