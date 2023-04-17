package controller

import (
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
	if params.UserID == 0 {
		return utils.ErrEmptyUserID
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
