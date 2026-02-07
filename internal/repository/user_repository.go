package repository

import (
	"sync"

	"user_service/internal/model"
	"user_service/pkg/errors"

	"github.com/google/uuid"
)

type UserRepository interface {
	Create(user *model.User) (*model.User, error)
	GetByID(id string) (model.User, error)
	GetAll() ([]*model.User, error)
	Update(user *model.User) (*model.User, error)
	Delete(id string) error
}

type InMemoryUserRepository struct {
	mu      sync.RWMutex
	users   map[string]model.User
	byEmail map[string]string
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users:   make(map[string]model.User),
		byEmail: make(map[string]string),
	}
}

// Create a new user in the repository
func (r *InMemoryUserRepository) Create(user *model.User) (*model.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, u := range r.users {
		if u.Email == user.Email {
			return nil, errors.ErrAlreadyExists
		}
	}

	if user.ID == "" {
		user.ID = uuid.New().String()
	}

	r.users[user.ID] = *user
	return user, nil
}

// GetByID retrieves a user by their ID
func (r *InMemoryUserRepository) GetByID(id string) (model.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, ok := r.users[id]
	if !ok {
		return model.User{}, errors.ErrNotFound
	}

	return user, nil
}

// GetAll retrieves all users from the repository
func (r *InMemoryUserRepository) GetAll() ([]*model.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	users := make([]*model.User, 0, len(r.users))
	for _, u := range r.users {
		// return pointer to each user
		userCopy := u
		users = append(users, &userCopy)
	}

	return users, nil
}

// Update a user by ID
func (r *InMemoryUserRepository) Update(user *model.User) (*model.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	existing, ok := r.users[user.ID]
	if !ok {
		return nil, errors.ErrNotFound
	}

	// Updating the fields here
	existing.Name = user.Name
	existing.Email = user.Email

	r.users[user.ID] = existing
	return &existing, nil
}

// Delete User by ID
func (r *InMemoryUserRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.users[id]; !ok {
		return errors.ErrNotFound
	}

	delete(r.users, id)
	return nil
}
