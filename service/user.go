package service

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/muhadyan/financial-planner/config"
	"github.com/muhadyan/financial-planner/model"
	"github.com/muhadyan/financial-planner/repository"
	"github.com/muhadyan/financial-planner/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepository     repository.UserRepository
	TimeRepository     repository.TimeRepository
	RoleRepository     repository.RoleRepository
	UserRoleRepository repository.UserRoleRepository
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
		Username: params.Username,
		Password: string(hashedPassword),
		Email:    params.Email,
		Fullname: params.Fullname,
	}
	_, err = c.UserRepository.InsertUser(&user)
	if err != nil {
		return nil, err
	}

	role, err := c.RoleRepository.GetRole(&model.Role{Name: utils.RoleUser})
	if err != nil {
		return nil, err
	}
	if role == nil {
		return nil, utils.ErrRoleNameUserNotExist
	}

	userRole := model.UserRole{
		UserID: int(user.ID),
		RoleID: int(role.ID),
	}
	_, err = c.UserRoleRepository.InsertUserRole(&userRole)
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

func (c *UserService) LogIn(params *model.LogInRequest) (*model.LogInResponse, error) {
	roles := []model.Role{}
	roleNames := []string{}

	user, err := c.validateLogIn(params)
	if err != nil {
		return nil, err
	}

	userRoles, err := c.UserRoleRepository.GetUserRoles(&model.UserRole{UserID: int(user.ID)})
	if err != nil {
		return nil, err
	}

	if len(userRoles) > 0 {
		roleIDs := make([]int, 0)

		for _, userRole := range userRoles {
			roleIDs = append(roleIDs, userRole.RoleID)
		}

		roles, err = c.RoleRepository.GetRolesByIDs(roleIDs)
		if err != nil {
			return nil, err
		}
	}

	for _, role := range roles {
		roleNames = append(roleNames, role.Name)
	}

	now := c.TimeRepository.TimeNow()
	payload := model.JWTPayload{
		ID:       int(user.ID),
		Username: user.Username,
		Email:    user.Email,
		Fullname: user.Fullname,
		IsActive: user.IsActive,
		Roles:    roleNames,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  now.Unix(),
			ExpiresAt: now.Add(time.Hour * 24).Unix(),
		},
	}
	token, err := utils.GenerateJWT(payload)
	if err != nil {
		return nil, err
	}

	_, err = c.UserRepository.UpdateUser(&model.User{
		ID:    user.ID,
		Token: &token,
	})
	if err != nil {
		return nil, err
	}

	resp := model.LogInResponse{
		ID:    user.ID,
		Token: token,
	}

	return &resp, nil
}
