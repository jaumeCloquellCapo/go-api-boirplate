package main

import (
	"ApiRest/app/controller"
	"ApiRest/app/middleware"
	"ApiRest/app/repository"
	"ApiRest/app/service"
	"ApiRest/config"
	"ApiRest/provider"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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

	cfg, errCfg := config.Load("dev")

	if errCfg != nil {
		log.Fatal(errCfg)
	}

	db, errDb := provider.InitializeDB(cfg)
	if errDb != nil {
		log.Fatal("Error InitializeDB")
		log.Fatal(errDb)
	}
	defer db.Close()

	cache := provider.InitializeCache(cfg)
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
	router.GET("/user/:id", middleware.TokenAuthMiddleware(), userController.GetById)

	router.Run()
}
