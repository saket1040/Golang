package model

type Expense struct {
	ID          string
	PaidBy      string
	Amount      float64
	Splits      []Split
	Description string
}

type Split struct {
	UserID string
	Amount float64
}
