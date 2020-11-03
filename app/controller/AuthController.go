package controller

import (
	"ApiRest/app/model"
	"ApiRest/app/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthController interface {
	Login(c *gin.Context)
}

type authController struct {
	authService service.AuthService
	userService service.UserService
}

func NewAuthController(authService service.AuthService, userService service.UserService) AuthController {
	return &authController{
		authService,
		userService,
	}
}

func (h *authController)Login(c *gin.Context) {
	var u model.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	//check if the user exist:
	user, err := h.userService.GetUserByEmail(u.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}


	//since after the user logged out, we destroyed that record in the database so that same jwt token can't be used twice. We need to create the token again
	token, err := h.authService.CreateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, token)
}

