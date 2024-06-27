package common

import "errors"

var (
	ErrRequestTimeout = errors.New("request timeout")
	ErrInternal       = errors.New("something went wrong")
)
