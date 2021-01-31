package route

import (
	"ApiRest/app/controller"
	"ApiRest/internal/dic"
	"ApiRest/internal/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sarulabs/dingo/generation/di"
	"net/http"
	"os"
	"time"
)

//Setup ...
func Setup(container di.Container) *gin.Engine {

	// Controllers
	ac := container.Get(dic.AuthController).(controller.AuthControllerInterface)
	uc := container.Get(dic.UserController).(controller.UserControllerInterface)

	// Middleware
	authMiddleware := container.Get(dic.AuthMiddleware).(middleware.AuthMiddlewareInterface)
	corsMiddleware := container.Get(dic.CorsMiddleware).(middleware.CorsMiddlewareInterface)

	gin.SetMode(os.Getenv("GIN_MODE"))
	r := gin.New()
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
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
		users := authorized.Group("/users")
		{
			users.GET("/:id", uc.FindUserById)
			users.DELETE("/:id", uc.RemoveUserById)
			users.PUT("/:id", uc.UpdateUserById)

		}
	}

	return r
}
