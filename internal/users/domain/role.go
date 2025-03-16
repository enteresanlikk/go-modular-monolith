package users_domain

import (
	common_domain "github.com/enteresanlikk/go-modular-monolith/internal/common/domain"
	"github.com/google/uuid"
)

type Role struct {
	common_domain.Entity

	Name string `json:"name" gorm:"not null"`
}

func (Role) TableName() string {
	return "users.roles"
}

func NewRole(name string) (*Role, error) {
	r := &Role{
		Entity: common_domain.Entity{
			ID: uuid.New(),
		},
		Name: name,
	}
	return r, nil
}
