package user

import "errors"

var (
	ErrEmptyUsername = errors.New("username is empty")
	ErrEmptyPassword = errors.New("password is empty")
)
