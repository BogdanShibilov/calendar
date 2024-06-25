package common

import "errors"

var (
	ErrMissingUsername = errors.New("username is missing")
	ErrMissingPassword = errors.New("password is missing")
)
