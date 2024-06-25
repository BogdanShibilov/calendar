package jwt

import "errors"

var (
	ErrExpiredToken       = errors.New("expired token")
	ErrMalformedToken     = errors.New("malformed token")
	ErrUnknownClaims      = errors.New("unknown claims")
	ErrNoSuchTokenForUser = errors.New("no such token for given user exist")
)
