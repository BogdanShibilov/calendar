package event

import "errors"

var (
	ErrEmptyName        = errors.New("name is empty")
	ErrInvalidId        = errors.New("invalid event id")
	ErrEmptyDescription = errors.New("description is empty")
)
