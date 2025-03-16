package todos_domain

import (
	common_domain "github.com/enteresanlikk/go-modular-monolith/internal/common/domain"
)

type Todo struct {
	common_domain.Entity

	Title     string `json:"title" gorm:"not null"`
	Completed bool   `json:"completed" gorm:"not null"`
}

func (Todo) TableName() string {
	return "todos.todos"
}
