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
