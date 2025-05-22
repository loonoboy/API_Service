package storage

import "errors"

var (
	ErrURLNotFound    = errors.New("url not found")
	ErrUsernameExists = errors.New("username exists")
)
