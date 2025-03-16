package users_presentation

import (
	users_application "github.com/enteresanlikk/go-modular-monolith/internal/users/application"
)

type UsersHandler struct {
	registerService *users_application.UserService
	loginService    *users_application.UserService
}

func NewUsersHandler(registerService *users_application.UserService, loginService *users_application.UserService) *UsersHandler {
	return &UsersHandler{
		registerService: registerService,
		loginService:    loginService,
	}
}
