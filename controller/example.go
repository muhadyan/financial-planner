package controller

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/muhadyan/financial-planner/service"
)

type ExampleController struct {
	ExampleService service.ExampleService
}

func (exampleController *ExampleController) GetExampleName(c echo.Context) error {
	return c.JSON(http.StatusOK, exampleController.ExampleService.GetExampleName())
}
