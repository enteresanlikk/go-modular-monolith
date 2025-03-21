package tenantsApplication

import "github.com/google/uuid"

type DeleteTenantRequest struct {
	ID uuid.UUID `json:"id" validate:"required"`
}

func (s *TenantService) DeleteTenant(req *DeleteTenantRequest) error {
	tenant, err := s.tenantRepository.FindByID(req.ID)
	if err != nil {
		return err
	}

	return s.tenantRepository.Delete(tenant.ID)
}
