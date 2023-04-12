package repository

type ExampleRepository interface {
	GetExampleName() string
}

type InMemoryExampleRepository struct{}

func (inMemoryExampleRepository *InMemoryExampleRepository) GetExampleName() string {
	return "John Doe"
}
