package users_domain

import (
	common_domain "github.com/enteresanlikk/go-modular-monolith/internal/common/domain"
	"github.com/google/uuid"
)

type UserRole struct {
	common_domain.Entity

	UserID string `json:"userId" gorm:"not null"`
	RoleID string `json:"roleId" gorm:"not null"`
}

func (UserRole) TableName() string {
	return "users.user_roles"
}

func NewUserRole(userID, roleID string) (*UserRole, error) {
	ur := &UserRole{
		Entity: common_domain.Entity{
			ID: uuid.New(),
		},
		UserID: userID,
		RoleID: roleID,
	}
	return ur, nil
}
