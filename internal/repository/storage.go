package repository

import "errors"

var (
	ErrUsernameNotFound = errors.New("username not found")
	ErrUsernameExists   = errors.New("username exists")
	ErrArticleNotFound  = errors.New("article not found")
)
