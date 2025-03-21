package tenantsApplication

import (
	tenantsDomain "github.com/enteresanlikk/go-modular-monolith/internal/tenants/domain"
)

type CreateTenantRequest struct {
	Alias string `json:"alias" validate:"required"`
	Name  string `json:"name" validate:"required"`
}

func (s *TenantService) CreateTenant(req *CreateTenantRequest) (*TenantResponse, error) {
	tenant, err := tenantsDomain.NewTenant(req.Alias, req.Name)
	if err != nil {
		return nil, err
	}

	if err := s.tenantRepository.Create(tenant); err != nil {
		return nil, err
	}

	return (&TenantResponse{}).FromTenant(tenant), nil
}
