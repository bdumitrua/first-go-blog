package auth

import (
	"database/sql"
	"errors"
)

type Repository interface {
	Login(loginDto *LoginDto) (*User, error)
}

type repositoryImpl struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repositoryImpl{db: db}
}

func (r *repositoryImpl) Login(loginDto *LoginDto) (*User, error) {
	var user User

	err := r.db.QueryRow("SELECT id, name, email FROM users WHERE email = ? AND WHERE password = ?", loginDto.Email, loginDto.Password).Scan(&user.ID, &user.Name, &user.Email)
	if err == sql.ErrNoRows {
		return nil, errors.New("invalid credentials")
	} else if err != nil {
		return nil, err
	}

	return &user, nil
}
