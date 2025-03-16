package todosPresentation

import (
	todosApplication "github.com/enteresanlikk/go-modular-monolith/internal/todos/application"
)

type TodosHandler struct {
	createTodoService  *todosApplication.TodoService
	getAllTodosService *todosApplication.TodoService
	getTodoByIdService *todosApplication.TodoService
	updateTodoService  *todosApplication.TodoService
	deleteTodoService  *todosApplication.TodoService
}

func NewTodosHandler(createTodoService *todosApplication.TodoService, getAllTodosService *todosApplication.TodoService, getTodoByIdService *todosApplication.TodoService, updateTodoService *todosApplication.TodoService, deleteTodoService *todosApplication.TodoService) *TodosHandler {
	return &TodosHandler{
		createTodoService:  createTodoService,
		getAllTodosService: getAllTodosService,
		getTodoByIdService: getTodoByIdService,
		updateTodoService:  updateTodoService,
		deleteTodoService:  deleteTodoService,
	}
}
