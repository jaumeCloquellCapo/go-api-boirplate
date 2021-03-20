package controller

import (
	errorNotFound "ApiRest/app/error"
	"ApiRest/app/model"
	"ApiRest/app/service"
	"ApiRest/internal/logger"
	"ApiRest/mock"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)


func TestMicroservice_Find(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userUC := mock.NewMockUserServiceCase(ctrl)

	apiLogger := logger.NewAPILogger()
	apiLogger.InitLogger()

	userController := NewUserController(userUC, apiLogger)

	reqValue := &model.CreateUser{
		Name:       "FirstName",
		LastName:   "LastName",
		Email:      "email@gmail.com",
		Country:    "es",
		Phone:      "es",
		PostalCode: "es",
	}

	t.Run("Correct", func(t *testing.T) {
		userRes := &model.User{
			ID:         1,
			Name:       reqValue.Name,
			LastName:   reqValue.LastName,
			Email:      reqValue.Email,
			Country:    reqValue.Country,
			Phone:      reqValue.Phone,
			PostalCode: reqValue.PostalCode,
		}

		userUC.EXPECT().FindById(1).Return(userRes, nil)

		router := gin.Default()
		router.GET("/api/users/:id", userController.Find)
		ts := httptest.NewServer(router)
		defer ts.Close()
		w := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/api/users/1", nil)

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("Incorrect", func(t *testing.T) {
		var err errorNotFound.NotFound
		userUC.EXPECT().FindById(2).Return(nil, &err)

		router := gin.Default()
		router.GET("/api/users/:id", userController.Find)
		ts := httptest.NewServer(router)
		defer ts.Close()
		w := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/api/users/2", nil)

		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusNotFound, w.Code)
	})

	t.Run("Incorrect_2", func(t *testing.T) {

		router := gin.Default()
		router.GET("/api/users/:id", userController.Find)
		ts := httptest.NewServer(router)
		defer ts.Close()
		w := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/api/users/pa", nil)

		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestNewUserController(t *testing.T) {
	type args struct {
		service service.UserServiceInterface
		logger  logger.Logger
	}
	tests := []struct {
		name string
		args args
		want UserControllerInterface
	}{
		{
			name: "success",
			args: args{
				service: nil,
				logger:  nil,
			},
			want: &userController{
				service: nil,
				logger:  nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserController(tt.args.service, tt.args.logger); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("User controller = %v, want %v", got, tt.want)
			}
		})
	}
}
