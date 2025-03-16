package todosDomain

import (
	commonDomain "github.com/enteresanlikk/go-modular-monolith/internal/common/domain"
	"github.com/google/uuid"
)

type Todo struct {
	commonDomain.Entity

	Title     string `json:"title" gorm:"not null"`
	Completed bool   `json:"completed" gorm:"not null"`
}

func (Todo) TableName() string {
	return "todos.todos"
}

func NewTodo(title string, completed bool) (*Todo, error) {
	t := &Todo{
		Entity: commonDomain.Entity{
			ID: uuid.New(),
		},
		Title:     title,
		Completed: completed,
	}
	return t, nil
}

func (t *Todo) Update(title string, completed bool) (*Todo, error) {
	t.Title = title
	t.Completed = completed
	return t, nil
}
