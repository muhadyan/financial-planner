package controller

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/muhadyan/financial-planner/model"
	"github.com/muhadyan/financial-planner/service"
)

type GoldController struct {
	GoldService service.GoldService
}

func (goldController *GoldController) Create(c echo.Context) error {
	params := new(model.CreateUserGoldRequest)
	userID := c.Get("id").(int)

	err := c.Bind(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.BasicResp{Message: err.Error()})
	}
	params.UserID = userID

	err = validateCreateUserGold(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.BasicResp{Message: err.Error()})
	}

	userGold, err := goldController.GoldService.Create(params)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.BasicResp{Message: err.Error()})
	}

	resp := model.BasicResp{
		Message: "Success",
		Data:    userGold,
	}

	return c.JSON(http.StatusOK, resp)
}
