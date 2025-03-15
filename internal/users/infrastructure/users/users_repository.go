package users

import (
	"errors"

	"github.com/enteresanlikk/go-modular-monolith/internal/users/domain/users"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrEmailAlreadyExist = errors.New("email already exists")
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) users.UserRepository {
	// Auto-migrate the users table
	db.AutoMigrate(&users.User{})

	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(user *users.User) error {
	// Check if email already exists
	var existingUser users.User
	result := r.db.Where("email = ?", user.Email).First(&existingUser)
	if result.Error == nil {
		return ErrEmailAlreadyExist
	} else if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}

	// Create new user
	result = r.db.Create(user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *UserRepository) FindByEmail(email string) (*users.User, error) {
	var user users.User
	result := r.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, result.Error
	}

	return &user, nil
}

func (r *UserRepository) FindByID(id uuid.UUID) (*users.User, error) {
	var user users.User
	result := r.db.First(&user, "id = ?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, result.Error
	}

	return &user, nil
}
