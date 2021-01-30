package service

import (
	"ApiRest/app/model"
	"ApiRest/app/repository"
)

type UserServiceInterface interface {
	GetUserById(id int) (user *model.User, err error)
	RemoveUserById(id int) error
	UpdateUserById(id int, user model.UpdateUser) error
	GetUsers() ([]model.User, error)
	GetUserByEmail(email string) (user *model.User, err error)
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
func (s *userService) GetUserById(id int) (user *model.User, err error) {
	return s.userRepo.GetUserById(id)
}

func (s *userService) RemoveUserById(id int) error {
	return s.userRepo.RemoveUserById(id)
}

func (s *userService) UpdateUserById(id int, user model.UpdateUser) error {
	return s.userRepo.UpdateUserById(id, user)
}

//GetUsers
func (s *userService) GetUsers() ([]model.User, error) {
	return s.userRepo.GetUsers()
}

//GetUserByEmail
func (s *userService) GetUserByEmail(email string) (user *model.User, err error) {
	return s.userRepo.GetUserByEmail(email)
}
