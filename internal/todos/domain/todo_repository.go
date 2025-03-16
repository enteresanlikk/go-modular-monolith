package todos_domain

import "github.com/google/uuid"

type TodoRepository interface {
	Create(todo *Todo) error
	FindByID(id uuid.UUID) (*Todo, error)
	FindAll() ([]*Todo, error)
	Update(todo *Todo) error
	Delete(id uuid.UUID) error
}
