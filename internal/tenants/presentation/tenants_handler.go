package tenantsPresentation

import tenantsApplication "github.com/enteresanlikk/go-modular-monolith/internal/tenants/application"

type TenantsHandler struct {
	createTenantService     *tenantsApplication.TenantService
	getAllTenantsService    *tenantsApplication.TenantService
	getTenantByAliasService *tenantsApplication.TenantService
	getTenantByIdService    *tenantsApplication.TenantService
	updateTenantService     *tenantsApplication.TenantService
	deleteTenantService     *tenantsApplication.TenantService
}

func NewTenantsHandler(
	createTenantService *tenantsApplication.TenantService,
	getAllTenantsService *tenantsApplication.TenantService,
	getTenantByAliasService *tenantsApplication.TenantService,
	getTenantByIdService *tenantsApplication.TenantService,
	updateTenantService *tenantsApplication.TenantService,
	deleteTenantService *tenantsApplication.TenantService,
) *TenantsHandler {
	return &TenantsHandler{
		createTenantService:     createTenantService,
		getAllTenantsService:    getAllTenantsService,
		getTenantByAliasService: getTenantByAliasService,
		getTenantByIdService:    getTenantByIdService,
		updateTenantService:     updateTenantService,
		deleteTenantService:     deleteTenantService,
	}
}
