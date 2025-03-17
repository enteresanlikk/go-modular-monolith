package usersDomain

import (
	commonDomain "github.com/enteresanlikk/go-modular-monolith/internal/common/domain"
	"github.com/google/uuid"
)

type Role struct {
	commonDomain.BaseEntity

	Name string `json:"name" gorm:"not null"`
}

func (Role) TableName() string {
	return "users.roles"
}

func NewRole(name string) (*Role, error) {
	r := &Role{
		BaseEntity: commonDomain.BaseEntity{
			ID: uuid.New(),
		},
		Name: name,
	}
	return r, nil
}
