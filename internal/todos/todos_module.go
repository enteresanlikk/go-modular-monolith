package todos

import (
	todosApplication "github.com/enteresanlikk/go-modular-monolith/internal/todos/application"
	todosInfrastructure "github.com/enteresanlikk/go-modular-monolith/internal/todos/infrastructure"
	todosPresentation "github.com/enteresanlikk/go-modular-monolith/internal/todos/presentation"
	"github.com/fasthttp/router"
	"gorm.io/gorm"
)

func Register(r *router.Router, db *gorm.DB) {
	todoRepo := todosInfrastructure.NewTodoRepository(db)
	createTodoService := todosApplication.NewTodoService(todoRepo)
	getAllTodosService := todosApplication.NewTodoService(todoRepo)
	getTodoByIdService := todosApplication.NewTodoService(todoRepo)
	updateTodoService := todosApplication.NewTodoService(todoRepo)
	deleteTodoService := todosApplication.NewTodoService(todoRepo)
	handler := todosPresentation.NewTodosHandler(createTodoService, getAllTodosService, getTodoByIdService, updateTodoService, deleteTodoService)

	todosGroup := r.Group("/todos")

	todosGroup.GET("/", handler.GetAllTodos)
	todosGroup.GET("/{id}", handler.GetTodoById)
	todosGroup.POST("/", handler.CreateTodo)
	todosGroup.PUT("/{id}", handler.UpdateTodo)
	todosGroup.DELETE("/{id}", handler.DeleteTodo)
}
