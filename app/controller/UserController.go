package controller

import (
	"ApiRest/app/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, user)
}

func (uc *userController) GetUsers(c *gin.Context) {

	user, err := uc.service.GetUsers()
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, user)
}
