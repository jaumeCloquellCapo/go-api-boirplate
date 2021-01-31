package testing

import (
	"ApiRest/app/model"
	"ApiRest/app/service"
	"ApiRest/internal/dic"
	"bytes"
	"encoding/json"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSignupCorrect(t *testing.T) {
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()

	expectedUser := model.CreateUser{
		Name:       gofakeit.Name(),
		Email:      gofakeit.Email(),
		Password:   gofakeit.HackerPhrase(),
		PostalCode: gofakeit.StreetNumber(),
		Country:    gofakeit.Country(),
		Phone:      gofakeit.Phone(),
		LastName:   gofakeit.LastName(),
	}

	body, _ := json.Marshal(expectedUser)

	req, err := http.NewRequest("POST", "/signup", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusCreated, rr.Code)

}

func TestSignupBadRequest(t *testing.T) {
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()

	// Test bad params
	expectedUser := model.CreateUser{
		Name:     gofakeit.Name(),
		Email:    "aa",
		Password: gofakeit.HackerPhrase(),
	}

	body, _ := json.Marshal(expectedUser)

	req, err := http.NewRequest("POST", "/signup", bytes.NewBuffer(body))

	if err != nil {
		t.Fatal(err)
	}

	router.ServeHTTP(rr, req)

	assert.NotEqual(t, http.StatusCreated, rr.Code)

}
func TestLoginUser(t *testing.T) {
	rr := httptest.NewRecorder()

	ar := container.Get(dic.AuthService).(service.AuthServiceInterface)

	expectedUser := model.CreateUser{
		Name:     gofakeit.Name(),
		Email:    gofakeit.Email(),
		Password: gofakeit.HackerPhrase(),
	}

	ar.SignUp(expectedUser)

	credentials := model.Credentials{
		Email:    expectedUser.Email,
		Password: expectedUser.Password,
	}

	body, _ := json.Marshal(credentials)
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(body))

	if err != nil {
		t.Fatal(err)
	}

	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestLoginUserBadRequest(t *testing.T) {
	rr := httptest.NewRecorder()
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("POST", "/login", nil)

	if err != nil {
		t.Fatal(err)
	}

	router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
}

func TestLoginNotExist(t *testing.T) {
	rr := httptest.NewRecorder()

	// Test bad params
	expectedUser := model.Credentials{
		Email:    gofakeit.Email(),
		Password: gofakeit.HackerPhrase(),
	}

	body, _ := json.Marshal(expectedUser)

	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(body))

	if err != nil {
		t.Fatal(err)
	}

	router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusNotFound, rr.Code)
}
