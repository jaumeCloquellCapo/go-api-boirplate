package main

import (
	"ApiRest/internal/dic"
	"ApiRest/internal/route"
	"flag"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var config string
func main() {

	flag.StringVar(&config, "env", "dev.env", "help message for flagname")
	flag.Parse()


	if err := godotenv.Load(config); err != nil {
		log.Fatalf("Error loading %v", "dev.env")
	}
	container := dic.InitContainer()
	router := route.Setup(container)
	router.Run(":" + os.Getenv("APP_PORT"))

}
