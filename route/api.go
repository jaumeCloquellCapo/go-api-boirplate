package route

import (
	"ApiRest/app/controller"
	"ApiRest/app/middleware"
	"ApiRest/dic"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Setup() *gin.Engine {

	// Controllers
	ac := dic.Container.Get(dic.AuthController).(controller.AuthControllerInterface)
	uc := dic.Container.Get(dic.UserController).(controller.UserControllerInterface)

	// Middleware
	authMiddleware := dic.Container.Get(dic.AuthMiddleware).(middleware.AuthMiddlewareInterface)
	corsMiddleware := dic.Container.Get(dic.CorsMiddleware).(middleware.CorsMiddlewareInterface)

	r := gin.New()
	//r.Use(limit.Limit(200)) // limit the number of current requests
	r.Use(gin.Recovery())
	r.Use(corsMiddleware.Handler())
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	// public endpoints
	r.POST("/login", ac.Login)
	r.POST("/signup", ac.SignUp)
	// private endpoints
	authorized := r.Group("/auth", authMiddleware.Handler())
	{
		//authorized.POST("/signup", ac.SignUp)
		authorized.POST("/logout", ac.Logout)
		authorized.GET("/users", uc.GetUsers)
		authorized.GET("/users/:id", uc.GetUserById)
	}

	return r
}
