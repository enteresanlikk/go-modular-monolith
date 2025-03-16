package users_domain

import (
	"os"
	"strconv"

	common_domain "github.com/enteresanlikk/go-modular-monolith/internal/common/domain"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
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

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()

	var bcryptCost, _ = strconv.Atoi(os.Getenv("BCRYPT_COST"))
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcryptCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return
}
