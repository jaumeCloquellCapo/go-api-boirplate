package testing

import (
	"ApiRest/app/model"
	"ApiRest/app/service"
	"ApiRest/dic"
	"ApiRest/route"
	"bytes"
	"encoding/json"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sarulabs/dingo/generation/di"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var router *gin.Engine
var container di.Container

func TestMain(m *testing.M) {

	err := godotenv.Load("../dev.env")

	if err != nil {
		log.Fatal("Error loading dev.env file")
	}

	gofakeit.Seed(0)

	container = dic.InitContainer()
	router = route.Setup(container)
	gin.SetMode(gin.TestMode)
	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestCreateUser(t *testing.T) {
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()

	expectedUser := model.CreateUser{
		Name:     gofakeit.Name(),
		Email:    gofakeit.Email(),
		Password: gofakeit.HackerPhrase(),
	}

	body, _ := json.Marshal(expectedUser)

	req, err := http.NewRequest("POST", "/signup", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusCreated, rr.Code)

}

func TestCreateBadUser(t *testing.T) {
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

func TestLoginUserBadBody(t *testing.T) {
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

func TestLoginFakeUser(t *testing.T) {
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
