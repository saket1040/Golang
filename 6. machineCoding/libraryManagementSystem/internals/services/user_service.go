package services

import (
	"fmt"

	"library/internals/models"
	"library/internals/repositories"
)

type UserService struct {
	userRepo repositories.UserRepository
	loanRepo repositories.LoanRepository
}

func NewUserService(userRepo repositories.UserRepository, loanRepo repositories.LoanRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
		loanRepo: loanRepo,
	}
}

func (s *UserService) GetLoanHistory(userID string) ([]*models.Loan, error) {
	if _, err := s.userRepo.FindByID(userID); err != nil {
		return nil, fmt.Errorf("user %s not found", userID)
	}

	loans, err := s.loanRepo.FindByUserID(userID)
	if err != nil {
		return nil, err
	}

	return loans, nil
}