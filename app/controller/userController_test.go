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

func TestUserController_UpdateUserById(t *testing.T) {
	us := &mock.MockUserService{}

	ctl := NewUserController(us)
	router := gin.Default()
	router.POST("/:id", ctl.UpdateUserById)
	ts := httptest.NewServer(router)
	defer ts.Close()

	credentials := model.UpdateUser{
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
	req := httptest.NewRequest("POST", "/1", bytes.NewBuffer(body))

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUserController_UpdateUserById1(t *testing.T) {
	us := &mock.MockUserService{}

	ctl := NewUserController(us)
	router := gin.Default()
	router.POST("/:id", ctl.UpdateUserById)
	ts := httptest.NewServer(router)
	defer ts.Close()

	credentials := model.UpdateUser{
		Name:       gofakeit.Name(),
		LastName:   gofakeit.LastName(),
		Email:      gofakeit.Email(),
	}

	body, _ := json.Marshal(credentials)

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/1", bytes.NewBuffer(body))

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUserController_UpdateUserById2(t *testing.T) {
	us := &mock.MockUserService{}

	ctl := NewUserController(us)
	router := gin.Default()
	router.POST("/:id", ctl.UpdateUserById)
	ts := httptest.NewServer(router)
	defer ts.Close()

	credentials := model.UpdateUser{
		Name:       gofakeit.Name(),
		LastName:   gofakeit.LastName(),
		Email:      "a",
		Country:    gofakeit.Country(),
		Phone:    	gofakeit.Phone(),
		PostalCode: "07440",
		Password:   gofakeit.HackerPhrase(),
	}

	body, _ := json.Marshal(credentials)

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/1", bytes.NewBuffer(body))

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}


func TestNewUserController(t *testing.T) {
	type args struct {
		service service.UserServiceInterface
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
			},
			want: &userController{
				service: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserController(tt.args.service); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("User controller = %v, want %v", got, tt.want)
			}
		})
	}
}
