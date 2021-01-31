package main

import (
	"ApiRest/cmd"
	"github.com/joho/godotenv"
	"log"
)

func main() {

	err := godotenv.Load("dev.env")
	if err != nil {
		log.Fatal("Error loading dev.env file")
	}

	cmd.Execute()

}
