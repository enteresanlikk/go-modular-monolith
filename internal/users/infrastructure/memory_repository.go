package infrastructure

import (
	"errors"
	"sync"

	"github.com/enteresanlikk/go-modular-monolith/internal/users/domain"
)

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrEmailAlreadyExist = errors.New("email already exists")
)

type InMemoryUserRepository struct {
	users  map[uint]*domain.User
	mu     sync.RWMutex
	lastID uint
}

func NewInMemoryUserRepository() domain.UserRepository {
	return &InMemoryUserRepository{
		users:  make(map[uint]*domain.User),
		lastID: 0,
	}
}

func (r *InMemoryUserRepository) Create(user *domain.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Check if email already exists
	for _, existingUser := range r.users {
		if existingUser.Email == user.Email {
			return ErrEmailAlreadyExist
		}
	}

	r.lastID++
	user.ID = r.lastID
	r.users[user.ID] = user
	return nil
}

func (r *InMemoryUserRepository) FindByEmail(email string) (*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, user := range r.users {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, ErrUserNotFound
}

func (r *InMemoryUserRepository) FindByID(id uint) (*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if user, exists := r.users[id]; exists {
		return user, nil
	}
	return nil, ErrUserNotFound
}
