package tenantsInfrastructure

import (
	tenantsDomain "github.com/enteresanlikk/go-modular-monolith/internal/tenants/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserTenantRepository struct {
	db *gorm.DB
}

func NewUserTenantRepository(db *gorm.DB) tenantsDomain.UserTenantRepository {
	db.Exec("CREATE SCHEMA IF NOT EXISTS tenants")

	db.AutoMigrate(&tenantsDomain.UserTenant{})

	return &UserTenantRepository{
		db: db,
	}
}

func (r *UserTenantRepository) Create(userTenant *tenantsDomain.UserTenant) error {
	return r.db.Create(userTenant).Error
}

func (r *UserTenantRepository) Update(userTenant *tenantsDomain.UserTenant) error {
	return r.db.Save(userTenant).Error
}

func (r *UserTenantRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&tenantsDomain.UserTenant{}, id).Error
}

func (r *UserTenantRepository) FindByID(id uuid.UUID) (*tenantsDomain.UserTenant, error) {
	var userTenant tenantsDomain.UserTenant
	result := r.db.Where("id = ?", id).First(&userTenant)
	return &userTenant, result.Error
}

func (r *UserTenantRepository) FindByUserID(userID uuid.UUID) ([]*tenantsDomain.UserTenant, error) {
	var userTenants []*tenantsDomain.UserTenant
	result := r.db.Where("user_id = ?", userID).Find(&userTenants)
	return userTenants, result.Error
}

func (r *UserTenantRepository) FindByTenantID(tenantID uuid.UUID) ([]*tenantsDomain.UserTenant, error) {
	var userTenants []*tenantsDomain.UserTenant
	result := r.db.Where("tenant_id = ?", tenantID).Find(&userTenants)
	return userTenants, result.Error
}
