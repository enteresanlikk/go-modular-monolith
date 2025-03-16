package todos_application

import (
	common_domain "github.com/enteresanlikk/go-modular-monolith/internal/common/domain"
	todos_domain "github.com/enteresanlikk/go-modular-monolith/internal/todos/domain"
	"github.com/google/uuid"
)

type CreateTodoRequest struct {
	Title     string `json:"title" validate:"required"`
	Completed bool   `json:"completed" validate:"required"`
}

func (s *TodoService) CreateTodo(req *CreateTodoRequest) (*TodoResponse, error) {
	todo := &todos_domain.Todo{
		Entity: common_domain.Entity{
			ID: uuid.New(),
		},
		Title:     req.Title,
		Completed: req.Completed,
	}

	if err := s.repo.Create(todo); err != nil {
		return nil, err
	}

	return &TodoResponse{
		ID:        todo.ID,
		Title:     todo.Title,
		Completed: todo.Completed,
		CreatedAt: todo.CreatedAt,
		UpdatedAt: todo.UpdatedAt,
	}, nil
}
