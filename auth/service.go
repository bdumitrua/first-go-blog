package auth

import (
	"errors"
	"first-blog-api/utils"
)

type Service interface {
	Login(loginDto *LoginDto) (string, error)
	Register(dto *UserCreateDTO) (string, error)
	Refresh(token string) (string, error)
}

type serviceImpl struct {
	repo Repository
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

func (s *serviceImpl) Register(userCreateDto *UserCreateDTO) (string, error) {
	hashedPassword, err := utils.HashPassword(userCreateDto.Password)
	if err != nil {
		return "error while processing password security", errors.New("error while preparing data")
	}

	userCreateDto.Password = hashedPassword
	return s.repo.CreateUser(userCreateDto)
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
