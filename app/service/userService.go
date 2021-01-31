package service

import (
	"ApiRest/app/model"
	"ApiRest/app/repository"
	"ApiRest/internal/helpers"
)

//UserServiceInterface ...
type UserServiceInterface interface {
	FindById(id int) (user *model.User, err error)
	RemoveById(id int) error
	UpdateById(id int, user model.UpdateUser) error
	FindAll() ([]model.User, error)
	FindByEmail(email string) (user *model.User, err error)
}

type userService struct {
	userRepo repository.UserRepositoryInterface
}

// NewUserService ...
func NewUserService(userRepo repository.UserRepositoryInterface) *userService {
	return &userService{
		userRepo,
	}
}

//FindById ...
func (s *userService) FindById(id int) (user *model.User, err error) {
	return s.userRepo.FindById(id)
}

//RemoveById ...
func (s *userService) RemoveById(id int) error {
	return s.userRepo.RemoveById(id)
}

//UpdateById ...
func (s *userService) UpdateById(id int, user model.UpdateUser) error {
	bytePassword := []byte(user.Password)
	var err error
	user.Password, err = helpers.HashAndSalt(bytePassword)
	if err != nil {
		return err
	}
	return s.userRepo.UpdateById(id, user)
}

//FindAll ...
func (s *userService) FindAll() ([]model.User, error) {
	return s.userRepo.FindAll()
}

//FindByEmail ...
func (s *userService) FindByEmail(email string) (user *model.User, err error) {
	return s.userRepo.FindByEmail(email)
}
