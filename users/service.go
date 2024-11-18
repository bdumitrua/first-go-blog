package users

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	GetById(userId int) (*User, error)
	CreateUser(createUserDto *UserCreateDTO) (string, error)
	UpdateUser(updateUserDto *UserUpdateDTO, userId int) (string, error)
}

type serviceImpl struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &serviceImpl{repo: repo}
}

func (s *serviceImpl) GetById(userId int) (*User, error) {
	return s.repo.GetById(userId)
}

func (s *serviceImpl) CreateUser(newUserDTO *UserCreateDTO) (string, error) {
	hashedPassword, err := HashPassword(newUserDTO.Password)
	if err != nil {
		return "error while processing password security", errors.New("error while preparing data")
	}

	newUserDTO.Password = hashedPassword
	return s.repo.CreateUser(newUserDTO)
}

func (s *serviceImpl) UpdateUser(updateUserDto *UserUpdateDTO, userId int) (string, error) {
	return s.repo.UpdateUser(updateUserDto, userId)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
