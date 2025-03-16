package presentation

import (
	"github.com/enteresanlikk/go-modular-monolith/internal/users/application"
)

type AuthHandler struct {
	registerService *application.UserService
	loginService    *application.UserService
}

func NewAuthHandler(registerService *application.UserService, loginService *application.UserService) *AuthHandler {
	return &AuthHandler{
		registerService: registerService,
		loginService:    loginService,
	}
}
