package controller

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/muhadyan/financial-planner/model"
	"github.com/muhadyan/financial-planner/service"
	"github.com/muhadyan/financial-planner/utils"
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

func (goldController *GoldController) GetUnrealized(c echo.Context) error {
	params := new(model.GetUnrealizedRequest)
	userID := c.Get("id").(int)
	params.UserID = userID

	err := c.Bind(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.BasicResp{Message: err.Error()})
	}

	err = validateGetUnrealized(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.BasicResp{Message: err.Error()})
	}

	unrealized, err := goldController.GoldService.GetUnrealized(params)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.BasicResp{Message: err.Error()})
	}

	totalPage := utils.TotalPage(len(unrealized), params.Limit)

	resp := model.Pagination{
		BasicResp: model.BasicResp{
			Message: "Success",
			Data:    unrealized,
		},
		Meta: model.Meta{
			Page:         params.Page,
			Limit:        params.Limit,
			TotalRecords: len(unrealized),
			TotalPages:   int(totalPage),
		},
	}

	return c.JSON(http.StatusOK, resp)
}
