package users_application

import (
	"os"
	"strconv"
	"strings"

	users_domain "github.com/enteresanlikk/go-modular-monolith/internal/users/domain"
	"golang.org/x/crypto/bcrypt"
)

type RegisterUserRequest struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required"`
	Password  string `json:"password" validate:"required"`
}

func (s *UserService) Register(req *RegisterUserRequest) (*TokenResponse, error) {
	var bcryptCost, _ = strconv.Atoi(os.Getenv("BCRYPT_COST"))
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcryptCost)
	if err != nil {
		return nil, err
	}

	user, err := users_domain.NewUser(
		req.FirstName,
		req.LastName,
		strings.ToUpper(req.Email),
		req.Email,
		string(hashedPassword),
	)
	if err != nil {
		return nil, err
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	claims := map[string]interface{}{
		"userId": user.ID.String(),
	}

	return s.tokenService.GenerateTokenPair(claims)
}
