package service

import (
	"time"

	"github.com/muhadyan/financial-planner/model"
	"github.com/muhadyan/financial-planner/repository"
	"github.com/muhadyan/financial-planner/utils"
)

type GoldService struct {
	UserGoldRepository repository.UserGoldRepository
	UserRepository     repository.UserRepository
}

func (c *GoldService) Create(params *model.CreateUserGoldRequest) (*model.CreateUserGoldResponse, error) {
	err := c.validateCreateUserGold(params)
	if err != nil {
		return nil, err
	}

	buyDate, err := time.Parse(utils.FormatDate, params.BuyDate)
	if err != nil {
		return nil, err
	}

	userGold, err := c.UserGoldRepository.InsertUserGold(&model.UserGold{
		UserID:   params.UserID,
		Weight:   params.Weight,
		BuyPrice: params.BuyPrice,
		BuyDate:  buyDate,
	})
	if err != nil {
		return nil, err
	}

	resp := model.CreateUserGoldResponse{
		ID:       int(userGold.ID),
		UserID:   userGold.UserID,
		Weight:   userGold.Weight,
		BuyPrice: userGold.BuyPrice,
		BuyDate:  userGold.BuyDate.String(),
	}

	return &resp, nil
}
