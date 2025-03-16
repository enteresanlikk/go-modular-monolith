package application

import users "github.com/enteresanlikk/go-modular-monolith/internal/users/domain"

type UserService struct {
	repo users.UserRepository
}

func NewUserService(repo users.UserRepository) *UserService {
	return &UserService{repo: repo}
}
