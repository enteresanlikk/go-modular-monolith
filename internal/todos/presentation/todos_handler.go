package todos_presentation

import (
	todos_application "github.com/enteresanlikk/go-modular-monolith/internal/todos/application"
)

type TodosHandler struct {
	createTodoService  *todos_application.TodoService
	getAllTodosService *todos_application.TodoService
	getTodoByIdService *todos_application.TodoService
	updateTodoService  *todos_application.TodoService
	deleteTodoService  *todos_application.TodoService
}

func NewTodosHandler(createTodoService *todos_application.TodoService, getAllTodosService *todos_application.TodoService, getTodoByIdService *todos_application.TodoService, updateTodoService *todos_application.TodoService, deleteTodoService *todos_application.TodoService) *TodosHandler {
	return &TodosHandler{
		createTodoService:  createTodoService,
		getAllTodosService: getAllTodosService,
		getTodoByIdService: getTodoByIdService,
		updateTodoService:  updateTodoService,
		deleteTodoService:  deleteTodoService,
	}
}
