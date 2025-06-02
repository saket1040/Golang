package repositories

import (
	"errors"
	"sync"

	"library/internals/models"
)

type BookRepository interface {
	Save(book *models.Book) error
	FindByID(bookID string) (*models.Book, error)
	UpdateStatus(bookID string, status models.BookStatus) error
}

type InMemoryBookRepository struct {
	books map[string]*models.Book
	mutex sync.RWMutex
}

func NewInMemoryBookRepository() *InMemoryBookRepository {
	return &InMemoryBookRepository{
		books: make(map[string]*models.Book),
	}
}

func (r *InMemoryBookRepository) Save(book *models.Book) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.books[book.ID]; exists {
		return errors.New("book already exists")
	}

	r.books[book.ID] = book
	return nil
}

func (r *InMemoryBookRepository) FindByID(bookID string) (*models.Book, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	book, exists := r.books[bookID]
	if !exists {
		return nil, errors.New("book not found")
	}
	return book, nil
}

func (r *InMemoryBookRepository) UpdateStatus(bookID string, status models.BookStatus) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	book, exists := r.books[bookID]
	if !exists {
		return errors.New("book not found")
	}
	book.Status = status
	return nil
}