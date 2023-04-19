package utils

import (
	"errors"
	"fmt"
)

var (
	ErrEmptyEmail    = errors.New("email cannot be empty")
	ErrEmptyFullname = errors.New("fullname cannot be empty")
	ErrEmptyPassword = errors.New("password cannot be empty")
	ErrEmptyUsername = errors.New("username cannot be empty")
	ErrEmptyBuyDate  = errors.New("gold buy date cannot be empty")

	ErrInvalidUserID   = errors.New("user id must be more than 0")
	ErrInvalidBuyPrice = errors.New("gold buy price must be more than 0")
	ErrInvalidBuyDate  = fmt.Errorf("gold buy date must be in %s format", FormatDate)

	ErrUsernameExist = errors.New("username already exist")
	ErrEmailExist    = errors.New("email already exist")

	ErrUserNotExist         = errors.New("user is not exist")
	ErrUserTokenNotExist    = errors.New("user token is not exist")
	ErrRoleNameUserNotExist = errors.New("role name user is not exist")

	ErrUserIDUsernameNotMatch = errors.New("user id and username do not match")
	ErrUserVerified           = errors.New("user is already verified")
	ErrWrongPassword          = errors.New("password is not correct")
)

var (
	RoleUser = "user"
)

const (
	FormatDate = `2006-01-02`
)
