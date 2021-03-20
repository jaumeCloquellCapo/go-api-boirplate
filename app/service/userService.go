package service

import (
	"ApiRest/app/model"
	"ApiRest/app/repository"
)

//UserServiceInterface define the user service interface methods
type UserServiceInterface interface {
	FindById(id int) (user *model.User, err error)
	RemoveById(id int) error
	UpdateById(id int, user model.UpdateUser) error
	Store(user model.CreateUser) (*model.User, error)
}

// userService handles communication with the user repository
type userService struct {
	userRepo repository.UserRepositoryInterface
}

// NewUserService implements the user service interface.
func NewUserService(userRepo repository.UserRepositoryInterface) *userService {
	return &userService{
		userRepo,
	}
}

// FindById implements the method to find a user model by primary key
func (s *userService) FindById(id int) (user *model.User, err error) {
	return s.userRepo.FindById(id)
}

// FindById implements the method to remove a user model by primary key
func (s *userService) RemoveById(id int) error {
	return s.userRepo.RemoveById(id)
}

// FindById implements the method to update a user model by primary key
func (s *userService) UpdateById(id int, user model.UpdateUser) error {
	return s.userRepo.UpdateById(id, user)
}

// FindById implements the method to store a new a user model
func (s *userService) Store(user model.CreateUser) (*model.User, error) {
	return s.userRepo.Create(user)
}
