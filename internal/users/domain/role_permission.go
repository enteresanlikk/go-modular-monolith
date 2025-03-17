package usersDomain

import (
	commonDomain "github.com/enteresanlikk/go-modular-monolith/internal/common/domain"
	"github.com/google/uuid"
)

type RolePermission struct {
	commonDomain.BaseEntity

	RoleID       string `json:"roleId" gorm:"not null"`
	PermissionID string `json:"permissionId" gorm:"not null"`
}

func (RolePermission) TableName() string {
	return "users.role_permissions"
}

func NewRolePermission(roleID, permissionID string) (*RolePermission, error) {
	rp := &RolePermission{
		BaseEntity: commonDomain.BaseEntity{
			ID: uuid.New(),
		},
		RoleID:       roleID,
		PermissionID: permissionID,
	}
	return rp, nil
}
