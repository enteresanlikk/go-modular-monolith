package usersInfrastructure

import (
	"errors"

	usersDomain "github.com/enteresanlikk/go-modular-monolith/internal/users/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) usersDomain.UserRepository {
	db.Exec("CREATE SCHEMA IF NOT EXISTS users")

	db.AutoMigrate(&usersDomain.User{})
	db.AutoMigrate(&usersDomain.Role{})
	db.AutoMigrate(&usersDomain.Permission{})
	db.AutoMigrate(&usersDomain.RolePermission{})
	db.AutoMigrate(&usersDomain.UserRole{})

	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(user *usersDomain.User) error {
	var existingUser usersDomain.User
	result := r.db.Where("email = ?", user.Email).First(&existingUser)
	if result.Error == nil {
		return usersDomain.ErrEmailAlreadyExist
	} else if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}

	result = r.db.Create(user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *UserRepository) FindByEmail(email string) (*usersDomain.User, error) {
	var user usersDomain.User
	result := r.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, usersDomain.ErrUserNotFound
		}
		return nil, result.Error
	}

	return &user, nil
}

func (r *UserRepository) FindByID(id uuid.UUID) (*usersDomain.User, error) {
	var user usersDomain.User
	result := r.db.First(&user, "id = ?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, usersDomain.ErrUserNotFound
		}
		return nil, result.Error
	}

	return &user, nil
}
