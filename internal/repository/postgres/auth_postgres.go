package postgres

import (
	"API_Service"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user API_Service.User) (int, error) {
	const op = "repository.auth_postgres.CreateUser"
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (API_Service.User, error) {
	const op = "repository.auth_postgres.GetUser"
	var user API_Service.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 and password_hash=$2", usersTable)
	err := r.db.Get(&user, query, username, password)
	if err != nil {
		return user, fmt.Errorf("%s: %w", op, err)
	}
	return user, err
}
