package service
import (
	"ApiRest/app/model"
	"ApiRest/app/repository"
)

type UserService interface {
	GetById(id int) (model.User, error)
}

type userService struct {
	userRepo repository.UserRepo
}

func NewUserService(userRepo repository.UserRepo) UserService {
	return &userService{
		userRepo,
	}
}

func (s *userService) GetById(id int) (model.User, error) {
	user, err := s.userRepo.GetById(id)
	if err != nil {
		print(err)
	}

	return user, err
}
}