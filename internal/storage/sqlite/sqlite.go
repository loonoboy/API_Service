package sqlite

import (
	"API_Service/internal/storage"
	"database/sql"
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
    	first_name TEXT,
    	last_name TEXT,
    	username TEXT NOT NULL UNIQUE
	);
	CREATE TABLE IF NOT EXISTS articles(
    	id INTEGER PRIMARY KEY,
    	user_id INTEGER NOT NULL,
    	title TEXT NOT NULL,
    	content TEXT,
    	FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE)
	`)
	if err != nil {
		return nil, fmt.Errorf("%s %w", op, err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) CreateUser(firsName string, lastName string, username string) error {
	const op = "storage.sqlite.CreateUser"

	stmt, err := s.db.Prepare("INSERT INTO users(first_name, last_name, username) VALUES(?, ?, ?)")
	if err != nil {
		return fmt.Errorf("%s %w", op, err)
	}
	_, err = stmt.Exec(firsName, lastName, username)
	if err != nil {
		if sqliteErr, ok := err.(sqlite3.Error); ok && sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
			return fmt.Errorf("%s %w", op, storage.ErrUsernameExists)
		}
		return fmt.Errorf("%s %w", op, err)
	}
	return nil
}
