package tenantsApplication

type GetTenantByAliasRequest struct {
	Alias string `json:"alias" validate:"required"`
}

func (s *TenantService) GetTenantByAlias(req *GetTenantByAliasRequest) (*TenantResponse, error) {
	tenant, err := s.tenantRepository.FindByAlias(req.Alias)
	if err != nil {
		return nil, err
	}

	return (&TenantResponse{}).FromTenant(tenant), nil
}
