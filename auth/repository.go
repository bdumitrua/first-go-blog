package auth

import (
	"database/sql"
)

type Repository interface {
	Login(loginDto *LoginDto) (*User, error)
	Logout(token string) (*User, error)
	Refresh(token string) (*User, error)
}

type repositoryImpl struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repositoryImpl{db: db}
}

func (r *repositoryImpl) Login(loginDto *LoginDto) (*User, error) {
	return nil, nil
}

func (r *repositoryImpl) Logout(token string) (*User, error) {
	return nil, nil
}

func (r *repositoryImpl) Refresh(token string) (*User, error) {
	return nil, nil
}
