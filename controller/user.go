package controller

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/muhadyan/financial-planner/model"
	"github.com/muhadyan/financial-planner/service"
)

type UserController struct {
	UserService service.UserService
}

func (userController *UserController) SignUp(c echo.Context) error {
	params := new(model.SignUpRequest)

	err := c.Bind(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.BasicResp{Message: err.Error()})
	}

	err = validateSignUp(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.BasicResp{Message: err.Error()})
	}

	user, err := userController.UserService.SignUp(params)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.BasicResp{Message: err.Error()})
	}

	resp := model.BasicResp{
		Message: "Success",
		Data:    user,
	}

	return c.JSON(http.StatusOK, resp)
}
