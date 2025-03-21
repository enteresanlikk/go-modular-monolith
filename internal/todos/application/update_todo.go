package todosApplication

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

	todo, err = todo.Update(req.Title, req.Completed)
	if err != nil {
		return nil, err
	}

	if err := s.repo.Update(todo); err != nil {
		return nil, err
	}

	return (&TodoResponse{}).FromTodo(todo), nil
}
