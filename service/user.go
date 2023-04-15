package service

import (
	"fmt"
	"time"

	"github.com/muhadyan/financial-planner/config"
	"github.com/muhadyan/financial-planner/model"
	"github.com/muhadyan/financial-planner/repository"
	"github.com/muhadyan/financial-planner/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepository repository.UserRepository
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
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
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
			VerifyAccountLink: fmt.Sprintf("%s/v1/user/verify?userID=%d&userName=%s", config.GetConfig().BaseUrl, user.ID, user.Username),
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
