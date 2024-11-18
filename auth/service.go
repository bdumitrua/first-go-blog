package auth

import (
	"errors"
	"first-blog-api/users"
	"first-blog-api/utils"
)

type Service interface {
	Login(loginDto *LoginDto) (*User, error)
	Register(dto *users.UserCreateDTO) (string, error)
	Logout(token string) (*User, error)
	Refresh(token string) (*User, error)
}

type serviceImpl struct {
	repo        Repository
	userService users.Service
}

func NewService(repo Repository) Service {
	return &serviceImpl{repo: repo}
}

func (s *serviceImpl) Login(loginDto *LoginDto) (*User, error) {
	hashedPassword, err := utils.HashPassword(loginDto.Password)
	if err != nil {
		return nil, errors.New("error while preparing data")
	}

	loginDto.Password = hashedPassword
	return s.repo.Login(loginDto)
}

func (s *serviceImpl) Register(dto *users.UserCreateDTO) (string, error) {
	return s.userService.CreateUser(dto)
}

func (s *serviceImpl) Logout(token string) (*User, error) {
	return s.repo.Logout(token)
}

func (s *serviceImpl) Refresh(token string) (*User, error) {
	return s.repo.Refresh(token)
}
