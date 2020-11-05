package main

import (
	"ApiRest/app/controller"
	"ApiRest/app/middleware"
	"ApiRest/app/repository"
	"ApiRest/app/service"
	"ApiRest/provider"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/twinj/uuid"
	"log"
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

	db := provider.InitializeDB()
	defer db.Close()

	cache := provider.InitializeCache()
	defer cache.Close()

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	authRepo := repository.NewAuthRepository(cache)
	authService := service.NewAuthService(authRepo)
	authController := controller.NewAuthController(authService, userService)

	router := gin.Default()
	router.Use(cors.Default())
	//router.Use(RequestIDMiddleware())

	router.POST("/login", authController.Login)
	router.GET("/user/:id", middleware.TokenAuthMiddleware(), userController.GetUserById)

	router.Run()
}
