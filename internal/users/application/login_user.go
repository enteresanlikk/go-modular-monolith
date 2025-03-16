package usersApplication

import (
	usersDomain "github.com/enteresanlikk/go-modular-monolith/internal/users/domain"
	"golang.org/x/crypto/bcrypt"
)

type LoginUserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type TokenResponse = usersDomain.TokenPair

func (s *UserService) Login(req *LoginUserRequest) (*TokenResponse, error) {
	user, err := s.repo.FindByEmail(req.Email)
	if err != nil {
		return nil, usersDomain.ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, usersDomain.ErrInvalidCredentials
	}

	claims := map[string]interface{}{
		"userId": user.ID.String(),
	}

	tokenPair, err := s.tokenService.GenerateTokenPair(claims)
	if err != nil {
		return nil, err
	}

	return tokenPair, nil
}
