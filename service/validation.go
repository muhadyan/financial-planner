package service

import (
	"github.com/muhadyan/financial-planner/model"
	"github.com/muhadyan/financial-planner/utils"
)

func (c *UserService) validateSignUp(params *model.SignUpRequest) error {
	username, err := c.UserRepository.GetUser(&model.User{Username: params.Username})
	if err != nil {
		return err
	}
	if username != nil {
		return utils.ErrUsernameExist
	}

	email, err := c.UserRepository.GetUser(&model.User{Email: params.Email})
	if err != nil {
		return err
	}
	if email != nil {
		return utils.ErrEmailExist
	}

	return nil
}

func (c *UserService) validateVerify(params *model.VerifyRequest) error {
	user, err := c.UserRepository.GetUser(&model.User{ID: uint(params.UserID)})
	if err != nil {
		return err
	}
	if user == nil {
		return utils.ErrUserNotExist
	}

	if user.Username != params.Username {
		return utils.ErrUserIDUsernameNotMatch
	}

	if user.IsActive {
		return utils.ErrUserVerified
	}

	return nil
}
