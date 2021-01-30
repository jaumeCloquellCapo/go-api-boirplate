package service

import (
	"ApiRest/app/model"
	"ApiRest/app/repository"
	"ApiRest/helpers"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type AuthServiceInterface interface {
	LoginService(user model.UserLogin) (tokenDetails model.TokenDetails, err error)
	SignUp(UserSignUp model.UserSignUp) (user model.User, tokenDetails model.TokenDetails, err error)
}

type authService struct {
	authRepository repository.AuthRepositoryInterface
	userRepository repository.UserRepositoryInterface
}

func NewAuthService(authRepository repository.AuthRepositoryInterface, userService repository.UserRepositoryInterface) AuthServiceInterface {
	return &authService{
		authRepository,
		userService,
	}
}

//CreateToken ...
func (m *authService) LoginService(userLogin model.UserLogin) (token model.TokenDetails, err error) {

	var user model.User

	if user, err = m.userRepository.GetUserByEmail(userLogin.Email); err != nil {
		return
	}

	//Compare the password form and database if match
	bytePassword := []byte(userLogin.Password)
	byteHashedPassword := []byte(user.Password)

	err = bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
	if err != nil {
		return token, err
	}

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

//signUp ...
func (m *authService) SignUp(UserSignUp model.UserSignUp) (user model.User, token model.TokenDetails, err error) {

	bytePassword := []byte(UserSignUp.Password)

	UserSignUp.Password, err = helpers.HashAndSalt(bytePassword)
	if err != nil {
		log.Println(err.Error())
		return
	}

	user, err = m.userRepository.CreateUser(UserSignUp)
	if err != nil {
		log.Println(err.Error())
		return
	}

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
