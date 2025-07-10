package router

import (
	"go_auth/config"
	database "go_auth/databases"
	"go_auth/databases/repositories"
	controller "go_auth/handlers/controllers"
	"go_auth/usecases"
)

type controllerCollections struct {
	userController controller.UserController
	// ... other controllers
}

func InitController(db database.DBs, cfg config.MainConfig) controllerCollections {

	userRepo := repositories.NewUserRepository(db.Movies)
	userUsecase := usecases.NewUserUsecase(userRepo)
	userController := controller.NewUserController(userUsecase)

	return controllerCollections{
		userController: userController,
		// ... initialize other controllers
	}
}
