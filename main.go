package main

import (
	"ApiRest/dic"
	"ApiRest/route"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/twinj/uuid"
	"log"
	"os"
)

// RequestIDMiddleware : create unique uuid to attach to every request
func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := uuid.NewV4()
		c.Writer.Header().Set("X-Request-Id", uuid.String())
		c.Next()
	}
}

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
