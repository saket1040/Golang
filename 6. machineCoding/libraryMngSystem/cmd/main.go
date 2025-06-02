package main

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id    string
	Name  string
	Email string
}

type Book struct {
	Id     string
	Title  string
	Author string
	Status BookStatus
}

type BookStatus string

const (
	AVAILABLE BookStatus = "AVAILABLE"
	BORROWED  BookStatus = "BORROWED"
)

type Loan struct {
	Id         string
	UserId     string
	BookId     string
	IssuedDate time.Time
	ReturnDate time.Time
	Status     LoanStatus
}

type LoanStatus string

const (
	ACTIVE   LoanStatus = "ACTIVE"
	INACTIVE LoanStatus = "INACTIVE"
)

type LibraryRepo struct {
	Users map[string]*User
	Books map[string]*Book
	Loans map[string]*Loan
	Mtx   *sync.Mutex
}

type RegistrationService struct {
	repo *LibraryRepo
}

func (r *RegistrationService) RegisterUser(user *User) error {
	if strings.TrimSpace(user.Email) == "" {
		return errors.New("user email is invalid")
	}

	if _, ok := r.repo.Users[user.Id]; ok {
		return errors.New("user already registered")
	}

	r.repo.Users[user.Id] = user
	return nil
}

func (r *RegistrationService) RegisterBook(book *Book) error {
	if strings.TrimSpace(book.Title) == "" {
		return errors.New("book title is invalid")
	}

	if _, ok := r.repo.Books[book.Id]; ok {
		return errors.New("book already registered")
	}

	r.repo.Books[book.Id] = book
	return nil
}

type BorrowService struct {
	repo *LibraryRepo
}

func (b *BorrowService) Borrow(bookId, userId string, returnDate string) error {
	// 	You’re using a new sync.Mutex inside each method, which protects nothing in practice — the mutex gets discarded after the call. Instead:
	//mtx := &sync.Mutex{}
	mtx := b.repo.Mtx
	mtx.Lock()
	defer mtx.Unlock()

	book, ok := b.repo.Books[bookId]
	if !ok {
		return errors.New("book doesn't exist")
	}
	if book.Status == BORROWED {
		return errors.New("book already borrowed")
	}

	if _, ok := b.repo.Users[userId]; !ok {
		return errors.New("user doesn't exist")
	}

	returnDateTm, err := time.Parse("2006-01-02", returnDate)
	if err != nil {
		return errors.New("invalid return date format, expected YYYY-MM-DD")
	}

	//returndate check not present

	loan := &Loan{
		Id:         uuid.NewString(),
		UserId:     userId,
		BookId:     bookId,
		IssuedDate: time.Now(),
		ReturnDate: returnDateTm,
		Status:     ACTIVE,
	}

	b.repo.Loans[loan.Id] = loan
	book.Status = BORROWED
	return nil
}

type ReturnService struct {
	repo *LibraryRepo
}

func (r *ReturnService) Return(loanId string) error {
	// mtx := &sync.Mutex{}
	mtx := r.repo.Mtx
	mtx.Lock()
	defer mtx.Unlock()

	loan, ok := r.repo.Loans[loanId]
	if !ok {
		return errors.New("loan doesn't exist")
	}
	if loan.Status == INACTIVE {
		return errors.New("book already returned")
	}

	loan.Status = INACTIVE
	r.repo.Books[loan.BookId].Status = AVAILABLE
	return nil
}

type UserService struct {
	repo *LibraryRepo
}

func (u *UserService) TrackUserHistory(userId string) ([]Loan, error) {
	if _, ok := u.repo.Users[userId]; !ok {
		return nil, fmt.Errorf("user with id %s doesn't exist", userId)
	}

	loans := make([]Loan, 0)
	for _, loan := range u.repo.Loans {
		if loan.UserId == userId {
			loans = append(loans, *loan)
		}
	}
	return loans, nil
}

func main() {
	// Initialize the library repository
	repo := &LibraryRepo{
		Users: make(map[string]*User),
		Books: make(map[string]*Book),
		Loans: make(map[string]*Loan),
		Mtx:   &sync.Mutex{},
	}

	// Initialize services
	registrationService := &RegistrationService{repo: repo}
	borrowService := &BorrowService{repo: repo}
	returnService := &ReturnService{repo: repo}
	userService := &UserService{repo: repo}

	// Register a user
	user := &User{Id: "user-1", Name: "John Doe", Email: "john.doe@example.com"}
	err := registrationService.RegisterUser(user)
	if err != nil {
		fmt.Println("Error registering user:", err)
		return
	}
	fmt.Println("User registered successfully")

	// Register a book
	book := &Book{Id: "book-1", Title: "Go Programming", Author: "John Smith", Status: AVAILABLE}
	err = registrationService.RegisterBook(book)
	if err != nil {
		fmt.Println("Error registering book:", err)
		return
	}
	fmt.Println("Book registered successfully")

	// Borrow a book
	returnDate := "2025-05-20"
	err = borrowService.Borrow(book.Id, user.Id, returnDate)
	if err != nil {
		fmt.Println("Error borrowing book:", err)
		return
	}
	fmt.Println("Book borrowed successfully")

	// Track user history
	loans, err := userService.TrackUserHistory(user.Id)
	if err != nil {
		fmt.Println("Error tracking user history:", err)
		return
	}
	fmt.Println("User loan history:")
	for _, loan := range loans {
		fmt.Printf("Loan ID: %s, Book ID: %s, Issued Date: %s, Return Date: %s, Status: %s\n",
			loan.Id, loan.BookId, loan.IssuedDate.Format("2006-01-02"), loan.ReturnDate.Format("2006-01-02"), loan.Status)
	}

	// Return the book
	var loanId string
	for id, loan := range repo.Loans {
		if loan.UserId == user.Id && loan.BookId == book.Id {
			loanId = id
			break
		}
	}

	err = returnService.Return(loanId)
	if err != nil {
		fmt.Println("Error returning book:", err)
		return
	}
	fmt.Println("Book returned successfully")
}
