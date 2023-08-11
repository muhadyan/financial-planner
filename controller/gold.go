package controller

import (
	"net/http"
	"strconv"

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

func (goldController *GoldController) FindAll(c echo.Context) error {
	params := new(model.FindAllUserGoldRequest)
	userID := c.Get("id").(int)
	params.UserID = userID

	err := c.Bind(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.BasicResp{Message: err.Error()})
	}

	err = validatePaginationReq(&params.PagiantionReq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.BasicResp{Message: err.Error()})
	}

	userGolds, err := goldController.GoldService.FindAll(params)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.BasicResp{Message: err.Error()})
	}

	totalPage := utils.TotalPage(len(userGolds), params.Limit)

	resp := model.Pagination{
		BasicResp: model.BasicResp{
			Message: "Success",
			Data:    userGolds,
		},
		Meta: model.Meta{
			Page:         params.Page,
			Limit:        params.Limit,
			TotalRecords: len(userGolds),
			TotalPages:   totalPage,
		},
	}

	return c.JSON(http.StatusOK, resp)
}

func (goldController *GoldController) Find(c echo.Context) error {
	userID := c.Get("id").(int)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.BasicResp{Message: err.Error()})
	}

	userGold, err := goldController.GoldService.Find(id, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.BasicResp{Message: err.Error()})
	}

	resp := model.BasicResp{
		Message: "Success",
		Data:    userGold,
	}

	return c.JSON(http.StatusOK, resp)
}

func (goldController *GoldController) Update(c echo.Context) error {
	params := new(model.UpdateUserGoldRequest)
	userID := c.Get("id").(int)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.BasicResp{Message: err.Error()})
	}

	err = c.Bind(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.BasicResp{Message: err.Error()})
	}
	params.ID = id
	params.UserID = userID

	err = validateUpdateUserGold(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.BasicResp{Message: err.Error()})
	}

	err = goldController.GoldService.Update(params)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.BasicResp{Message: err.Error()})
	}

	resp := model.BasicResp{
		Message: "Success",
	}

	return c.JSON(http.StatusOK, resp)
}

func (goldController *GoldController) Delete(c echo.Context) error {
	userID := c.Get("id").(int)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.BasicResp{Message: err.Error()})
	}

	err = goldController.GoldService.Delete(id, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.BasicResp{Message: err.Error()})
	}

	resp := model.BasicResp{
		Message: "Success",
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

	err = validatePaginationReq(&params.PagiantionReq)
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
			TotalPages:   totalPage,
		},
	}

	return c.JSON(http.StatusOK, resp)
}

func (goldController *GoldController) Sell(c echo.Context) error {
	params := new(model.SellGoldRequest)
	userID := c.Get("id").(int)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.BasicResp{Message: err.Error()})
	}

	params.ID = id
	params.UserID = userID

	err = c.Bind(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.BasicResp{Message: err.Error()})
	}

	err = validateSellUserGold(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.BasicResp{Message: err.Error()})
	}

	err = goldController.GoldService.Sell(params)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.BasicResp{Message: err.Error()})
	}

	resp := model.BasicResp{
		Message: "Success",
	}

	return c.JSON(http.StatusOK, resp)
}

func (goldController *GoldController) Dashboard(c echo.Context) error {
	params := new(model.DashboardUserGoldRequest)
	userID := c.Get("id").(int)
	params.UserID = userID

	err := c.Bind(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.BasicResp{Message: err.Error()})
	}

	dashboard, err := goldController.GoldService.Dashboard(params)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.BasicResp{Message: err.Error()})
	}

	resp := model.BasicResp{
		Message: "Success",
		Data:    dashboard,
	}

	return c.JSON(http.StatusOK, resp)
}
