package controller

import (
	"ApiRest/app/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	errorNotFound "ApiRest/app/error"
)

// UserController : interface
type UserControllerInterface interface {
	GetUserById(c *gin.Context)
	GetUsers(c *gin.Context)
}

type userController struct {
	service service.UserServiceInterface
}

func NewUserController(service service.UserServiceInterface) UserControllerInterface {
	return &userController{
		service,
	}
}

//GetById
func (uc *userController) GetUserById(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	user, err := uc.service.GetUserById(id)
	if err, ok := err.(errorNotFound.IErrorNotFound); ok && err.IsNotFound() {
		c.Status(http.StatusNotFound)
	}

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, user)
}

func (uc *userController) GetUsers(c *gin.Context) {

	user, err := uc.service.GetUsers()

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, user)
}
