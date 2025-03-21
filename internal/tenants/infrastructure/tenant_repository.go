package tenantsInfrastructure

import (
	tenantsDomain "github.com/enteresanlikk/go-modular-monolith/internal/tenants/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TenantRepository struct {
	db *gorm.DB
}

func NewTenantRepository(db *gorm.DB) tenantsDomain.TenantRepository {
	db.Exec("CREATE SCHEMA IF NOT EXISTS tenants")

	db.AutoMigrate(&tenantsDomain.Tenant{})

	return &TenantRepository{
		db: db,
	}
}

func (r *TenantRepository) Create(tenant *tenantsDomain.Tenant) error {
	return r.db.Create(tenant).Error
}

func (r *TenantRepository) Update(tenant *tenantsDomain.Tenant) error {
	return r.db.Save(tenant).Error
}

func (r *TenantRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&tenantsDomain.Tenant{}, id).Error
}

func (r *TenantRepository) FindByID(id uuid.UUID) (*tenantsDomain.Tenant, error) {
	var tenant tenantsDomain.Tenant
	result := r.db.Where("id = ?", id).First(&tenant)
	return &tenant, result.Error
}

func (r *TenantRepository) FindByAlias(alias string) (*tenantsDomain.Tenant, error) {
	var tenant tenantsDomain.Tenant
	result := r.db.Where("alias = ?", alias).First(&tenant)
	return &tenant, result.Error
}

func (r *TenantRepository) FindAll() ([]*tenantsDomain.Tenant, error) {
	var tenants []*tenantsDomain.Tenant
	result := r.db.Find(&tenants)
	return tenants, result.Error
}
