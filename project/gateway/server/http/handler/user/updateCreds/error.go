package updateCreds

import "errors"

var (
	ErrAccessTokenForNonExistingUser = errors.New("access token for non existing user")
	ErrMissingNewUsername            = errors.New("new username is required")
	ErrMissingNewPassword            = errors.New("new password is required")
)
