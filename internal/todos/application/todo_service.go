package todos_application

import todos_domain "github.com/enteresanlikk/go-modular-monolith/internal/todos/domain"

type TodoService struct {
	repo todos_domain.TodoRepository
}

func NewTodoService(repo todos_domain.TodoRepository) *TodoService {
	return &TodoService{repo: repo}
}
