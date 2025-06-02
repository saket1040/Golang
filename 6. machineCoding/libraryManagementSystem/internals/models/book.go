package models

type Book struct {
	ID     string
	Title  string
	Author string
	Status BookStatus
}

type BookStatus string

const (
	AVAILABLE BookStatus = "AVAILABLE"
	BORROWED  BookStatus = "BORROWED"
)