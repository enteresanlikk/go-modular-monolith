package todosApplication

import (
	"time"

	todosDomain "github.com/enteresanlikk/go-modular-monolith/internal/todos/domain"
	"github.com/google/uuid"
)

type TodoResponse struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (t *TodoResponse) FromTodo(todo *todosDomain.Todo) *TodoResponse {
	t.ID = todo.ID
	t.Title = todo.Title
	t.Completed = todo.Completed
	t.CreatedAt = todo.CreatedAt
	t.UpdatedAt = todo.UpdatedAt
	return t
}
