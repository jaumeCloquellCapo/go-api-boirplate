package controller

import (
	"ApiRest/app/model"
	"ApiRest/app/service"
	"ApiRest/mock"
	"bytes"
	"encoding/json"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestAuthController_Login(t *testing.T) {
	ms := &mock.MockAuthService{}
	us := &mock.MockUserService{}

	ctl := NewAuthController(ms, us)
	router := gin.Default()
	router.POST("/login", ctl.Login)
	ts := httptest.NewServer(router)
	defer ts.Close()

	credentials := model.Credentials{
		Email:    gofakeit.Email(),
		Password:  gofakeit.HackerPhrase(),
	}

	body, _ := json.Marshal(credentials)

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/login", bytes.NewBuffer(body))

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestAuthController_Login2(t *testing.T) {
	ms := &mock.MockAuthService{}
	us := &mock.MockUserService{}

	ctl := NewAuthController(ms, us)
	router := gin.Default()
	router.POST("/login", ctl.Login)
	ts := httptest.NewServer(router)
	defer ts.Close()

	credentials := model.Credentials{
		Email:     "--",
		Password:  gofakeit.HackerPhrase(),
	}

	body, _ := json.Marshal(credentials)

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/login", bytes.NewBuffer(body))

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestAuthController_Login3(t *testing.T) {
	ms := &mock.MockAuthService{}
	us := &mock.MockUserService{}

	ctl := NewAuthController(ms, us)
	router := gin.Default()
	router.POST("/login", ctl.Login)
	ts := httptest.NewServer(router)
	defer ts.Close()


	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/login", nil)

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestAuthController_SignUp(t *testing.T) {
	ms := &mock.MockAuthService{}
	us := &mock.MockUserService{}

	ctl := NewAuthController(ms, us)
	router := gin.Default()
	router.POST("/signup", ctl.SignUp)
	ts := httptest.NewServer(router)
	defer ts.Close()


	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/signup", nil)

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestAuthController_SignUp1(t *testing.T) {
	ms := &mock.MockAuthService{}
	us := &mock.MockUserService{}

	ctl := NewAuthController(ms, us)
	router := gin.Default()
	router.POST("/signup", ctl.SignUp)
	ts := httptest.NewServer(router)
	defer ts.Close()


	credentials := model.CreateUser{
		Name:       gofakeit.Name(),
		LastName:   gofakeit.LastName(),
		Email:      gofakeit.Email(),
		Country:    gofakeit.Country(),
		Phone:    	gofakeit.Phone(),
		PostalCode: "07440",
		Password:   gofakeit.HackerPhrase(),
	}

	body, _ := json.Marshal(credentials)

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/signup", bytes.NewBuffer(body))

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestAuthController_SignUp2(t *testing.T) {
	ms := &mock.MockAuthService{}
	us := &mock.MockUserService{}

	ctl := NewAuthController(ms, us)
	router := gin.Default()
	router.POST("/signup", ctl.SignUp)
	ts := httptest.NewServer(router)
	defer ts.Close()


	credentials := model.CreateUser{
		Name:       gofakeit.Name(),
		LastName:   gofakeit.LastName(),
		Email:      "bad",
		Country:    gofakeit.Country(),
		Phone:    	gofakeit.Phone(),
		PostalCode: "07440",
		Password:   gofakeit.HackerPhrase(),
	}

	body, _ := json.Marshal(credentials)

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/signup", bytes.NewBuffer(body))

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestNewAuthController(t *testing.T) {
	type args struct {
		authService service.AuthServiceInterface
		userService service.UserServiceInterface
	}
	tests := []struct {
		name string
		args args
		want AuthControllerInterface
	}{
		{
			name: "success",
			args: args{
				authService: nil,
				userService: nil,
			},
			want: &authController{
				authService: nil,
				userService: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAuthController(tt.args.authService, tt.args.userService); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Auth controller = %v, want %v", got, tt.want)
			}
		})
	}
}
