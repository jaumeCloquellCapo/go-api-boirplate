package route

import (
	"ApiRest/app/controller"
	"ApiRest/app/middleware"
	"ApiRest/dic"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Setup() *gin.Engine {

	r := gin.New()
	r.Use(gin.Recovery())

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	ac := dic.Container.Get(dic.AuthController).(controller.AuthControllerInterface)
	uc := dic.Container.Get(dic.UserController).(controller.UserControllerInterface)
	authMiddleware := dic.Container.Get(dic.AuthMiddleware).(middleware.AuthMiddlewareInterface)

	r.POST("/login", ac.Login)
	r.POST("/logout", ac.Logout)
	authorized := r.Group("/auth", authMiddleware.Handler())
	{
		authorized.POST("/register", ac.SignUp)
		authorized.GET("/users", uc.GetUsers)
		authorized.GET("/users/:id", uc.GetUserById)
	}

	return r
}
