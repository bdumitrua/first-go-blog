package users

import (
	"database/sql"
	"errors"
	"fmt"
)

type Repository interface {
	GetById(userId int) (*User, error)
	CreateUser(newUserDto *UserCreateDTO) (string, error)
	UpdateUser(updateUserDto *UserUpdateDTO, userId int) (string, error)
}

type repositoryImpl struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repositoryImpl{db: db}
}

func (r *repositoryImpl) GetById(userId int) (*User, error) {
	var user User

	err := r.db.QueryRow("SELECT id, name, email FROM users WHERE id = ?", userId).Scan(&user.ID, &user.Name, &user.Email)
	if err == sql.ErrNoRows {
		return nil, errors.New("User not found")
	} else if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *repositoryImpl) CreateUser(newUserDto *UserCreateDTO) (string, error) {
	result, err := r.db.Exec("INSERT INTO users (name, email, password) VALUES (?, ?, ?)", newUserDto.Name, newUserDto.Email, newUserDto.Password)
	if err != nil {
		return "", err
	}

	id, _ := result.LastInsertId()
	return fmt.Sprintf("User created successfully with ID %d", id), nil
}

func (r *repositoryImpl) UpdateUser(updateUserDto *UserUpdateDTO, userId int) (string, error) {
	result, err := r.db.Exec("UPDATE users SET name = ? WHERE id = ?", updateUserDto.Name, userId)
	if err != nil {
		return "", err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return "", errors.New("User not found")
	}

	return "User updated successfully", nil
}
