package repository

import (
	"database/sql"
	"user-service/internal/model"
)

type UserRepository struct {
	DB *sql.DB
}

func (r *UserRepository) CreateUser(user model.User) error {
	query := "INSERT INTO users (email, password) VALUES ($1, $2)"
	_, err := r.DB.Exec(query, user.Email, user.Password)
	return err
}

func (r *UserRepository) GetUserByEmail(email string) (*model.User, error) {
	query := "SELECT id, email, password FROM users WHERE email=$1"
	row := r.DB.QueryRow(query, email)

	var user model.User
	err := row.Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
