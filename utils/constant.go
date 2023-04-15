package utils

import (
	"errors"
)

var (
	ErrEmptyEmail    = errors.New("email cannot be empty")
	ErrEmptyFullname = errors.New("fullname cannot be empty")
	ErrEmptyPassword = errors.New("password cannot be empty")
	ErrEmptyUsername = errors.New("username cannot be empty")

	ErrUsernameExist = errors.New("username already exist")
	ErrEmailExist = errors.New("email already exist")
)
