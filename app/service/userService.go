package service

import (
	"ApiRest/app/model"
	"ApiRest/app/repository"
	"log"
)

type UserServiceInterface interface {
	GetUserById(id int) (model.User, error)
	RemoveUserById(id int) error
	GetUsers() ([]model.User, error)
	GetUserByEmail(email string) (model.User, error)
}

type userService struct {
	userRepo repository.UserRepositoryInterface
}

//NewUserService
func NewUserService(userRepo repository.UserRepositoryInterface) *userService {
	return &userService{
		userRepo,
	}
}

//GetById
func (s *userService) GetUserById(id int) (model.User, error) {
	user, err := s.userRepo.GetUserById(id)
	if err != nil {
		log.Println(err.Error())
	}

	return user, err
}

func (s *userService) RemoveUserById(id int) error {
	err := s.userRepo.RemoveUserById(id)
	if err != nil {
		log.Println(err.Error())
	}

	return err
}

//GetUsers
func (s *userService) GetUsers() ([]model.User, error) {
	users, err := s.userRepo.GetUsers()
	if err != nil {
		log.Println(err.Error())
	}

	return users, err
}

//GetUserByEmail
func (s *userService) GetUserByEmail(email string) (model.User, error) {
	user, err := s.userRepo.GetUserByEmail(email)
	if err != nil {
		log.Println(err.Error())
	}

	return user, err
}
