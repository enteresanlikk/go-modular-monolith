package usersInfrastructure

import (
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

	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(user *usersDomain.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) FindByEmail(email string) (*usersDomain.User, error) {
	var user usersDomain.User
	result := r.db.Where("email = ?", email).First(&user)
	return &user, result.Error
}

func (r *UserRepository) FindByID(id uuid.UUID) (*usersDomain.User, error) {
	var user usersDomain.User
	result := r.db.Where("id = ?", id).First(&user)
	return &user, result.Error
}
