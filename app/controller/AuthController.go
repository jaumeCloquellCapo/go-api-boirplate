package controller

import (
	"ApiRest/app/model"
	"ApiRest/app/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AuthController : interface AuthController
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

func (h *authController) Login(c *gin.Context) {
	var userLogin model.UserLogin
	var err error
	var user model.User

	if err := c.ShouldBindJSON(&userLogin); err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	if user, err = h.userService.GetUserByEmail(userLogin.Email); err != nil {
		c.Writer.WriteHeader(http.StatusNotFound)
		return
	}

	//since after the user logged out, we destroyed that record in the database so that same jwt token can't be used twice. We need to create the token again
	tokenDetail, err := h.authService.LoginService(user)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, tokenDetail.Token)
}
