package tenantsDomain

import (
	commonDomain "github.com/enteresanlikk/go-modular-monolith/internal/common/domain"
	"github.com/google/uuid"
)

type Tenant struct {
	commonDomain.BaseEntity

	Alias string `json:"alias" gorm:"uniqueIndex;not null"`
	Name  string `json:"name" gorm:"not null"`
}

func (t *Tenant) TableName() string {
	return "tenants.tenants"
}

func NewTenant(alias, name string) (*Tenant, error) {
	t := &Tenant{
		BaseEntity: commonDomain.BaseEntity{
			ID: uuid.New(),
		},
		Alias: alias,
		Name:  name,
	}

	return t, nil
}

func (t *Tenant) Update(alias, name string) (*Tenant, error) {
	t.Alias = alias
	t.Name = name
	return t, nil
}
