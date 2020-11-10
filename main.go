package main

import (
	"ApiRest/dic"
	"ApiRest/route"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load("dev.env")
	if err != nil {
		log.Fatal("Error loading dev.env file")
	}
	//cmd.Execute()
	dic.InitContainer()

	router := route.Setup()
	router.Run(":" + os.Getenv("APP_PORT"))
}
