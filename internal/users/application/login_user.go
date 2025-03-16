package users_application

import (
	"time"

	users_domain "github.com/enteresanlikk/go-modular-monolith/internal/users/domain"
	"golang.org/x/crypto/bcrypt"
)

type LoginUserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

func (s *UserService) Login(req *LoginUserRequest) (*TokenResponse, error) {
	user, err := s.repo.FindByEmail(req.Email)
	if err != nil {
		return nil, users_domain.ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, users_domain.ErrInvalidCredentials
	}

	accessToken := "dummy-token"  // TODO: Implement proper JWT token generation
	refreshToken := "dummy-token" // TODO: Implement proper JWT token generation
	expiresAt := time.Now().Add(time.Hour * 24).Unix()

	return &TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    expiresAt,
	}, nil
}
