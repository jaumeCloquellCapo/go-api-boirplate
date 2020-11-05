package controller

import (
	"ApiRest/app/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// UserController : interface
type UserController interface {
	GetUserById(c *gin.Context)
}

type userController struct {
	service service.UserService
}

func NewUserController(service service.UserService) UserController {
	return &userController{
		service,
	}
}

//GetById
func (h *userController) GetUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	user, err := h.service.GetUserById(id)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, user)
	return
}
