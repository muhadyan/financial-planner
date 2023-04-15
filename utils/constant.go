package utils

import (
	"errors"
)

var (
	ErrEmptyEmail    = errors.New("email cannot be empty")
	ErrEmptyFullname = errors.New("fullname cannot be empty")
	ErrEmptyPassword = errors.New("password cannot be empty")
	ErrEmptyUsername = errors.New("username cannot be empty")
	ErrEmptyUserID   = errors.New("user id cannot be empty")

	ErrUsernameExist = errors.New("username already exist")
	ErrEmailExist    = errors.New("email already exist")

	ErrUserNotExist = errors.New("user is not exist")

	ErrUserIDUsernameNotMatch = errors.New("user id and username do not match")
	ErrUserVerified           = errors.New("user is already verified")
)
