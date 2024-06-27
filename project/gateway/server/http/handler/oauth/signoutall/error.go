package signoutall

import "errors"

var (
	ErrMissingAccessToken = errors.New("access token is required")
	ErrInvalidToken       = errors.New("invalid token")
)
