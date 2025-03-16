package users_application

import (
	"os"
	"strconv"
	"strings"
	"time"

	users_domain "github.com/enteresanlikk/go-modular-monolith/internal/users/domain"
	"golang.org/x/crypto/bcrypt"
)

type RegisterUserRequest struct {
	FirstName       string `json:"firstName" binding:"required"`
	LastName        string `json:"lastName" binding:"required"`
	Email           string `json:"email" binding:"required,email"`
	Password        string `json:"password" binding:"required,min=6"`
	ConfirmPassword string `json:"confirmPassword" binding:"required,eqfield=Password"`
}

type RegisterUserResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	ExpiresAt    int64  `json:"expiresAt"`
}

func (s *UserService) Register(req *RegisterUserRequest) (*RegisterUserResponse, error) {
	if req.Password != req.ConfirmPassword {
		return nil, users_domain.ErrPasswordMismatch
	}

	var bcryptCost, _ = strconv.Atoi(os.Getenv("BCRYPT_COST"))
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcryptCost)
	if err != nil {
		return nil, err
	}

	user := &users_domain.User{
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

	return &RegisterUserResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    expiresAt,
	}, nil
}
