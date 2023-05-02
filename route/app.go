package route

import (
	"github.com/muhadyan/financial-planner/controller"
	"github.com/muhadyan/financial-planner/repository"
	"github.com/muhadyan/financial-planner/service"
)

type AppModels struct {
	Example controller.ExampleController
	User    controller.UserController
	Gold    controller.GoldController
}

func App() AppModels {
	// repository
	inMemoryExampleRepository := &repository.InMemoryExampleRepository{}
	userRepository := &repository.UserRepositoryCtx{}
	timeRepository := &repository.TimeRepositoryCtx{}
	roleRepository := &repository.RoleRepositoryCtx{}
	userRoleRepository := &repository.UserRoleRepositoryCtx{}
	userGoldRepository := &repository.UserGoldRepositoryCtx{}
	currentGoldRepository := &repository.CurrentGoldRepositoryCtx{}

	// service
	exampleService := service.ExampleService{
		ExampleRepository: inMemoryExampleRepository,
	}
	userService := service.UserService{
		UserRepository:     userRepository,
		TimeRepository:     timeRepository,
		RoleRepository:     roleRepository,
		UserRoleRepository: userRoleRepository,
	}
	goldService := service.GoldService{
		UserGoldRepository:    userGoldRepository,
		UserRepository:        userRepository,
		CurrentGoldRepository: currentGoldRepository,
	}

	// controller
	exampleController := controller.ExampleController{
		ExampleService: exampleService,
	}
	userController := controller.UserController{
		UserService: userService,
	}
	goldController := controller.GoldController{
		GoldService: goldService,
	}

	return AppModels{
		Example: exampleController,
		User:    userController,
		Gold:    goldController,
	}
}
