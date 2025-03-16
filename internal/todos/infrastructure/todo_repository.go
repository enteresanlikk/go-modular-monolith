package todosInfrastructure

import (
	todosDomain "github.com/enteresanlikk/go-modular-monolith/internal/todos/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TodoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) todosDomain.TodoRepository {
	db.Exec("CREATE SCHEMA IF NOT EXISTS todos")

	db.AutoMigrate(&todosDomain.Todo{})

	return &TodoRepository{
		db: db,
	}
}

func (r *TodoRepository) Create(todo *todosDomain.Todo) error {
	return r.db.Create(todo).Error
}

func (r *TodoRepository) FindByID(id uuid.UUID) (*todosDomain.Todo, error) {
	var todo todosDomain.Todo
	result := r.db.Where("id = ?", id).First(&todo)
	return &todo, result.Error
}

func (r *TodoRepository) FindAll() ([]*todosDomain.Todo, error) {
	var todos []*todosDomain.Todo
	result := r.db.Find(&todos)
	return todos, result.Error
}

func (r *TodoRepository) Update(todo *todosDomain.Todo) error {
	return r.db.Save(todo).Error
}

func (r *TodoRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&todosDomain.Todo{}, id).Error
}
