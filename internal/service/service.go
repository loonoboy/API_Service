package service

import (
	"API_Service"
	"API_Service/internal/repository"
)

type Authorization interface {
	CreateUser(user API_Service.User) (int, error)
}

type Article interface {
}

type Service struct {
	Authorization
	Article
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
	}
}
