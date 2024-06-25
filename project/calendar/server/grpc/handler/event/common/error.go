package common

import "errors"

var (
	ErrEmptyName        = errors.New("name is empty")
	ErrEmptyDescription = errors.New("description is empty")
)
