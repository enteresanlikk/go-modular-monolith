package users_domain

import "errors"

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrPasswordMismatch   = errors.New("password and confirm password do not match")
)
