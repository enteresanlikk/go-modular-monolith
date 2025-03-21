package tenantsDomain

import (
	"github.com/google/uuid"
)

type TenantRepository interface {
	Create(tenant *Tenant) error
	Update(tenant *Tenant) error
	Delete(id uuid.UUID) error
	FindByID(id uuid.UUID) (*Tenant, error)
	FindByAlias(alias string) (*Tenant, error)
	FindAll() ([]*Tenant, error)
}
