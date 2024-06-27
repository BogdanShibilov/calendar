package refresh

import "errors"

var (
	ErrMissingRefreshToken = errors.New("refresh token required")
	ErrInvalidToken        = errors.New("invalid token")
	ErrUnauthenticated     = errors.New("wrong refresh token")
)
