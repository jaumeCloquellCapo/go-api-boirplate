package service

import (
	error2 "ApiRest/app/error"
	"ApiRest/app/model"
	"ApiRest/app/model/billing"
	"ApiRest/app/repository"
)

//BillingServiceInterface define the user service interface methods
type BillingServiceInterface interface {
	AddBilling(user model.User, payment billing.Payment) error
	GetPaymentAdapter(customer billing.CreateCustomer) (*billing.Payment, error)
}

// billingService handles communication with the user repository
type billingService struct {
	paymentRepo repository.BillingRepositoryInterface
}

// NewUserService implements the user service interface.
func NewBillingService(paymentRepo repository.BillingRepositoryInterface) *billingService {
	return &billingService{
		paymentRepo,
	}
}

// FindByID implements the method to store a new a user model
func (s *billingService) AddBilling(user model.User, payment billing.Payment) error {

	key, err := payment.PaymentMethod.CreateCustomer(payment.CustomerParams)
	if err != nil {
		return err
	}

	return s.paymentRepo.CreateBillingService(payment.Identify, key, user.ID)
}


// FindByID implements the method to store a new a user model
func (s *billingService) GetPaymentAdapter(customer billing.CreateCustomer) (*billing.Payment, error) {
	p, err := billing.GetPaymentAdapter(customer.Identify)

	if err != nil {
		return nil, error2.InvalidPaymentMethod
	}

	return &billing.Payment{
		Identify:       customer.Identify,
		CustomerParams: customer.CustomerParams,
		PaymentMethod:  p,
	}, err
}
