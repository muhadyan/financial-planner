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

func (c *GoldService) FindAll(params *model.FindAllUserGoldRequest) ([]model.UserGold, error) {
	err := c.validateExistUser(params.UserID)
	if err != nil {
		return nil, err
	}

	resp, err := c.UserGoldRepository.GetUserGolds(&model.UserGoldViewParams{
		UserID: params.UserID,
		BasedFilter: model.BasedFilter{
			Limit: params.Limit,
			Page:  params.Page,
		},
	})
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *GoldService) Find(id, userID int) (*model.UserGold, error) {
	userGold, err := c.UserGoldRepository.GetUserGold(&model.UserGold{
		ID: uint(id),
	})
	if err != nil {
		return nil, err
	}

	err = c.validateFindUserGold(userGold, userID)
	if err != nil {
		return nil, err
	}

	return userGold, nil
}

func (c *GoldService) Update(params *model.UpdateUserGoldRequest) error {
	err := c.validateUpdateUserGold(params)
	if err != nil {
		return err
	}

	buyDate, err := time.Parse(utils.FormatDate, params.BuyDate)
	if err != nil {
		return err
	}

	userGold := model.UserGold{
		ID:       uint(params.ID),
		Weight:   params.Weight,
		BuyPrice: params.BuyPrice,
		BuyDate:  buyDate,
	}
	_, err = c.UserGoldRepository.UpdateUserGold(&userGold)
	if err != nil {
		return err
	}

	return nil
}

func (c *GoldService) Delete(id, userID int) error {
	err := c.validateDeleteUserGold(id, userID)
	if err != nil {
		return err
	}

	err = c.UserGoldRepository.DeleteUserGold(id)
	if err != nil {
		return err
	}

	return nil
}

func (c *GoldService) GetUnrealized(params *model.GetUnrealizedRequest) ([]model.GetUnrealizedResponse, error) {
	resp := []model.GetUnrealizedResponse{}

	err := c.validateExistUser(params.UserID)
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
				Percent:      utils.RoundFloat((unrealized/userGold.BuyPrice)*100, 2),
				AvgBuyPrice:  utils.RoundFloat(userGold.BuyPrice/userGold.Weight, 2),
				CurrentPrice: currentPrice,
				BuyDate:      userGold.BuyDate.Format(utils.FormatDate),
			})
		}
	}

	return resp, nil
}

func (c *GoldService) Sell(params *model.SellGoldRequest) error {
	err := c.validateSellGold(params)
	if err != nil {
		return err
	}

	sellDate, err := time.Parse(utils.FormatDate, params.SellDate)
	if err != nil {
		return err
	}

	userGold := model.UserGold{
		ID:        uint(params.ID),
		SellPrice: &params.SellPrice,
		SellDate:  &sellDate,
	}
	_, err = c.UserGoldRepository.UpdateUserGold(&userGold)
	if err != nil {
		return err
	}

	return nil
}

func (c *GoldService) Dashboard(params *model.DashboardUserGoldRequest) (*model.DashboardUserGoldResponse, error) {
	var (
		totalWeight           float64
		totalSoldWeight       float64
		totalBuyPrice         float64
		totalSellPrice        float64
		highestRealize        float64
		highestRealizePercent float64
		lowestRealize         float64
		lowestRealizePercent  float64
		totalRealize          float64
	)

	err := c.validateDashboardUserGold(params)
	if err != nil {
		return nil, err
	}

	userGolds, err := c.UserGoldRepository.GetAllUserGolds(&model.UserGold{
		UserID: params.UserID,
	})
	if err != nil {
		return nil, err
	}

	for _, userGold := range userGolds {
		totalWeight += userGold.Weight
		totalBuyPrice += userGold.BuyPrice

		if userGold.SellPrice != nil {
			totalSoldWeight += userGold.Weight
			totalSellPrice += *userGold.SellPrice

			realize := *userGold.SellPrice - userGold.BuyPrice

			if realize > highestRealize {
				highestRealize = realize
				highestRealizePercent = ((*userGold.SellPrice - userGold.BuyPrice) / userGold.BuyPrice) * 100
			}

			if realize > lowestRealize {
				lowestRealize = realize
				lowestRealizePercent = ((*userGold.SellPrice - userGold.BuyPrice) / userGold.BuyPrice) * 100
			}

			totalRealize += realize

		}
	}

	// still not complete for some response
	resp := &model.DashboardUserGoldResponse{
		UserID:                       params.UserID,
		TotalWeight:                  totalWeight,
		AvgBuyPricePerGram:           totalBuyPrice / totalWeight,
		AvgSellPricePerGram:          totalSellPrice / totalSoldWeight,
		HighestRealizedMargin:        highestRealize,
		HighestRealizedMarginPercent: highestRealizePercent,
		LowestRealizedMargin:         lowestRealize,
		LowestRealizedMarginPercent:  lowestRealizePercent,
		AvgRealizedMargin:            totalRealize / totalSoldWeight,
	}

	return resp, nil
}
