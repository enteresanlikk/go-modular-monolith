package users_domain

import (
	common_domain "github.com/enteresanlikk/go-modular-monolith/internal/common/domain"
	"github.com/google/uuid"
)

type RolePermission struct {
	common_domain.Entity

	RoleID       string `json:"roleId" gorm:"not null"`
	PermissionID string `json:"permissionId" gorm:"not null"`
}

func (RolePermission) TableName() string {
	return "users.role_permissions"
}

func NewRolePermission(roleID, permissionID string) (*RolePermission, error) {
	rp := &RolePermission{
		Entity: common_domain.Entity{
			ID: uuid.New(),
		},
		RoleID:       roleID,
		PermissionID: permissionID,
	}
	return rp, nil
}
