package todos

import (
	"net/http"

	todosApplication "github.com/enteresanlikk/go-modular-monolith/internal/todos/application"
	todosInfrastructure "github.com/enteresanlikk/go-modular-monolith/internal/todos/infrastructure"
	todosPresentation "github.com/enteresanlikk/go-modular-monolith/internal/todos/presentation"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func Register(mux *mux.Router, db *gorm.DB) {
	todoRepo := todosInfrastructure.NewTodoRepository(db)
	createTodoService := todosApplication.NewTodoService(todoRepo)
	getAllTodosService := todosApplication.NewTodoService(todoRepo)
	getTodoByIdService := todosApplication.NewTodoService(todoRepo)
	updateTodoService := todosApplication.NewTodoService(todoRepo)
	deleteTodoService := todosApplication.NewTodoService(todoRepo)
	handler := todosPresentation.NewTodosHandler(createTodoService, getAllTodosService, getTodoByIdService, updateTodoService, deleteTodoService)

	todosRouter := mux.PathPrefix("/todos").Subrouter()

	todosRouter.HandleFunc("", handler.GetAllTodos).Methods(http.MethodGet)
	todosRouter.HandleFunc("/{id}", handler.GetTodoById).Methods(http.MethodGet)
	todosRouter.HandleFunc("", handler.CreateTodo).Methods(http.MethodPost)
	todosRouter.HandleFunc("/{id}", handler.UpdateTodo).Methods(http.MethodPut)
	todosRouter.HandleFunc("/{id}", handler.DeleteTodo).Methods(http.MethodDelete)
}
