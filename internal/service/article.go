package service

import (
	"API_Service/internal/dto"
	"API_Service/internal/repository"
)

const coutArticles = 10

type ArticleService struct {
	repo repository.Article
}

func NewArticleService(repo repository.Article) *ArticleService {
	return &ArticleService{
		repo: repo,
	}
}

func (s *ArticleService) CreateArticle(userId int, article dto.Article) (int, error) {
	id, err := s.repo.ArticleDB.CreateArticle(userId, article)
	if err != nil {
		return 0, err
	}
	errCh := make(chan error, 1)
	go func() {
		errCh <- s.RefreshRecentArticles()
	}()

	if err = <-errCh; err != nil {
		return 0, err
	}

	return id, nil
}

func (s *ArticleService) GetAllById(userId int) ([]dto.Article, error) {
	return s.repo.ArticleDB.GetAllById(userId)
}

func (s *ArticleService) GetArticleById(userId, articleId int) (dto.Article, error) {
	return s.repo.ArticleDB.GetArticleById(userId, articleId)
}

func (s *ArticleService) DeleteArticleById(userId, articleId int) error {
	err := s.repo.ArticleDB.DeleteArticleById(userId, articleId)
	if err != nil {
		return err
	}
	errCh := make(chan error, 1)
	go func() {
		errCh <- s.RefreshRecentArticles()
	}()

	if err = <-errCh; err != nil {
		return err
	}

	return nil
}

func (s *ArticleService) UpdateArticleById(userId, articleId int, input dto.UpdateArticle) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.ArticleDB.UpdateArticleById(userId, articleId, input)
}

func (s *ArticleService) GetAllArticles() ([]dto.Article, error) {
	var data []dto.Article
	data, err := s.repo.ArticleRedis.GetArticles()
	if data == nil || err != nil {
		data, err = s.repo.ArticleDB.GetLastArticles(coutArticles)
		if err != nil {
			return nil, err
		}
	}
	return data, nil
}

func (s *ArticleService) WarmupRecentArticles() error {
	articles, err := s.repo.GetLastArticles(coutArticles)
	if err != nil {
		return err
	}
	return s.repo.ArticleRedis.SetRecentArticles(articles)
}

func (s *ArticleService) RefreshRecentArticles() error {
	articles, err := s.repo.GetLastArticles(coutArticles)
	if err != nil {
		return err
	}
	return s.repo.ArticleRedis.SetRecentArticles(articles)
}
