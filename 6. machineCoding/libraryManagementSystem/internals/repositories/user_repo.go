package repositories

import (
	"errors"
	"sync"

	"library/internals/models"
)

type UserRepository interface {
	Save(user *models.User) error
	FindByID(userID string) (*models.User, error)
}

type InMemoryUserRepository struct {
	users map[string]*models.User
	mutex sync.RWMutex
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make(map[string]*models.User),
	}
}

func (r *InMemoryUserRepository) Save(user *models.User) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.users[user.ID]; exists {
		return errors.New("user already exists")
	}

	r.users[user.ID] = user
	return nil
}

func (r *InMemoryUserRepository) FindByID(userID string) (*models.User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	user, exists := r.users[userID]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}