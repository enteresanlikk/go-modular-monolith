package todos

import (
	todosApplication "github.com/enteresanlikk/go-modular-monolith/internal/todos/application"
	todosInfrastructure "github.com/enteresanlikk/go-modular-monolith/internal/todos/infrastructure"
	todosPresentation "github.com/enteresanlikk/go-modular-monolith/internal/todos/presentation"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Register(r *fiber.App, db *gorm.DB) {
	todoRepo := todosInfrastructure.NewTodoRepository(db)
	createTodoService := todosApplication.NewTodoService(todoRepo)
	getAllTodosService := todosApplication.NewTodoService(todoRepo)
	getTodoByIdService := todosApplication.NewTodoService(todoRepo)
	updateTodoService := todosApplication.NewTodoService(todoRepo)
	deleteTodoService := todosApplication.NewTodoService(todoRepo)
	handler := todosPresentation.NewTodosHandler(createTodoService, getAllTodosService, getTodoByIdService, updateTodoService, deleteTodoService)

	todosGroup := r.Group("/todos")

	todosGroup.Get("/", handler.GetAllTodos)
	todosGroup.Get("/:id", handler.GetTodoById)
	todosGroup.Post("/", handler.CreateTodo)
	todosGroup.Put("/:id", handler.UpdateTodo)
	todosGroup.Delete("/:id", handler.DeleteTodo)
}
