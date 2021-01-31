package controller

import (
	errorNotFound "ApiRest/app/error"
	"ApiRest/app/model"
	"ApiRest/app/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// UserControllerInterface ...
type UserControllerInterface interface {
	FindUserById(c *gin.Context)
	RemoveUserById(c *gin.Context)
	UpdateUserById(c *gin.Context)
	FindAllUsers(c *gin.Context)
}

type userController struct {
	service service.UserServiceInterface
}

// NewUserController ...
func NewUserController(service service.UserServiceInterface) UserControllerInterface {
	return &userController{
		service,
	}
}

//FindUserById ...
func (uc *userController) FindUserById(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	user, err := uc.service.FindById(id)

	if err != nil {
		if _, ok := err.(*errorNotFound.NotFound); ok {
			c.Status(http.StatusNotFound)
			return
		}
		log.Print(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, user)
}

//RemoveUserById ...
func (uc *userController) RemoveUserById(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	err = uc.service.RemoveById(id)

	if err != nil {
		if _, ok := err.(*errorNotFound.NotFound); ok {
			c.Status(http.StatusNotFound)
			return
		}
		log.Print(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

//UpdateUserById ...
func (uc *userController) UpdateUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	var user model.UpdateUser

	if err := c.ShouldBindJSON(&user); err != nil {
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	err = uc.service.UpdateById(id, user)

	if err != nil {
		if _, ok := err.(*errorNotFound.NotFound); ok {
			c.Status(http.StatusNotFound)
			return
		}
		log.Print(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

//FindAllUsers ...
func (uc *userController) FindAllUsers(c *gin.Context) {

	user, err := uc.service.FindAll()

	if err != nil {
		log.Print(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, user)
}
