package todosApplication

import (
	todosDomain "github.com/enteresanlikk/go-modular-monolith/internal/todos/domain"
)

type CreateTodoRequest struct {
	Title     string `json:"title" validate:"required"`
	Completed bool   `json:"completed" validate:"required"`
}

func (s *TodoService) CreateTodo(req *CreateTodoRequest) (*TodoResponse, error) {
	todo, err := todosDomain.NewTodo(req.Title, req.Completed)
	if err != nil {
		return nil, err
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
