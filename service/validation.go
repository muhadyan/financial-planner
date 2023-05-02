package service

import (
	"github.com/muhadyan/financial-planner/model"
	"github.com/muhadyan/financial-planner/utils"
	"golang.org/x/crypto/bcrypt"
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

func (c *UserService) validateLogIn(params *model.LogInRequest) (*model.User, error) {
	user, err := c.UserRepository.GetUser(&model.User{
		Username: params.Username,
		IsActive: true,
	})
	if err != nil {
		return nil, err
	}
	if user == nil {
		user, err = c.UserRepository.GetUser(&model.User{
			Email:    params.Username,
			IsActive: true,
		})
		if err != nil {
			return nil, err
		}
		if user == nil {
			return nil, utils.ErrUserNotExist
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(params.Password))
	if err != nil {
		return nil, utils.ErrWrongPassword
	}

	return user, nil
}

func (c *GoldService) validateCreateUserGold(params *model.CreateUserGoldRequest) error {
	user, err := c.UserRepository.GetUser(&model.User{
		ID:       uint(params.UserID),
		IsActive: true,
	})
	if err != nil {
		return err
	}

	if user == nil {
		return utils.ErrUserNotExist
	}

	return nil
}

func (c *GoldService) validateGetUnrealized(params *model.GetUnrealizedRequest) error {
	user, err := c.UserRepository.GetUser(&model.User{
		ID:       uint(params.UserID),
		IsActive: true,
	})
	if err != nil {
		return err
	}

	if user == nil {
		return utils.ErrUserNotExist
	}

	return nil
}
