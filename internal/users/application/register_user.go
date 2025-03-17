package usersApplication

import (
	"errors"
	"os"
	"strconv"
	"strings"

	usersDomain "github.com/enteresanlikk/go-modular-monolith/internal/users/domain"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type RegisterUserRequest struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required"`
	Password  string `json:"password" validate:"required"`
}

func (s *UserService) Register(req *RegisterUserRequest) (*TokenResponse, error) {
	findedUser, err := s.repo.FindByEmail(req.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if findedUser != nil {
		return nil, usersDomain.ErrEmailAlreadyExist
	}

	var bcryptCost, _ = strconv.Atoi(os.Getenv("BCRYPT_COST"))
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcryptCost)
	if err != nil {
		return nil, err
	}

	user, err := usersDomain.NewUser(
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
