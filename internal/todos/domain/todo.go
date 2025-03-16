package todos_domain

import (
	common_domain "github.com/enteresanlikk/go-modular-monolith/internal/common/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Todo struct {
	common_domain.Entity

	Title     string `json:"title" gorm:"not null"`
	Completed bool   `json:"completed" gorm:"not null"`
}

func (Todo) TableName() string {
	return "todos.todos"
}

func (t *Todo) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = uuid.New()
	return
}
