package tests

import (
	"ApiRest/internal/dic"
	"ApiRest/internal/route"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sarulabs/dingo/generation/di"
	"log"
	"os"
	"testing"
)

var router *gin.Engine
var container di.Container

func TestMain(m *testing.M) {
	//gin.SetMode(gin.TestMode)

	err := godotenv.Load("../dev.env")
	gofakeit.Seed(0)

	if err != nil {
		log.Fatal("Error loading dev.env file")
	}

	container = dic.InitContainer()
	router = route.Setup(container)
	exitVal := m.Run()
	os.Exit(exitVal)
}
