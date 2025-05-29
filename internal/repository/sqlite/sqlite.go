package sqlite

import (
	"API_Service/internal/repository"
	"database/sql"
	"errors"
	"fmt"
	"github.com/mattn/go-sqlite3"
	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	db *sql.DB
}

func New(storagePath string) (*Storage, error) {
	const op = "storage.sqlite.New"

	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s %w", op, err)
	}
	_, err = db.Exec(`PRAGMA foreign_keys = ON`)
	if err != nil {
		return nil, fmt.Errorf("%s %w", op, err)
	}
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS users(
		id INTEGER PRIMARY KEY,
    	username TEXT NOT NULL UNIQUE,
    	password TEXT
	);
	CREATE TABLE IF NOT EXISTS articles(
    	id INTEGER PRIMARY KEY,
    	user_id INTEGER NOT NULL UNIQUE,
    	title TEXT NOT NULL,
    	content TEXT NOT NULL,
    	FOREIGN KEY (user_id) REFERENCES users(username) ON DELETE CASCADE)
	`)
	if err != nil {
		return nil, fmt.Errorf("%s %w", op, err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) CreateUser(username, password string) error {
	const op = "storage.sqlite.CreateUser"

	stmt, err := s.db.Prepare("INSERT INTO users(username, password) VALUES(?, ?)")
	if err != nil {
		return fmt.Errorf("%s: prepare statement: %w", op, err)
	}
	_, err = stmt.Exec(username, password)
	if err != nil {
		if sqliteErr, ok := err.(sqlite3.Error); ok && sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
			return fmt.Errorf("%s %w", op, repository.ErrUsernameExists)
		}
		return fmt.Errorf("%s %w", op, err)
	}
	return nil
}

func (s *Storage) SaveArticle(userId, title, content string) error {
	const op = "storage.sqlite.CreateArticle"

	stmt, err := s.db.Prepare("INSERT INTO articles(user_id, title, content) VALUES(?, ?, ?)")
	if err != nil {
		return fmt.Errorf("%s: prepare statement: %w", op, err)
	}
	_, err = stmt.Exec(userId, title, content)
	if err != nil {
		if sqliteErr, ok := err.(sqlite3.Error); ok && sqliteErr.ExtendedCode == sqlite3.ErrConstraintForeignKey {
			return fmt.Errorf("%s %w", op, repository.ErrUsernameNotFound)
		}
		return fmt.Errorf("%s %w", op, err)
	}
	return nil
}

func (s *Storage) GetArticles(userId string) (string, string, error) {
	const op = "storage.sqlite.GetArticles"

	stmt, err := s.db.Prepare("SELECT title, content FROM articles WHERE user_id = ?")
	if err != nil {
		return "", "", fmt.Errorf("%s: prepare statement: %w", op, err)
	}

	var title, content string
	err = stmt.QueryRow(userId).Scan(&title, &content)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", "", repository.ErrArticleNotFound
		}
		return "", "", fmt.Errorf("%s: execute statement: %w", op, err)
	}
	return title, content, nil
}

// TODO: emplement metod
// func (s *Storage) DeleteArticle(user_id, title string) error
