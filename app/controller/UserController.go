package controller

import (
	"ApiRest/app/service"
	"net/http"
)
type UserController interface{
	GetById(w http.ResponseWriter, r *http.Request)
}

type userController struct {
	service service.UserService
}

func NewUserController(service service.UserService) UserController {
	return &userController{
		service,
	}
}

func (h *userController) GetById(w http.ResponseWriter, r *http.Request){
	return
}