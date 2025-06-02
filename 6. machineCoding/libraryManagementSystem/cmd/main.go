package main

import (
	"fmt"
	"time"

	"library/internals/dtos"
	"library/internals/models"
	"library/internals/repositories"
	"library/internals/services"
)

func main() {
	// Setup Repositories
	userRepo := repositories.NewInMemoryUserRepository()
	bookRepo := repositories.NewInMemoryBookRepository()
	loanRepo := repositories.NewInMemoryLoanRepository()

	// Setup Services
	userService := services.NewUserService(userRepo, loanRepo)
	loanService := services.NewLoanService(userRepo, bookRepo, loanRepo)

	// Register a User
	user := &models.User{ID: "user-1", Name: "Alice", Email: "alice@example.com"}
	err := userRepo.Save(user)
	if err != nil {
		fmt.Println("Error registering user:", err)
		return
	}

	// Register a Book
	book := &models.Book{ID: "book-1", Title: "GoLang 101", Author: "Bob", Status: models.AVAILABLE}
	err = bookRepo.Save(book)
	if err != nil {
		fmt.Println("Error registering book:", err)
		return
	}

	// Borrow a Book
	err = loanService.Borrow(&dtos.BorrowRequest{
		BookID:     book.ID,
		UserID:     user.ID,
		ReturnDate: time.Now().AddDate(0, 0, 7).Format("2006-01-02"),
	})
	if err != nil {
		fmt.Println("Error borrowing book:", err)
		return
	}
	fmt.Println("Book borrowed successfully")

	// Track User Loan History
	loans, err := userService.GetLoanHistory(user.ID)
	if err != nil {
		fmt.Println("Error tracking history:", err)
		return
	}
	for _, loan := range loans {
		fmt.Printf("Loan ID: %s, Book ID: %s, Status: %s\n", loan.ID, loan.BookID, loan.Status)
	}

	// Return Book
	if len(loans) > 0 {
		err = loanService.Return(& dtos.ReturnRequest{LoanID: loans[0].ID})
		if err != nil {
			fmt.Println("Error returning book:", err)
			return
		}
		fmt.Println("Book returned successfully")
	}
}
