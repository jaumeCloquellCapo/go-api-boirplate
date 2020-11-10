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
	r.POST("/login", ac.Login)
	r.POST("/logout", ac.Logout)
	r.POST("/signUp", ac.SignUp)

	authMiddleware := dic.Container.Get(dic.AuthMiddleware).(middleware.AuthMiddlewareInterface)

	authorized := r.Group("/auth", authMiddleware.Handler())
	uc := dic.Container.Get(dic.UserController).(controller.UserControllerInterface)
	authorized.GET("/users", uc.GetUsers)
	authorized.GET("/users/:id", uc.GetUserById)

	return r
}
