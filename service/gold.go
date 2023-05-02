package service

import (
	"strconv"
	"time"

	"github.com/muhadyan/financial-planner/config"
	"github.com/muhadyan/financial-planner/model"
	"github.com/muhadyan/financial-planner/repository"
	"github.com/muhadyan/financial-planner/utils"
)

type GoldService struct {
	UserGoldRepository    repository.UserGoldRepository
	UserRepository        repository.UserRepository
	CurrentGoldRepository repository.CurrentGoldRepository
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

func (c *GoldService) GetUnrealized(params *model.GetUnrealizedRequest) ([]model.GetUnrealizedResponse, error) {
	resp := []model.GetUnrealizedResponse{}

	err := c.validateGetUnrealized(params)
	if err != nil {
		return nil, err
	}

	userGolds, err := c.UserGoldRepository.GetUserGolds(&model.UserGoldViewParams{
		UserID:        params.UserID,
		NullSellPrice: true,
		BasedFilter: model.BasedFilter{
			Limit: params.Limit,
			Page:  params.Page,
		},
	})
	if err != nil {
		return nil, err
	}

	if userGolds != nil {
		currentGold, err := c.CurrentGoldRepository.GetCurrentGold()
		if err != nil {
			return nil, err
		}

		currentPrice := currentGold.Data[0].Buy
		minNominalTax, err := strconv.ParseFloat(config.GetConfig().MinNominalTax, 64)
		if err != nil {
			return nil, err
		}

		// check users npwp
		pph22, err := strconv.ParseFloat(config.GetConfig().PPH22, 64)
		if err != nil {
			return nil, err
		}

		for _, userGold := range userGolds {
			var unrealized float64
			currentTotalPrice := currentPrice * userGold.Weight

			if currentTotalPrice > minNominalTax {
				pph22Nominal := currentTotalPrice * pph22 / 100
				unrealized = (currentTotalPrice - pph22Nominal) - userGold.BuyPrice
			} else {
				unrealized = currentTotalPrice - userGold.BuyPrice
			}

			resp = append(resp, model.GetUnrealizedResponse{
				ID:           int(userGold.ID),
				UserID:       userGold.UserID,
				Weight:       userGold.Weight,
				BuyPrice:     userGold.BuyPrice,
				Unrealized:   unrealized,
				Percent:      utils.RoundFloat(((unrealized / userGold.BuyPrice) * 100), 2),
				AvgBuyPrice:  utils.RoundFloat((userGold.BuyPrice / userGold.Weight), 2),
				CurrentPrice: currentPrice,
				BuyDate:      userGold.BuyDate.Format(utils.FormatDate),
			})
		}
	}

	return resp, nil
}
