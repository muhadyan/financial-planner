package service

import (
	"testing"

	"github.com/muhadyan/financial-planner/repository/mocks"
)

func TestExampleService_GetExampleName(t *testing.T) {
	mockExampleRepo := new(mocks.ExampleRepository)

	tests := []struct {
		name       string
		want       string
		beforeFunc func() *ExampleService
	}{
		{
			name: "Success",
			want: "John Doe",
			beforeFunc: func() *ExampleService {
				s := &ExampleService{
					ExampleRepository: mockExampleRepo,
				}

				mockExampleRepo.On("GetExampleName").Return("John Doe").Once()

				return s
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.beforeFunc()
			if got := s.GetExampleName(); got != tt.want {
				t.Errorf("ExampleService.GetExampleName() = %v, want %v", got, tt.want)
			}
		})
	}
}
