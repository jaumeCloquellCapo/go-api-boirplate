package service

import (
	"ApiRest/app/model"
	"ApiRest/app/repository"
	"log"
)

type AuthServiceInterface interface {
	LoginService(user model.User) (tokenDetails model.TokenDetails, err error)
}

type authService struct {
	authRepository repository.AuthRepositoryInterface
}

func NewAuthService(authRepository repository.AuthRepositoryInterface) AuthServiceInterface {
	return &authService{
		authRepository,
	}
}

//CreateToken ...
func (m *authService) LoginService(user model.User) (token model.TokenDetails, err error) {

	//Compare the password form and database if match
	//bytePassword := []byte(form.Password)
	//byteHashedPassword := []byte(user.Password)
	//err = bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
	//if err != nil {
	//	return user, token, errors.New("Invalid password")
	//}

	token, err = m.authRepository.CreateToken(user)
	if err != nil {
		log.Println(err.Error())
		return
	}
	err = m.authRepository.CreateAuth(user, token)
	if err != nil {
		log.Println(err.Error())
		return
	}

	return
}
