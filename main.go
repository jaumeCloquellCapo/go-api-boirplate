package main

import (
	"ApiRest/internal/dic"
	"ApiRest/internal/logger"
	"ApiRest/internal/route"
	"flag"
	"github.com/joho/godotenv"
	"os"
)

var config string

func main() {

	flag.StringVar(&config, "env", "dev.env", "Environment name")
	flag.Parse()

	logger := logger.NewAPILogger()
	logger.InitLogger()

	if err := godotenv.Load(config); err != nil {
		logger.Fatalf(err.Error())
	}
	container := dic.InitContainer()
	router := route.Setup(container, logger)
	router.Run(":" + os.Getenv("APP_PORT"))

}
