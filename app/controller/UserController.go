package controller

import (
	"ApiRest/app/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserController interface {
	GetById(c *gin.Context)
}

type userController struct {
	service service.UserService
}

func NewUserController(service service.UserService) UserController {
	return &userController{
		service,
	}
}

func (h *userController) GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, "")
	}

	_, err = h.service.GetById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
	}
	c.JSON(http.StatusBadRequest, "")
	return
}


