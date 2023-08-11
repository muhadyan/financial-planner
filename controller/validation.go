package controller

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/muhadyan/financial-planner/model"
	"github.com/muhadyan/financial-planner/utils"
)

func validateSignUp(params *model.SignUpRequest) error {
	if params.Email == "" {
		return utils.ErrEmptyEmail
	}

	if params.Fullname == "" {
		return utils.ErrEmptyFullname
	}

	if params.Password == "" {
		return utils.ErrEmptyPassword
	}

	if params.Username == "" {
		return utils.ErrEmptyUsername
	}

	return nil
}

func validateVerify(params *model.VerifyRequest) error {
	if params.UserID <= 0 {
		return utils.ErrInvalidUserID
	}

	if params.Username == "" {
		return utils.ErrEmptyUsername
	}

	return nil
}

func validateLogIn(params *model.LogInRequest) error {
	if params.Username == "" {
		return utils.ErrEmptyUsername
	}

	if params.Password == "" {
		return utils.ErrEmptyPassword
	}

	return nil
}

func validateCreateUserGold(params *model.CreateUserGoldRequest) error {
	if params.Weight <= 0.0 {
		return utils.ErrInvalidWeight
	}

	if params.BuyPrice <= 0.0 {
		return utils.ErrInvalidBuyPrice
	}

	if params.BuyDate == "" {
		return utils.ErrEmptyBuyDate
	}

	err := validation.Validate(params.BuyDate, validation.Date(utils.FormatDate))
	if err != nil {
		return utils.ErrInvalidBuyDate
	}

	return nil
}

func validateUpdateUserGold(params *model.UpdateUserGoldRequest) error {
	if params.Weight <= 0.0 {
		return utils.ErrInvalidWeight
	}

	if params.BuyPrice <= 0.0 {
		return utils.ErrInvalidBuyPrice
	}

	if params.BuyDate == "" {
		return utils.ErrEmptyBuyDate
	}

	err := validation.Validate(params.BuyDate, validation.Date(utils.FormatDate))
	if err != nil {
		return utils.ErrInvalidBuyDate
	}

	return nil
}

func validateSellUserGold(params *model.SellGoldRequest) error {
	if params.SellPrice <= 0.0 {
		return utils.ErrInvalidSellPrice
	}

	if params.SellDate == "" {
		return utils.ErrEmptySellDate
	}

	err := validation.Validate(params.SellDate, validation.Date(utils.FormatDate))
	if err != nil {
		return utils.ErrInvalidSellDate
	}

	return nil
}

func validatePaginationReq(params *model.PagiantionReq) error {
	if params.Page < 1 {
		return utils.ErrInvalidPage
	}

	if params.Limit < 1 {
		return utils.ErrInvalidPage
	}

	return nil
}
