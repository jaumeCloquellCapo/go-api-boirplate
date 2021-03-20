package controller

import (
	errorNotFound "ApiRest/app/error"
	"ApiRest/app/model/billing"
	"ApiRest/app/service"
	"ApiRest/internal/logger"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
)

//BillingControllerInterface define the services controller interface methods
type BillingControllerInterface interface {
	AddCustomer(c *gin.Context)
}

// billingController handles communication with the external service
type billingController struct {
	service  service.BillingServiceInterface
	uservice service.UserServiceInterface
	logger   logger.Logger
}

// NewBillingController implements the user controller interface.
func NewBillingController(service service.BillingServiceInterface, uservice service.UserServiceInterface, logger logger.Logger) BillingControllerInterface {
	return &billingController{
		service,
		uservice,
		logger,
	}
}

// Store implements the method to validate the params to store a  new payment method and handle the service
func (uc *billingController) AddCustomer(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	user, err := uc.uservice.FindByID(id)
	if err != nil {
		uc.logger.Error(err.Error())
		c.Status(errorNotFound.ParseError(err))
		return
	}

	var rq billing.CreateCustomer

	if err := c.ShouldBindJSON(&rq); err != nil {
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	validate := validator.New()

	err = validate.Struct(rq)
	if err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}
	p, err := uc.service.GetPaymentAdapter(rq)
	if err != nil {
		uc.logger.Error(err.Error())
		c.Status(errorNotFound.ParseError(err))
		return
	}
	err = uc.service.AddBilling(*user, *p)

	if err != nil {
		uc.logger.Error(err.Error())
		c.Status(errorNotFound.ParseError(err))
		return
	}

	c.Status(http.StatusOK)
}
