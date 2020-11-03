package main

import (
	"ApiRest/app/controller"
	"ApiRest/app/repository"
	"ApiRest/app/service"
	"ApiRest/config"
	"ApiRest/provider"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	cfg, errCfg := config.Load("dev")

	if errCfg != nil {
		panic(errCfg)
	}

	db, errDb := provider.InitializeDB(cfg)

	defer db.Close()

	if errDb != nil {
		panic(errDb)
	}

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/user/:id", userController.GetById)

	router.Run()
}
