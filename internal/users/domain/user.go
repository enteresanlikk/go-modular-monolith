package users_domain

import (
	common_domain "github.com/enteresanlikk/go-modular-monolith/internal/common/domain"
)

type User struct {
	common_domain.Entity

	FirstName string `json:"firstName" gorm:"not null"`
	LastName  string `json:"lastName" gorm:"not null"`
	Username  string `json:"username" gorm:"uniqueIndex;not null"`
	Email     string `json:"email" gorm:"uniqueIndex;not null"`
	Password  string `json:"-" gorm:"not null"`
}

func (u *User) TableName() string {
	return "users.users"
}
