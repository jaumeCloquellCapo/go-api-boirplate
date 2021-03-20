package controller

import (
	errorNotFound "ApiRest/app/error"
	"ApiRest/app/model"
	"ApiRest/app/service"
	"ApiRest/internal/logger"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
)

//UserControllerInterface define the user controller interface methods
type UserControllerInterface interface {
	Find(c *gin.Context)
	Destroy(c *gin.Context)
	Update(c *gin.Context)
	Store(c *gin.Context)
}

// userController handles communication with the user service
type userController struct {
	service service.UserServiceInterface
	logger  logger.Logger
}

// NewUserController implements the user controller interface.
func NewUserController(service service.UserServiceInterface, logger logger.Logger) UserControllerInterface {
	return &userController{
		service,
		logger,
	}
}

// Find implements the method to handle the service to find a user by the primary key
func (uc *userController) Find(c *gin.Context) {

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
		uc.logger.Error(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, user)
}

// Destroy implements the method to validate the params to store a  new user and handle the service
func (uc *userController) Destroy(c *gin.Context) {

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
		uc.logger.Error(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

// Update implements the method to validate teh params to update a user and handle the service
func (uc *userController) Update(c *gin.Context) {
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

	validate := validator.New()
	err = validate.Struct(user)
	if err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = uc.service.UpdateById(id, user)

	if err != nil {
		if _, ok := err.(*errorNotFound.NotFound); ok {
			c.Status(http.StatusNotFound)
			return
		}
		uc.logger.Error(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

// Store implements the method to validate the params to store a  new user and handle the service
func (uc *userController) Store(c *gin.Context) {

	var rq model.CreateUser

	if err := c.ShouldBindJSON(&rq); err != nil {
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	validate := validator.New()
	err := validate.Struct(rq)
	if err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = uc.service.Store(rq)

	if err != nil {
		if _, ok := err.(*errorNotFound.NotFound); ok {
			c.Status(http.StatusNotFound)
			return
		}
		uc.logger.Error(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}
