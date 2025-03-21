package tenantsApplication

import (
	"github.com/google/uuid"
)

type UpdateTenantRequest struct {
	ID    uuid.UUID `json:"id" validate:"required"`
	Alias string    `json:"alias" validate:"required"`
	Name  string    `json:"name" validate:"required"`
}

func (s *TenantService) UpdateTenant(req *UpdateTenantRequest) (*TenantResponse, error) {
	tenant, err := s.tenantRepository.FindByID(req.ID)
	if err != nil {
		return nil, err
	}

	tenant, err = tenant.Update(req.Alias, req.Name)
	if err != nil {
		return nil, err
	}

	if err := s.tenantRepository.Update(tenant); err != nil {
		return nil, err
	}

	return (&TenantResponse{}).FromTenant(tenant), nil
}
