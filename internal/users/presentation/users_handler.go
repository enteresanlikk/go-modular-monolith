package usersPresentation

import (
	usersApplication "github.com/enteresanlikk/go-modular-monolith/internal/users/application"
)

type UsersHandler struct {
	registerService *usersApplication.UserService
	loginService    *usersApplication.UserService
}

func NewUsersHandler(registerService *usersApplication.UserService, loginService *usersApplication.UserService) *UsersHandler {
	return &UsersHandler{
		registerService: registerService,
		loginService:    loginService,
	}
}
