package route

import (
	"github.com/muhadyan/financial-planner/controller"
	"github.com/muhadyan/financial-planner/repository"
	"github.com/muhadyan/financial-planner/service"
)

type AppModels struct {
	Example controller.ExampleController
	User    controller.UserController
}

func App() AppModels {
	inMemoryExampleRepository := &repository.InMemoryExampleRepository{}
	userRepository := &repository.UserRepositoryCtx{}
	timeRepository := &repository.TimeRepositoryCtx{}

	exampleService := service.ExampleService{
		ExampleRepository: inMemoryExampleRepository,
	}
	userService := service.UserService{
		UserRepository: userRepository,
		TimeRepository: timeRepository,
	}

	exampleController := controller.ExampleController{
		ExampleService: exampleService,
	}
	userController := controller.UserController{
		UserService: userService,
	}

	return AppModels{
		Example: exampleController,
		User:    userController,
	}
}
