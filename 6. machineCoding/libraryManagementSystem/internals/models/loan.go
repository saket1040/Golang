package models

import "time"

type Loan struct {
	ID         string
	UserID     string
	BookID    string
	IssuedDate time.Time
	ReturnDate time.Time
	Status     LoanStatus
}

type LoanStatus string

const (
	ACTIVE   LoanStatus = "ACTIVE"
	INACTIVE LoanStatus = "INACTIVE"
)