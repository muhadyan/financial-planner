package service

import "github.com/muhadyan/financial-planner/repository"

type ExampleService struct {
	ExampleRepository repository.ExampleRepository
}

func (exampleService *ExampleService) GetExampleName() string {
	return exampleService.ExampleRepository.GetExampleName()
}
