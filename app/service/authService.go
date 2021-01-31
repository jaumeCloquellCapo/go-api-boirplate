package service

import (
	"ApiRest/app/model"
	"ApiRest/app/repository"
	"ApiRest/helpers"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type AuthServiceInterface interface {
	Login(user model.Credentials) (tokenDetails model.TokenDetails, err error)
	Logout(accessUUID string) error
	SignUp(UserSignUp model.CreateUser) (user *model.User, tokenDetails model.TokenDetails, err error)
	GetAuth(AccessUUID string) (int64, error)
}

type authService struct {
	authRepository repository.AuthRepositoryInterface
	userRepository repository.UserRepositoryInterface
}

//NewAuthService ...
func NewAuthService(authRepository repository.AuthRepositoryInterface, userService repository.UserRepositoryInterface) AuthServiceInterface {
	return &authService{
		authRepository,
		userService,
	}
}

//Logout
func (m *authService) Logout(accessUUID string) error {
	return m.authRepository.DeleteAuth(accessUUID)
}

//GetAuth
func (m *authService) GetAuth(accessUUID string) (int64, error) {
	return m.authRepository.GetAuth(accessUUID)
}

//CreateToken ...
func (m *authService) Login(userLogin model.Credentials) (token model.TokenDetails, err error) {

	var user *model.User

	if user, err = m.userRepository.FindByEmail(userLogin.Email); err != nil {
		return token, err
	}

	//ok := helpers.ComparePasswords(userLogin.Password, *user.Password)

	// Compare the password form and database if match
	bytePassword := []byte(userLogin.Password)
	byteHashedPassword := []byte(*user.Password)
	err = bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
	if err != nil {
		fmt.Print(err.Error())
		return token, err
	}

	token, err = m.authRepository.CreateToken(*user)
	if err != nil {
		log.Println(err.Error())
		return
	}

	err = m.authRepository.CreateAuth(*user, token)
	if err != nil {
		log.Println(err.Error())
		return
	}

	return
}

//signUp ...
func (m *authService) SignUp(UserSignUp model.CreateUser) (user *model.User, token model.TokenDetails, err error) {

	bytePassword := []byte(UserSignUp.Password)

	UserSignUp.Password, err = helpers.HashAndSalt(bytePassword)
	if err != nil {
		log.Println(err.Error())
		return
	}

	user, err = m.userRepository.Create(UserSignUp)
	if err != nil {
		log.Println(err.Error())
		return
	}

	token, err = m.authRepository.CreateToken(*user)
	if err != nil {
		log.Println(err.Error())
		return
	}

	err = m.authRepository.CreateAuth(*user, token)
	if err != nil {
		log.Println(err.Error())
		return
	}
	return
}
