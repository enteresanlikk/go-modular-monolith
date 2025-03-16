package usersApplication

import (
	usersDomain "github.com/enteresanlikk/go-modular-monolith/internal/users/domain"
)

type UserService struct {
	repo         usersDomain.UserRepository
	tokenService usersDomain.TokenService
}

func NewUserService(repo usersDomain.UserRepository, tokenService usersDomain.TokenService) *UserService {
	return &UserService{
		repo:         repo,
		tokenService: tokenService,
	}
}
