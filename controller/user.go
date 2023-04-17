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

func (userController *UserController) Verify(c echo.Context) error {
	params := new(model.VerifyRequest)

	err := c.Bind(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.BasicResp{Message: err.Error()})
	}

	err = validateVerify(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.BasicResp{Message: err.Error()})
	}

	err = userController.UserService.Verify(params)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.BasicResp{Message: err.Error()})
	}

	resp := model.BasicResp{
		Message: "Thank you! Your account is verified.",
	}

	return c.JSON(http.StatusOK, resp)
}

func (userController *UserController) LogIn(c echo.Context) error {
	params := new(model.LogInRequest)

	err := c.Bind(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.BasicResp{Message: err.Error()})
	}

	err = validateLogIn(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.BasicResp{Message: err.Error()})
	}

	login, err := userController.UserService.LogIn(params)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.BasicResp{Message: err.Error()})
	}

	resp := model.BasicResp{
		Message: "Success",
		Data:    login,
	}

	return c.JSON(http.StatusOK, resp)
}
