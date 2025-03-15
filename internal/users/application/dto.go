package application

import "github.com/enteresanlikk/go-modular-monolith/internal/users/domain"

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type RegisterRequest struct {
	Email           string `json:"email" binding:"required,email"`
	Password        string `json:"password" binding:"required,min=6"`
	ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=Password"`
}

type AuthResponse struct {
	Token string      `json:"token"`
	User  domain.User `json:"user"`
}
