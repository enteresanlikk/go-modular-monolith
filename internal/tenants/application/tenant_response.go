package tenantsApplication

import (
	"time"

	tenantsDomain "github.com/enteresanlikk/go-modular-monolith/internal/tenants/domain"
	"github.com/google/uuid"
)

type TenantResponse struct {
	ID        uuid.UUID `json:"id"`
	Alias     string    `json:"alias"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (t *TenantResponse) FromTenant(tenant *tenantsDomain.Tenant) *TenantResponse {
	return &TenantResponse{
		ID:        tenant.ID,
		Alias:     tenant.Alias,
		Name:      tenant.Name,
		CreatedAt: tenant.CreatedAt,
		UpdatedAt: tenant.UpdatedAt,
	}
}
