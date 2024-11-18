package auth

import (
	"errors"
	"first-blog-api/users"
	"first-blog-api/utils"
)

type Service interface {
	Login(loginDto *LoginDto) (string, error)
	Register(dto *users.UserCreateDTO) (string, error)
	Refresh(token string) (string, error)
}

type serviceImpl struct {
	repo        Repository
	userService users.Service
}

func NewService(repo Repository) Service {
	return &serviceImpl{repo: repo}
}

func (s *serviceImpl) Login(loginDto *LoginDto) (string, error) {
	hashedPassword, err := utils.HashPassword(loginDto.Password)
	if err != nil {
		return "", errors.New("error while preparing data")
	}

	loginDto.Password = hashedPassword

	user, err := s.repo.Login(loginDto)
	if err != nil {
		return "", err
	}

	newToken, err := GenerateJWT(user.ID)
	if err != nil {
		return "", err
	}

	return newToken, nil
}

func (s *serviceImpl) Register(dto *users.UserCreateDTO) (string, error) {
	return s.userService.CreateUser(dto)
}

func (s *serviceImpl) Refresh(token string) (string, error) {
	claims, err := ValidateJWT(token)
	if err != nil {
		return "", err
	}

	newToken, err := GenerateJWT(claims.UserId)
	if err != nil {
		return "", err
	}

	return newToken, nil
}
