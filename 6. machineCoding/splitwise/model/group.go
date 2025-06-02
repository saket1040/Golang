package model

type Group struct {
	ID       string
	Name     string
	Members  []string // User IDs
	Expenses []string // Expense IDs
}
