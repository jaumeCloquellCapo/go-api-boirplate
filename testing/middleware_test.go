package testing

import (
	"ApiRest/app/model"
	"ApiRest/app/service"
	"ApiRest/dic"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMiddlewareCorrectToken(t *testing.T) {
	rr := httptest.NewRecorder()

	ar := container.Get(dic.AuthService).(service.AuthServiceInterface)

	newUser := model.CreateUser{
		Name:       gofakeit.Name(),
		Email:      gofakeit.Email(),
		Password:   gofakeit.HackerPhrase(),
		PostalCode: gofakeit.StreetNumber(),
		Country:    gofakeit.Country(),
		Phone:      gofakeit.Phone(),
		LastName:   gofakeit.LastName(),
	}
	//var user model.User
	user, td, err := ar.SignUp(newUser)
	if err != nil {
		t.Fatal(err.Error())
	}

	var id = user.ID

	req, err := http.NewRequest("GET", fmt.Sprint("/auth/users/", id), nil)

	req.Header.Set("Authorization", "Bearer "+td.AccessToken)

	if err != nil {
		t.Fatal(err)
	}

	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

}

func TestMiddlewareIncorrectToken(t *testing.T) {
	rr := httptest.NewRecorder()

	ar := container.Get(dic.AuthService).(service.AuthServiceInterface)

	newUser := model.CreateUser{
		Name:       gofakeit.Name(),
		Email:      gofakeit.Email(),
		Password:   gofakeit.HackerPhrase(),
		PostalCode: gofakeit.StreetNumber(),
		Country:    gofakeit.Country(),
		Phone:      gofakeit.Phone(),
		LastName:   gofakeit.LastName(),
	}
	//var user model.User
	user, _, err := ar.SignUp(newUser)
	if err != nil {
		t.Fatal(err.Error())
	}

	var id = user.ID

	req, err := http.NewRequest("GET", fmt.Sprint("/auth/users/", id), nil)

	req.Header.Set("Authorization", "Bearer ")

	if err != nil {
		t.Fatal(err)
	}

	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusUnauthorized, rr.Code)

}
