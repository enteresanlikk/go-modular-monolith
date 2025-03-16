package users_domain

import (
	common_domain "github.com/enteresanlikk/go-modular-monolith/internal/common/domain"
	"github.com/google/uuid"
)

type Permission struct {
	common_domain.Entity

	Name string `json:"name" gorm:"not null"`
}

func (Permission) TableName() string {
	return "users.permissions"
}

func NewPermission(name string) (*Permission, error) {
	p := &Permission{
		Entity: common_domain.Entity{
			ID: uuid.New(),
		},
		Name: name,
	}
	return p, nil
}
