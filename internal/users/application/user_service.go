package users_application

import users_domain "github.com/enteresanlikk/go-modular-monolith/internal/users/domain"

type UserService struct {
	repo users_domain.UserRepository
}

func NewUserService(repo users_domain.UserRepository) *UserService {
	return &UserService{repo: repo}
}
