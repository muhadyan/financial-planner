package service

import (
	"fmt"

	"github.com/muhadyan/financial-planner/config"
	"github.com/muhadyan/financial-planner/model"
	"github.com/muhadyan/financial-planner/repository"
	"github.com/muhadyan/financial-planner/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepository repository.UserRepository
	TimeRepository repository.TimeRepository
}

func (c *UserService) SignUp(params *model.SignUpRequest) (*model.SignUpResponse, error) {
	err := c.validateSignUp(params)
	if err != nil {
		return nil, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(params.Password), 10)
	if err != nil {
		return nil, err
	}

	user := model.User{
		Username:  params.Username,
		Password:  string(hashedPassword),
		Email:     params.Email,
		Fullname:  params.Fullname,
		CreatedAt: c.TimeRepository.TimeNow(),
		UpdatedAt: c.TimeRepository.TimeNow(),
	}
	_, err = c.UserRepository.InsertUser(&user)
	if err != nil {
		return nil, err
	}

	go func() {
		templateEmail := "./template/verify_account.html"
		sendRequest := model.SendMail{
			SendTo:            user.Email,
			Username:          user.Username,
			UserFullname:      user.Fullname,
			VerifyAccountLink: fmt.Sprintf("%s/v1/user/verify?user_id=%d&username=%s", config.GetConfig().BaseUrl, user.ID, user.Username),
		}
		subject := fmt.Sprintf("Account Verification For %s!", user.Fullname)
		utils.SendMail(templateEmail, sendRequest, subject)
	}()

	resp := model.SignUpResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Fullname: user.Fullname,
	}

	return &resp, nil
}

func (c *UserService) Verify(params *model.VerifyRequest) error {
	err := c.validateVerify(params)
	if err != nil {
		return err
	}

	user := model.User{
		ID:       uint(params.UserID),
		IsActive: true,
	}
	_, err = c.UserRepository.UpdateUser(&user)
	if err != nil {
		return err
	}

	return nil
}
