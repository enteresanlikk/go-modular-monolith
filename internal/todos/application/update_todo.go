package todos_application

import "github.com/google/uuid"

type UpdateTodoRequest struct {
	ID        uuid.UUID `json:"id" validate:"required"`
	Title     string    `json:"title" validate:"required"`
	Completed bool      `json:"completed" validate:"required"`
}

func (s *TodoService) UpdateTodo(req *UpdateTodoRequest) (*TodoResponse, error) {
	todo, err := s.repo.FindByID(req.ID)
	if err != nil {
		return nil, err
	}

	todo.Title = req.Title
	todo.Completed = req.Completed

	if err := s.repo.Update(todo); err != nil {
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
