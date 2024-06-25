package common

import "errors"

var (
	ErrMissingUsername    = errors.New("missing username")
	ErrMissingTokenString = errors.New("missing token string")
)
