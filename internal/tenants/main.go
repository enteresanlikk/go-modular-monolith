package tenants

import (
	tenantsApplication "github.com/enteresanlikk/go-modular-monolith/internal/tenants/application"
	tenantsInfrastructure "github.com/enteresanlikk/go-modular-monolith/internal/tenants/infrastructure"
	tenantsPresentation "github.com/enteresanlikk/go-modular-monolith/internal/tenants/presentation"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Register(r *fiber.App, db *gorm.DB) {
	tenantRepo := tenantsInfrastructure.NewTenantRepository(db)

	createTenantService := tenantsApplication.NewTenantService(tenantRepo)
	getAllTenantsService := tenantsApplication.NewTenantService(tenantRepo)
	getTenantByAliasService := tenantsApplication.NewTenantService(tenantRepo)
	getTenantByIdService := tenantsApplication.NewTenantService(tenantRepo)
	updateTenantService := tenantsApplication.NewTenantService(tenantRepo)
	deleteTenantService := tenantsApplication.NewTenantService(tenantRepo)

	handler := tenantsPresentation.NewTenantsHandler(createTenantService, getAllTenantsService, getTenantByAliasService, getTenantByIdService, updateTenantService, deleteTenantService)

	tenantsGroup := r.Group("/tenants")

	tenantsGroup.Get("/", handler.GetAllTenants)
	tenantsGroup.Get("/:id", handler.GetTenantById)
	tenantsGroup.Post("/", handler.CreateTenant)
	tenantsGroup.Put("/:id", handler.UpdateTenant)
	tenantsGroup.Delete("/:id", handler.DeleteTenant)
}
