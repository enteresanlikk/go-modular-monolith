package tenantsDomain

import (
	commonDomain "github.com/enteresanlikk/go-modular-monolith/internal/common/domain"
	"github.com/google/uuid"
)

type TenantRole string

const (
	TenantRoleMember TenantRole = "member"
	TenantRoleAdmin  TenantRole = "admin"
)

type UserTenant struct {
	commonDomain.BaseEntity

	UserID   uuid.UUID  `json:"user_id" gorm:"type:uuid;not null"`
	TenantID uuid.UUID  `json:"tenant_id" gorm:"type:uuid;not null"`
	Role     TenantRole `json:"role" gorm:"not null;default:'member'"`
}

func (u *UserTenant) TableName() string {
	return "tenants.user_tenants"
}

func NewUserTenant(userID, tenantID uuid.UUID, role TenantRole) *UserTenant {
	return &UserTenant{
		BaseEntity: commonDomain.BaseEntity{
			ID: uuid.New(),
		},
		UserID:   userID,
		TenantID: tenantID,
		Role:     role,
	}
}
