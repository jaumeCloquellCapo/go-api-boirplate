package controller

import (
	errorNotFound "ApiRest/app/error"
	"ApiRest/app/middleware"
	"ApiRest/app/model"
	"ApiRest/app/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AuthController : interface AuthController
type AuthControllerInterface interface {
	Login(c *gin.Context)
	Logout(c *gin.Context)
	SignUp(c *gin.Context)
}

type authController struct {
	authService service.AuthServiceInterface
	userService service.UserServiceInterface
}

func NewAuthController(authService service.AuthServiceInterface, userService service.UserServiceInterface) AuthControllerInterface {
	return &authController{
		authService,
		userService,
	}
}

func (h *authController) Login(c *gin.Context) {

	var userLogin model.Credentials

	var err error

	if err := c.ShouldBindJSON(&userLogin); err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	//since after the user logged out, we destroyed that record in the database so that same jwt token can't be used twice. We need to create the token again
	tokenDetail, err := h.authService.Login(userLogin)

	if err != nil {
		if _, ok := err.(*errorNotFound.ErrorNotFound); ok {
			c.Status(http.StatusNotFound)
			return
		}
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, tokenDetail)
}

func (h *authController) Logout(c *gin.Context) {

	tokenAuth, err := middleware.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = h.authService.Logout(tokenAuth.AccessUUID)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	c.Writer.WriteHeader(http.StatusOK)
}

func (h *authController) SignUp(c *gin.Context) {

	var UserSignUp model.CreateUser

	if err := c.ShouldBindJSON(&UserSignUp); err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	//since after the user logged out, we destroyed that record in the database so that same jwt token can't be used twice. We need to create the token again
	_, tokenDetail, err := h.authService.SignUp(UserSignUp)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, tokenDetail)

}
