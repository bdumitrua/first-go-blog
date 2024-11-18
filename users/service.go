package users

import (
	"errors"
	"first-blog-api/utils"
)

type Service interface {
	GetById(userId int) (*User, error)
	CreateUser(userCreateDto *UserCreateDTO) (string, error)
	UpdateUser(userUpdateDto *UserUpdateDTO, userId int) (string, error)
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

func (s *serviceImpl) CreateUser(userCreateDto *UserCreateDTO) (string, error) {
	hashedPassword, err := utils.HashPassword(userCreateDto.Password)
	if err != nil {
		return "error while processing password security", errors.New("error while preparing data")
	}

	userCreateDto.Password = hashedPassword
	return s.repo.CreateUser(userCreateDto)
}

func (s *serviceImpl) UpdateUser(userUpdateDto *UserUpdateDTO, userId int) (string, error) {
	return s.repo.UpdateUser(userUpdateDto, userId)
}
