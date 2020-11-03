package service

import (
	"ApiRest/app/model"
	"ApiRest/app/repository"
)

type UserService interface {
	GetById(id int) (model.User, error)
	GetUserByEmail(email string)(model.User, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo,
	}
}

func (s *userService) GetById(id int) (model.User, error) {
	user, err := s.GetById(id)
	if err != nil {
		print(err)
	}

	return user, err
}


func (s *userService) GetUserByEmail(email string) (model.User, error) {
	user, err := s.GetUserByEmail(email)
	if err != nil {
		print(err)
	}

	return user, err
}
