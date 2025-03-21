package tenants

import (
	tenantsApplication "github.com/enteresanlikk/go-modular-monolith/internal/tenants/application"
	tenantsInfrastructure "github.com/enteresanlikk/go-modular-monolith/internal/tenants/infrastructure"
	tenantsPresentation "github.com/enteresanlikk/go-modular-monolith/internal/tenants/presentation"
	"github.com/fasthttp/router"
	"gorm.io/gorm"
)

func Register(r *router.Router, db *gorm.DB) {
	tenantRepo := tenantsInfrastructure.NewTenantRepository(db)

	createTenantService := tenantsApplication.NewTenantService(tenantRepo)
	getAllTenantsService := tenantsApplication.NewTenantService(tenantRepo)
	getTenantByAliasService := tenantsApplication.NewTenantService(tenantRepo)
	getTenantByIdService := tenantsApplication.NewTenantService(tenantRepo)
	updateTenantService := tenantsApplication.NewTenantService(tenantRepo)
	deleteTenantService := tenantsApplication.NewTenantService(tenantRepo)

	handler := tenantsPresentation.NewTenantsHandler(createTenantService, getAllTenantsService, getTenantByAliasService, getTenantByIdService, updateTenantService, deleteTenantService)

	tenantsGroup := r.Group("/tenants")

	tenantsGroup.GET("/", handler.GetAllTenants)
	tenantsGroup.GET("/{id}", handler.GetTenantById)
	tenantsGroup.POST("/", handler.CreateTenant)
	tenantsGroup.PUT("/{id}", handler.UpdateTenant)
	tenantsGroup.DELETE("/{id}", handler.DeleteTenant)
}
