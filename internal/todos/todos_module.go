package todos

import (
	"net/http"

	todos_application "github.com/enteresanlikk/go-modular-monolith/internal/todos/application"
	todos_infrastructure "github.com/enteresanlikk/go-modular-monolith/internal/todos/infrastructure"
	todos_presentation "github.com/enteresanlikk/go-modular-monolith/internal/todos/presentation"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func Register(mux *mux.Router, db *gorm.DB) {
	todoRepo := todos_infrastructure.NewTodoRepository(db)
	createTodoService := todos_application.NewTodoService(todoRepo)
	getAllTodosService := todos_application.NewTodoService(todoRepo)
	getTodoByIdService := todos_application.NewTodoService(todoRepo)
	updateTodoService := todos_application.NewTodoService(todoRepo)
	deleteTodoService := todos_application.NewTodoService(todoRepo)
	handler := todos_presentation.NewTodosHandler(createTodoService, getAllTodosService, getTodoByIdService, updateTodoService, deleteTodoService)

	todosRouter := mux.PathPrefix("/todos").Subrouter()

	todosRouter.HandleFunc("/", handler.GetAllTodos).Methods(http.MethodGet)
	todosRouter.HandleFunc("/{id}", handler.GetTodoById).Methods(http.MethodGet)
	todosRouter.HandleFunc("/", handler.CreateTodo).Methods(http.MethodPost)
	todosRouter.HandleFunc("/{id}", handler.UpdateTodo).Methods(http.MethodPut)
	todosRouter.HandleFunc("/{id}", handler.DeleteTodo).Methods(http.MethodDelete)
}
