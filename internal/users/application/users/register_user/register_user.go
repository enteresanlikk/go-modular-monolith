package register_user

import (
	"time"

	"github.com/enteresanlikk/go-modular-monolith/internal/users/domain/users"
	"golang.org/x/crypto/bcrypt"
)

type RegisterUserRequest struct {
	Email           string `json:"email" binding:"required,email"`
	Password        string `json:"password" binding:"required,min=6"`
	ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=Password"`
}

type RegisterUserResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresAt    int64  `json:"expires_at"`
}

type UserService struct {
	repo users.UserRepository
}

func NewUserService(repo users.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Register(req *RegisterUserRequest) (*RegisterUserResponse, error) {
	if req.Password != req.ConfirmPassword {
		return nil, users.ErrPasswordMismatch
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &users.User{
		Email:     req.Email,
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
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
