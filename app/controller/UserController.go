package controller

import (
	errorNotFound "ApiRest/app/error"
	"ApiRest/app/model"
	"ApiRest/app/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// UserController : interface
type UserControllerInterface interface {
	GetUserById(c *gin.Context)
	RemoveUserById(c *gin.Context)
	UpdateUserById(c *gin.Context)
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
		if _, ok := err.(*errorNotFound.ErrorNotFound); ok {
			c.Status(http.StatusNotFound)
			return
		}
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (uc *userController) RemoveUserById(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	err = uc.service.RemoveUserById(id)

	if err != nil {
		if _, ok := err.(*errorNotFound.ErrorNotFound); ok {
			c.Status(http.StatusNotFound)
			return
		}
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, "")
}

func (uc *userController) UpdateUserById(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	var user model.UpdateUser

	if err := c.ShouldBindJSON(&user); err != nil {
		c.Writer.WriteHeader(http.StatusNotAcceptable)
		return
	}

	err = uc.service.UpdateUserById(id, user)

	if err != nil {
		if _, ok := err.(*errorNotFound.ErrorNotFound); ok {
			c.Status(http.StatusNotFound)
			return
		}
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, "")
}

func (uc *userController) GetUsers(c *gin.Context) {

	user, err := uc.service.GetUsers()

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, user)
}
