package usersDomain

import (
	commonDomain "github.com/enteresanlikk/go-modular-monolith/internal/common/domain"
	"github.com/google/uuid"
)

type Role struct {
	commonDomain.Entity

	Name string `json:"name" gorm:"not null"`
}

func (Role) TableName() string {
	return "users.roles"
}

func NewRole(name string) (*Role, error) {
	r := &Role{
		Entity: commonDomain.Entity{
			ID: uuid.New(),
		},
		Name: name,
	}
	return r, nil
}
