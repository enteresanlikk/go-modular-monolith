package users_domain

import (
	common_domain "github.com/enteresanlikk/go-modular-monolith/internal/common/domain"
	"github.com/google/uuid"
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

func (u *User) Create(firstName, lastName, username, email, password string) (*User, error) {
	u.ID = uuid.New()
	u.FirstName = firstName
	u.LastName = lastName
	u.Username = username
	u.Email = email
	u.Password = password
	return u, nil
}
