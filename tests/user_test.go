package tests

import (
	"ApiRest/app/model"
	"ApiRest/app/service"
	"ApiRest/internal/dic"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUpdateUser(t *testing.T) {
	rr := httptest.NewRecorder()

	ar := container.Get(dic.AuthService).(service.AuthServiceInterface)
	ac := container.Get(dic.UserService).(service.UserServiceInterface)

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

	userToUpdate := model.UpdateUser{
		Name:       gofakeit.Name(),
		Email:      gofakeit.Email(),
		Password:   gofakeit.HackerPhrase(),
		PostalCode: gofakeit.StreetNumber(),
		Country:    gofakeit.Country(),
		Phone:      gofakeit.Phone(),
		LastName:   gofakeit.LastName(),
	}

	body, _ := json.Marshal(userToUpdate)

	req, err := http.NewRequest("PUT", fmt.Sprint("/auth/users/", id), bytes.NewBuffer(body))

	req.Header.Set("Authorization", "Bearer "+td.AccessToken)

	if err != nil {
		t.Fatal(err)
	}

	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	userUpdated, err := ac.FindById(id)

	if err != nil {
		t.Fatal(err.Error())
	}
	if userUpdated.Name != userToUpdate.Name || userUpdated.Email != userToUpdate.Email || userUpdated.Phone != userToUpdate.Phone {
		t.Fatal("Not Equal")
	}
	//assert.EqualValues(t, userToUpdate, userUpdated)
}
