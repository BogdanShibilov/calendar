package signup

import "errors"

var (
	ErrMissingUsername       = errors.New("username is required")
	ErrMissingPassword       = errors.New("password is required")
	ErrUsernameAlreadyExists = errors.New("user with this username already exists")
)
