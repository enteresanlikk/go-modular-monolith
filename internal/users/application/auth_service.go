package application

import (
	"errors"
	"time"

	"github.com/enteresanlikk/go-modular-monolith/internal/users/domain"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrPasswordMismatch   = errors.New("password and confirm password do not match")
)

type AuthService interface {
	Register(req *RegisterRequest) (*AuthResponse, error)
	Login(req *LoginRequest) (*AuthResponse, error)
}

type authService struct {
	userRepo domain.UserRepository
}

func NewAuthService(userRepo domain.UserRepository) AuthService {
	return &authService{
		userRepo: userRepo,
	}
}

func (s *authService) Register(req *RegisterRequest) (*AuthResponse, error) {
	if req.Password != req.ConfirmPassword {
		return nil, ErrPasswordMismatch
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		Email:     req.Email,
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	token := "dummy-token" // TODO: Implement proper JWT token generation

	return &AuthResponse{
		Token: token,
		User:  *user,
	}, nil
}

func (s *authService) Login(req *LoginRequest) (*AuthResponse, error) {
	user, err := s.userRepo.FindByEmail(req.Email)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, ErrInvalidCredentials
	}

	token := "dummy-token" // TODO: Implement proper JWT token generation

	return &AuthResponse{
		Token: token,
		User:  *user,
	}, nil
}
