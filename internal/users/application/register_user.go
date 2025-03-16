package users_application

import (
	"strings"
	"time"

	users_domain "github.com/enteresanlikk/go-modular-monolith/internal/users/domain"
)

type RegisterUserRequest struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required"`
	Password  string `json:"password" validate:"required"`
}

type RegisterUserResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	ExpiresAt    int64  `json:"expiresAt"`
}

func (s *UserService) Register(req *RegisterUserRequest) (*RegisterUserResponse, error) {
	user := &users_domain.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Username:  strings.ToUpper(req.Email),
		Email:     req.Email,
		Password:  req.Password,
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	accessToken := "dummy-token"  // TODO: Implement proper JWT token generation
	refreshToken := "dummy-token" // TODO: Implement proper JWT token generation
	expiresAt := time.Now().Add(time.Hour * 24).Unix()

	return &RegisterUserResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    expiresAt,
	}, nil
}
