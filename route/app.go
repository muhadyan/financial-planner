package route

import (
	"github.com/muhadyan/financial-planner/controller"
	"github.com/muhadyan/financial-planner/repository"
	"github.com/muhadyan/financial-planner/service"
)

type AppModels struct {
	Example controller.ExampleController
}

func App() AppModels {
	inMemoryExampleRepository := &repository.InMemoryExampleRepository{}

	exampleService := service.ExampleService{
		ExampleRepository: inMemoryExampleRepository,
	}

	exampleController := controller.ExampleController{
		ExampleService: exampleService,
	}

	return AppModels{
		Example: exampleController,
	}
}
