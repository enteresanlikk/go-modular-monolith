package tenantsApplication

import "github.com/google/uuid"

type GetTenantByIdRequest struct {
	ID uuid.UUID `json:"id" validate:"required"`
}

func (s *TenantService) GetTenantById(req *GetTenantByIdRequest) (*TenantResponse, error) {
	tenant, err := s.tenantRepository.FindByID(req.ID)
	if err != nil {
		return nil, err
	}

	return (&TenantResponse{}).FromTenant(tenant), nil
}
