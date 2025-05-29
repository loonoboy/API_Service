package service

import "API_Service/internal/repository"

type Authorization interface {
}

type Article interface {
}

type Service struct {
	Authorization
	Article
}

func NewService(repo *repository.Repository) *Service {
	return &Service{}
}
