package route

import (
	"ApiRest/app/controller"
	"ApiRest/app/service"
	"ApiRest/internal/dic"
	"ApiRest/internal/logger"
	"ApiRest/internal/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sarulabs/dingo/generation/di"
	"net/http"
	"os"
	"time"
)

// Setup returns initialized routes.
func Setup(container di.Container, logger logger.Logger) *gin.Engine {
	// ac := container.Get(dic.AuthController).(controller.AuthControllerInterface)

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

	r.Use(gin.Recovery())

	// Middleware initialization
	corsMiddleware := container.Get(dic.CorsMiddleware).(middleware.CorsMiddlewareInterface)
	r.Use(corsMiddleware.Handler())

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	// v1 Routes

	//uc := container.Get(dic.UserController).(controller.UserControllerInterface)
	uc := controller.NewUserController(container.Get(dic.UserService).(service.UserServiceInterface), logger)
	v1 := r.Group("/api")
	{
		users := v1.Group("/users")
		{
			users.GET("/:id", uc.Find)
			users.DELETE("/:id", uc.Destroy)
			users.PUT("/:id", uc.Update)

		}
	}

	return r
}
