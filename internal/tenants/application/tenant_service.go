package tenantsApplication

import tenantsDomain "github.com/enteresanlikk/go-modular-monolith/internal/tenants/domain"

type TenantService struct {
	tenantRepository tenantsDomain.TenantRepository
}

func NewTenantService(tenantRepository tenantsDomain.TenantRepository) *TenantService {
	return &TenantService{
		tenantRepository: tenantRepository,
	}
}
