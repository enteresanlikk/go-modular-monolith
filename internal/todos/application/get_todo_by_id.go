package todos_application

import "github.com/google/uuid"

type GetTodoByIdRequest struct {
	ID uuid.UUID `json:"id" validate:"required"`
}

func (s *TodoService) GetTodoById(req *GetTodoByIdRequest) (*TodoResponse, error) {
	todo, err := s.repo.FindByID(req.ID)
	if err != nil {
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
