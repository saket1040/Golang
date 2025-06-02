package repositories

import (
	"errors"
	"sync"

	"library/internals/models"
)

type LoanRepository interface {
	Save(loan *models.Loan) error
	FindByID(loanID string) (*models.Loan, error)
	FindByUserID(userID string) ([]*models.Loan, error)
	UpdateStatus(loanID string, status models.LoanStatus) error
}

type InMemoryLoanRepository struct {
	loans map[string]*models.Loan
	mutex sync.RWMutex
}

func NewInMemoryLoanRepository() *InMemoryLoanRepository {
	return &InMemoryLoanRepository{
		loans: make(map[string]*models.Loan),
	}
}

func (r *InMemoryLoanRepository) Save(loan *models.Loan) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.loans[loan.ID]; exists {
		return errors.New("loan already exists")
	}
	r.loans[loan.ID] = loan
	return nil
}

func (r *InMemoryLoanRepository) FindByID(loanID string) (*models.Loan, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	loan, exists := r.loans[loanID]
	if !exists {
		return nil, errors.New("loan not found")
	}
	return loan, nil
}

func (r *InMemoryLoanRepository) FindByUserID(userID string) ([]*models.Loan, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	var results []*models.Loan
	for _, loan := range r.loans {
		if loan.UserID == userID {
			results = append(results, loan)
		}
	}
	return results, nil
}

func (r *InMemoryLoanRepository) UpdateStatus(loanID string, status models.LoanStatus) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	loan, exists := r.loans[loanID]
	if !exists {
		return errors.New("loan not found")
	}
	loan.Status = status
	return nil
}