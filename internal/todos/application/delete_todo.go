package todosApplication

import "github.com/google/uuid"

type DeleteTodoRequest struct {
	ID uuid.UUID `json:"id" validate:"required"`
}

func (s *TodoService) DeleteTodo(req *DeleteTodoRequest) error {
	todo, err := s.repo.FindByID(req.ID)
	if err != nil {
		return err
	}

	return s.repo.Delete(todo.ID)
}
