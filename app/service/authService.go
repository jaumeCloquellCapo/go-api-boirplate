package service

import (
	"ApiRest/app/model"
	"ApiRest/app/repository"
)

type AuthService interface {
	CreateToken(user model.User) (token string, err error)
}

type authService struct {
	authRepository repository.AuthRepository
}

func NewAuthService(authRepository repository.AuthRepository) AuthService {
	return &authService{
		authRepository,
	}
}

func (s *authService) CreateToken(user model.User) (token string, err error) {
	token, err = s.authRepository.CreateToken(user)
	return
}
