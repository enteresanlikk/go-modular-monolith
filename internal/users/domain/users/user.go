package users

import (
	"github.com/enteresanlikk/go-modular-monolith/internal/common"
)

type User struct {
	common.Entity

	Email    string `json:"email" gorm:"uniqueIndex;not null"`
	Password string `json:"-" gorm:"not null"`
}
