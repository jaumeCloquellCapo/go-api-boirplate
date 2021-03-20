package dic

import (
	"ApiRest/app/repository"
	"ApiRest/app/service"
	"ApiRest/internal/middleware"
	"ApiRest/internal/storage"
	"github.com/sarulabs/dingo/generation/di"
	"log"
)

const DbService = "db"
const CacheService = "cache"

const AuthMiddleware = "middleware.auth"
const CorsMiddleware = "middleware.cors"

const UserRepository = "repository.user"
const UserService = "service.user"
const UserController = "controller.user"

const AuthRepository = "repository.auth"
const AuthService = "service.auth"
const AuthController = "controller.auth"

// dependency injection container
func InitContainer() di.Container {
	builder, err := di.NewBuilder()
	if err != nil {
		log.Fatal(err.Error())
	}
	RegisterServices(builder)
	return builder.Build()
}

func RegisterServices(builder *di.Builder) {
	builder.Add(di.Def{
		Name: DbService,
		Build: func(ctn di.Container) (interface{}, error) {
			return storage.InitializeDB(), nil
		},
		Close: func(obj interface{}) error {
			obj.(*storage.DbStore).Close()
			return nil
		},
	})
	builder.Add(di.Def{
		Name: CacheService,
		Build: func(ctn di.Container) (interface{}, error) {
			return storage.InitializeCache(), nil
		},
		Close: func(obj interface{}) error {
			obj.(*storage.DbCache).Close()
			return nil
		},
	})

	builder.Add(di.Def{
		Name: CorsMiddleware,
		Build: func(ctn di.Container) (interface{}, error) {
			return middleware.NewCorsMiddleware(), nil
		},
	})

	builder.Add(di.Def{
		Name: UserRepository,
		Build: func(ctn di.Container) (interface{}, error) {
			return repository.NewUserRepository(ctn.Get(DbService).(*storage.DbStore)), nil
		},
	})

	builder.Add(di.Def{
		Name: UserService,
		Build: func(ctn di.Container) (interface{}, error) {
			return service.NewUserService(ctn.Get(UserRepository).(repository.UserRepositoryInterface)), nil
		},
	})

	/*builder.Add(di.Def{
		Name: UserController,
		Build: func(ctn di.Container) (interface{}, error) {
			return controller.NewUserController(ctn.Get(UserService).(service.UserServiceInterface)), nil
		},
	}) */

}
