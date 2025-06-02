package services

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"library/internals/dtos"
	"library/internals/models"
	"library/internals/repositories"
)

type LoanService struct {
	userRepo repositories.UserRepository
	bookRepo repositories.BookRepository
	loanRepo repositories.LoanRepository
}

func NewLoanService(userRepo repositories.UserRepository, bookRepo repositories.BookRepository, loanRepo repositories.LoanRepository) *LoanService {
	return &LoanService{
		userRepo: userRepo,
		bookRepo: bookRepo,
		loanRepo: loanRepo,
	}
}

func (s *LoanService) Borrow(req *dtos.BorrowRequest) error {
	user, err := s.userRepo.FindByID(req.UserID)
	if err != nil {
		return err
	}

	book, err := s.bookRepo.FindByID(req.BookID)
	if err != nil {
		return err
	}

	if book.Status != models.AVAILABLE {
		return errors.New("book is not available")
	}

	returnDate, err := time.Parse("2006-01-02", req.ReturnDate)
	if err != nil {
		return errors.New("invalid return date format, expected YYYY-MM-DD")
	}

	loan := &models.Loan{
		ID:         uuid.NewString(),
		UserID:     user.ID,
		BookID:     book.ID,
		IssuedDate: time.Now(),
		ReturnDate: returnDate,
		Status:     models.ACTIVE,
	}

	if err := s.loanRepo.Save(loan); err != nil {
		return err
	}

	return s.bookRepo.UpdateStatus(book.ID, models.BORROWED)
}

func (s *LoanService) Return(req *dtos.ReturnRequest) error {
	loan, err := s.loanRepo.FindByID(req.LoanID)
	if err != nil {
		return err
	}

	if loan.Status == models.INACTIVE {
		return errors.New("loan already inactive (returned)")
	}

	if err := s.loanRepo.UpdateStatus(loan.ID, models.INACTIVE); err != nil {
		return err
	}

	return s.bookRepo.UpdateStatus(loan.BookID, models.AVAILABLE)
}