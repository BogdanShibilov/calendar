package add

import "errors"

var (
	ErrMissingName        = errors.New("event name is required")
	ErrMissingDescription = errors.New("event description is required")
	ErrEventAlreadyExists = errors.New("event already exists")
)
