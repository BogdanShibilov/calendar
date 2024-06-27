package signin

import "errors"

var (
	ErrMissingUsername  = errors.New("username is required")
	ErrMissingPassword  = errors.New("password is required")
	ErrWrongCredentials = errors.New("wrong credentials")
)
