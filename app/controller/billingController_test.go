package controller

import (
	error2 "ApiRest/app/error"
	"ApiRest/app/model"
	"ApiRest/app/model/billing"
	"ApiRest/app/service"
	"ApiRest/internal/logger"
	"ApiRest/mock"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)
func TestBillingController_Store(t *testing.T) {

	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userUC := mock.NewMockUserServiceCase(ctrl)
	billUC := mock.NewMockBillingServiceCase(ctrl)

	// payUc := mock.NewMockPaymentAdapterCase(ctrl)

	apiLogger := logger.NewAPILogger()
	apiLogger.InitLogger()

	billingController := NewBillingController(billUC, userUC, apiLogger)

	t.Run("UserNotFound", func(t *testing.T) {
		userUC.EXPECT().FindById(1).Return(nil, error2.ErrNotFound)

		router := gin.Default()
		router.POST("/api/users/:id/paypal", billingController.AddCustomer)
		ts := httptest.NewServer(router)
		defer ts.Close()
		w := httptest.NewRecorder()


		var customer = billing.CreateCustomer{
			Identify:       billing.AccountStripe,
			CustomerParams:  billing.CustomerParams{
				Email: "test3@test.com",
				Desc:  "a 3rd test customer",
				Card: &billing.CardParams{
					Name:     "serRes.Name",
					Number:   "userRes.Cif",
					ExpYear:  time.Now().Year() + 1,
					ExpMonth: 1,
				},
			},
		}

		body, _ := json.Marshal(customer)
		req := httptest.NewRequest("POST", "/api/users/1/paypal", bytes.NewBuffer(body))

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}


func TestBillingController_Store2(t *testing.T) {

	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userUC := mock.NewMockUserServiceCase(ctrl)
	billUC := mock.NewMockBillingServiceCase(ctrl)

	// payUc := mock.NewMockPaymentAdapterCase(ctrl)

	apiLogger := logger.NewAPILogger()
	apiLogger.InitLogger()

	billingController := NewBillingController(billUC, userUC, apiLogger)

	t.Run("InvalidIdentifyPayment", func(t *testing.T) {


		userExpected := model.User{
			ID:         1,
			Name:       "a",
			Cif:        "a",
			Country:    "a",
			PostalCode: "a",
		}

		userRes := &model.User{
			ID:         userExpected.ID,
			Name:       userExpected.Name,
			Cif:        userExpected.Cif,
			Country:    userExpected.Country,
			PostalCode: userExpected.PostalCode,
		}

		userUC.EXPECT().FindById(1).Return(userRes, nil)

		const badIdentify billing.Identify = "strbadIdentifyipe"

		var customer = billing.CreateCustomer{
			Identify:       badIdentify,
			CustomerParams:  billing.CustomerParams{
				Email: "11@test.com",
				Desc:  "a 3rd test customer",
				Card: &billing.CardParams{
					Name:     userRes.Name,
					Number:   userRes.Cif,
					ExpYear:  time.Now().Year() + 1,
					ExpMonth: 1,
				},
			},
		}

		billUC.EXPECT().GetPaymentAdapter(customer).Return(nil, error2.InvalidPaymentMethod)


		router := gin.Default()
		router.POST("/api/users/:id/paypal", billingController.AddCustomer)

		ts := httptest.NewServer(router)
		defer ts.Close()

		w := httptest.NewRecorder()
		body, _ := json.Marshal(customer)
		req := httptest.NewRequest("POST", "/api/users/1/paypal", bytes.NewBuffer(body))

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusServiceUnavailable, w.Code)
	})
}

func TestBillingController_Store3(t *testing.T) {

	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userUC := mock.NewMockUserServiceCase(ctrl)
	billUC := mock.NewMockBillingServiceCase(ctrl)

	// payUc := mock.NewMockPaymentAdapterCase(ctrl)

	apiLogger := logger.NewAPILogger()
	apiLogger.InitLogger()

	billingController := NewBillingController(billUC, userUC, apiLogger)

	t.Run("Correct", func(t *testing.T) {
		userRes := &model.User{
			ID:         1,
			Name:       "a",
			Cif:        "a",
			Country:    "a",
			PostalCode: "a",
		}

		userUC.EXPECT().FindById(1).Return(userRes, nil)


		var customer = billing.CreateCustomer{
			Identify: billing.AccountStripe,
			CustomerParams: billing.CustomerParams{
				Email: "test3@test.com",
				Desc:  "a 3rd test customer",
				Card: &billing.CardParams{
					Name:     userRes.Name,
					Number:   userRes.Cif,
					ExpYear:  time.Now().Year() + 1,
					ExpMonth: 1,
				},
			},
		}

		p := &billing.Payment{
			Identify:      customer.Identify,
			CustomerParams: customer.CustomerParams,
			PaymentMethod: &mock.FakeAdapter{},
		}

		billUC.EXPECT().GetPaymentAdapter(customer).Return(p, nil)
		billUC.EXPECT().AddBilling(*userRes, *p).Return(nil)

		router := gin.Default()
		router.POST("/api/users/:id/paypal", billingController.AddCustomer)
		ts := httptest.NewServer(router)
		defer ts.Close()
		w := httptest.NewRecorder()

		body, _ := json.Marshal(customer)
		req := httptest.NewRequest("POST", "/api/users/1/paypal", bytes.NewBuffer(body))

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})
}



func TestNewBillingController(t *testing.T) {
	type args struct {
		uservice service.UserServiceInterface
		service  service.BillingServiceInterface
		logger   logger.Logger
	}
	tests := []struct {
		name string
		args args
		want BillingControllerInterface
	}{
		{
			name: "success",
			args: args{
				service:  nil,
				uservice: nil,
				logger:   nil,
			},
			want: &billingController{
				service:  nil,
				uservice: nil,
				logger:   nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBillingController(tt.args.service, tt.args.uservice, tt.args.logger); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("User controller = %v, want %v", got, tt.want)
			}
		})
	}
}
