package usersDomain

import (
	commonDomain "github.com/enteresanlikk/go-modular-monolith/internal/common/domain"
	"github.com/google/uuid"
)

type Permission struct {
	commonDomain.BaseEntity

	Name string `json:"name" gorm:"not null"`
}

func (Permission) TableName() string {
	return "users.permissions"
}

func NewPermission(name string) (*Permission, error) {
	p := &Permission{
		BaseEntity: commonDomain.BaseEntity{
			ID: uuid.New(),
		},
		Name: name,
	}
	return p, nil
}
