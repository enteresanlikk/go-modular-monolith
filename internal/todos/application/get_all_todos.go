package todos_application

type GetAllTodosRequest struct{}

func (s *TodoService) GetAllTodos(req *GetAllTodosRequest) ([]*TodoResponse, error) {
	todos, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	todosResponse := make([]*TodoResponse, len(todos))
	for i, todo := range todos {
		todosResponse[i] = &TodoResponse{
			ID:        todo.ID,
			Title:     todo.Title,
			Completed: todo.Completed,
			CreatedAt: todo.CreatedAt,
			UpdatedAt: todo.UpdatedAt,
		}
	}

	return todosResponse, nil
}
