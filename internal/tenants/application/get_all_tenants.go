package tenantsApplication

type GetAllTenantsRequest struct{}

func (s *TenantService) GetAllTenants(req *GetAllTenantsRequest) ([]*TenantResponse, error) {
	tenants, err := s.tenantRepository.FindAll()
	if err != nil {
		return nil, err
	}

	tenantsResponse := make([]*TenantResponse, len(tenants))
	for i, tenant := range tenants {
		tenantsResponse[i] = (&TenantResponse{}).FromTenant(tenant)
	}

	return tenantsResponse, nil
}
