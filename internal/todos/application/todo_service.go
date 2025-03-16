package todosApplication

import todosDomain "github.com/enteresanlikk/go-modular-monolith/internal/todos/domain"

type TodoService struct {
	repo todosDomain.TodoRepository
}

func NewTodoService(repo todosDomain.TodoRepository) *TodoService {
	return &TodoService{repo: repo}
}
