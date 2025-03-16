package todos_infrastructure

import (
	todos_domain "github.com/enteresanlikk/go-modular-monolith/internal/todos/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TodoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) todos_domain.TodoRepository {
	db.Exec("CREATE SCHEMA IF NOT EXISTS todos")

	db.AutoMigrate(&todos_domain.Todo{})

	return &TodoRepository{
		db: db,
	}
}

func (r *TodoRepository) Create(todo *todos_domain.Todo) error {
	return r.db.Create(todo).Error
}

func (r *TodoRepository) FindByID(id uuid.UUID) (*todos_domain.Todo, error) {
	var todo todos_domain.Todo
	result := r.db.Where("id = ?", id).First(&todo)
	return &todo, result.Error
}

func (r *TodoRepository) FindAll() ([]*todos_domain.Todo, error) {
	var todos []*todos_domain.Todo
	result := r.db.Find(&todos)
	return todos, result.Error
}

func (r *TodoRepository) Update(todo *todos_domain.Todo) error {
	return r.db.Save(todo).Error
}

func (r *TodoRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&todos_domain.Todo{}, id).Error
}
