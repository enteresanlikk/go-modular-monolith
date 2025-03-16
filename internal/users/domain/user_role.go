package usersDomain

import (
	commonDomain "github.com/enteresanlikk/go-modular-monolith/internal/common/domain"
	"github.com/google/uuid"
)

type UserRole struct {
	commonDomain.Entity

	UserID string `json:"userId" gorm:"not null"`
	RoleID string `json:"roleId" gorm:"not null"`
}

func (UserRole) TableName() string {
	return "users.user_roles"
}

func NewUserRole(userID, roleID string) (*UserRole, error) {
	ur := &UserRole{
		Entity: commonDomain.Entity{
			ID: uuid.New(),
		},
		UserID: userID,
		RoleID: roleID,
	}
	return ur, nil
}
