package service

import (
	"ApiRest/app/model"
	"ApiRest/app/repository"
	"ApiRest/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func TestNewUserService(t *testing.T) {
	type args struct {
		userRepository repository.UserRepositoryInterface
	}
	tests := []struct {
		name string
		args args
		want UserServiceInterface
	}{
		{
			name: "success",
			args: args{
				userRepository: nil,
			},
			want: &userService{
				userRepo: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserService(tt.args.userRepository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserInit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_Store(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	userR := mock.NewMockUserPGRepository(ctrl)
	userService := NewUserService(userR)

	reqValue := model.CreateUser{
		Name:       "a",
		LastName:   "a",
		Email:      "a@a.com",
		Country:    "a",
		Phone:      "a",
		PostalCode: "a",
	}

	t.Run("Store", func(t *testing.T) {
		t.Parallel()

		user := model.CreateUser{
			Name:       reqValue.Name,
			LastName:   reqValue.LastName,
			Email:      reqValue.Email,
			Country:    reqValue.Country,
			Phone:      reqValue.Phone,
			PostalCode: reqValue.PostalCode,
		}

		userID := int(1)
		userRes := &model.User{
			ID:         userID,
			Name:       user.Name,
			LastName:   user.LastName,
			Email:      user.Email,
			Country:    user.Country,
			Phone:      user.Phone,
			PostalCode: user.PostalCode,
		}
		var err error

		userR.EXPECT().Create(user).Return(userRes, err)

		response, err := userService.Store(reqValue)

		require.NoError(t, err)
		require.NotNil(t, response)
	})
}
