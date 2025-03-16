package users_domain

import "errors"

var (
	ErrInvalidCredentials = errors.New("invalid_credentials")
	ErrPasswordMismatch   = errors.New("password_and_confirm_password_do_not_match")
	ErrUserNotFound       = errors.New("user_not_found")
	ErrEmailAlreadyExist  = errors.New("email_already_exists")
)
