package auth

import "errors"

var (
	ErrUnexpectedSigningMethod = errors.New("unexpected signing method")
	ErrInvalidSessionToken     = errors.New("invalid session token")
)
