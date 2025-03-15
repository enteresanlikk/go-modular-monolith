package login_user

import (
	"time"

	"github.com/enteresanlikk/go-modular-monolith/internal/users/domain/users"
	"golang.org/x/crypto/bcrypt"
)

type LoginUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginResponse struct {
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

func (s *UserService) Login(req *LoginUserRequest) (*LoginResponse, error) {
	user, err := s.repo.FindByEmail(req.Email)
	if err != nil {
		return nil, users.ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, users.ErrInvalidCredentials
	}

	accessToken := "dummy-token"  // TODO: Implement proper JWT token generation
	refreshToken := "dummy-token" // TODO: Implement proper JWT token generation
	expiresAt := time.Now().Add(time.Hour * 24).Unix()

	return &LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    expiresAt,
	}, nil
}
