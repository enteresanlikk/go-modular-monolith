package users_application

import (
	"os"
	"strconv"
	"strings"
	"time"

	common_domain "github.com/enteresanlikk/go-modular-monolith/internal/common/domain"
	users_domain "github.com/enteresanlikk/go-modular-monolith/internal/users/domain"
	"github.com/google/uuid"
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

	user := &users_domain.User{
		Entity: common_domain.Entity{
			ID: uuid.New(),
		},
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Username:  strings.ToUpper(req.Email),
		Email:     req.Email,
		Password:  string(hashedPassword),
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
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
