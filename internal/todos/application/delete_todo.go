package todos_application

import "github.com/google/uuid"

type DeleteTodoRequest struct {
	ID uuid.UUID `json:"id" validate:"required"`
}

func (s *TodoService) DeleteTodo(req *DeleteTodoRequest) error {
	return s.repo.Delete(req.ID)
}
