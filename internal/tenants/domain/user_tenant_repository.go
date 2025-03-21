package tenantsDomain

import "github.com/google/uuid"

type UserTenantRepository interface {
	Create(userTenant *UserTenant) error
	Update(userTenant *UserTenant) error
	Delete(id uuid.UUID) error
	FindByID(id uuid.UUID) (*UserTenant, error)
	FindByUserID(userID uuid.UUID) ([]*UserTenant, error)
	FindByTenantID(tenantID uuid.UUID) ([]*UserTenant, error)
}
