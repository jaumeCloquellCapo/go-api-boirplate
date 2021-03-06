package service

import (
	"ApiRest/app/model"
	"ApiRest/app/model/billing"
	"ApiRest/app/repository"
	"ApiRest/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
	"time"
)

func TestNewBillingService(t *testing.T) {
	type args struct {
		paymentRepo repository.BillingRepositoryInterface
	}
	tests := []struct {
		name string
		args args
		want BillingServiceInterface
	}{
		{
			name: "success",
			args: args{
				paymentRepo: nil,
			},
			want: &billingService{
				paymentRepo: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBillingService(tt.args.paymentRepo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserInit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBillingService_AddBilling(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	billingR := mock.NewMockBillingPGRepository(ctrl)
	billingService := NewBillingService(billingR)

	t.Run("InvalidPayment", func(t *testing.T) {
		t.Parallel()
		user := model.User{
			ID:         1,
			Name:       "1",
			Cif:        "1",
			Country:    "1",
			PostalCode: "1",
		}



		p := billing.Payment{
			Identify:       billing.AccountPaypal,
			CustomerParams: billing.CustomerParams{
				Email: "test3@test.com",
				Desc:  "a 3rd test customer",
				Card: &billing.CardParams{
					Name:     user.Name,
					Number:   user.Cif,
					ExpYear:  time.Now().Year() + 1,
					ExpMonth: 1,
				},
			},
			PaymentMethod: &mock.FakeAdapter{},
		}



		var err error

		billingR.EXPECT().CreateBillingService(p.Identify, "fake", user.ID).Return(err) // identify model.Identify, PaymentUserKey string, userID int)

		err = billingService.AddBilling(user, p)

		require.NoError(t, err)
	})
}


func TestBillingService_GetPaymentAdapter(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	billingR := mock.NewMockBillingPGRepository(ctrl)
	billingService := NewBillingService(billingR)

	t.Run("InvalidPayment", func(t *testing.T) {
		t.Parallel()
		user := model.User{
			ID:         1,
			Name:       "1",
			Cif:        "1",
			Country:    "1",
			PostalCode: "1",
		}



		p := billing.CreateCustomer{
			Identify:       billing.AccountPaypal,
			CustomerParams: billing.CustomerParams{
				Email: "test3@test.com",
				Desc:  "a 3rd test customer",
				Card: &billing.CardParams{
					Name:     user.Name,
					Number:   user.Cif,
					ExpYear:  time.Now().Year() + 1,
					ExpMonth: 1,
				},
			},
		}



		var err error

		presult, err := billingService.GetPaymentAdapter(p)

		require.NotNil(t, presult)
		require.NoError(t, err)
	})
}

func TestBillingService_GetPaymentAdapter2(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	billingR := mock.NewMockBillingPGRepository(ctrl)
	billingService := NewBillingService(billingR)

	t.Run("InvalidPayment", func(t *testing.T) {
		t.Parallel()
		user := model.User{
			ID:         1,
			Name:       "1",
			Cif:        "1",
			Country:    "1",
			PostalCode: "1",
		}



		p := billing.CreateCustomer{
			Identify:       billing.Identify("bad"),
			CustomerParams: billing.CustomerParams{
				Email: "test3@test.com",
				Desc:  "a 3rd test customer",
				Card: &billing.CardParams{
					Name:     user.Name,
					Number:   user.Cif,
					ExpYear:  time.Now().Year() + 1,
					ExpMonth: 1,
				},
			},
		}

		var err error

		presult, err := billingService.GetPaymentAdapter(p)

		require.Nil(t, presult)
		require.Error(t, err)
	})
}


