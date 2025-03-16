package users_infrastructure

import (
	"errors"

	users_domain "github.com/enteresanlikk/go-modular-monolith/internal/users/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) users_domain.UserRepository {
	db.Exec("CREATE SCHEMA IF NOT EXISTS users")

	db.AutoMigrate(&users_domain.User{})

	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(user *users_domain.User) error {
	var existingUser users_domain.User
	result := r.db.Where("email = ?", user.Email).First(&existingUser)
	if result.Error == nil {
		return users_domain.ErrEmailAlreadyExist
	} else if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}

	result = r.db.Create(user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *UserRepository) FindByEmail(email string) (*users_domain.User, error) {
	var user users_domain.User
	result := r.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, users_domain.ErrUserNotFound
		}
		return nil, result.Error
	}

	return &user, nil
}

func (r *UserRepository) FindByID(id uuid.UUID) (*users_domain.User, error) {
	var user users_domain.User
	result := r.db.First(&user, "id = ?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, users_domain.ErrUserNotFound
		}
		return nil, result.Error
	}

	return &user, nil
}
