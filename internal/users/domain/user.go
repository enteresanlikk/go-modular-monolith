package usersDomain

import (
	commonDomain "github.com/enteresanlikk/go-modular-monolith/internal/common/domain"
	"github.com/google/uuid"
)

type User struct {
	commonDomain.BaseEntity

	FirstName string `json:"firstName" gorm:"not null"`
	LastName  string `json:"lastName" gorm:"not null"`
	Username  string `json:"username" gorm:"uniqueIndex;not null"`
	Email     string `json:"email" gorm:"uniqueIndex;not null"`
	Password  string `json:"-" gorm:"not null"`
}

func (u *User) TableName() string {
	return "users.users"
}

func NewUser(firstName, lastName, username, email, password string) (*User, error) {
	u := &User{
		BaseEntity: commonDomain.BaseEntity{
			ID: uuid.New(),
		},
		FirstName: firstName,
		LastName:  lastName,
		Username:  username,
		Email:     email,
		Password:  password,
	}
	return u, nil
}
