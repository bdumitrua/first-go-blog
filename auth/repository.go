package auth

import (
	"database/sql"
	"errors"
	"first-blog-api/utils"
	"fmt"
)

type Repository interface {
	Login(loginDto *LoginDto) (*User, error)
	CreateUser(newUserDto *UserCreateDTO) (string, error)
}

type repositoryImpl struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repositoryImpl{db: db}
}

func (r *repositoryImpl) Login(loginDto *LoginDto) (*User, error) {
	var user User
	var passwordHash string

	err := r.db.QueryRow("SELECT id, name, email, password FROM users WHERE email = ?", loginDto.Email).Scan(&user.ID, &user.Name, &user.Email, &passwordHash)
	if err == sql.ErrNoRows {
		return nil, errors.New("invalid credentials")
	} else if err != nil {
		return nil, err
	}

	if utils.CheckPasswordHash(loginDto.Password, passwordHash) {
		return nil, errors.New("invalid credentials")
	}

	return &user, nil
}

func (r *repositoryImpl) CreateUser(newUserDto *UserCreateDTO) (string, error) {
	dbEmail := ""
	err := r.db.QueryRow("SELECT email FROM users WHERE email = ?", newUserDto.Email).Scan(&dbEmail)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}

	if dbEmail != "" {
		return "", errors.New("email already taken")
	}

	result, err := r.db.Exec("INSERT INTO users (name, email, password) VALUES (?, ?, ?)", newUserDto.Name, newUserDto.Email, newUserDto.Password)
	if err != nil {
		return "", err
	}

	id, _ := result.LastInsertId()
	return fmt.Sprintf("User created successfully with ID %d", id), nil
}
