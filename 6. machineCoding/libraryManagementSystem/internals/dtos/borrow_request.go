package dtos

type BorrowRequest struct {
	BookID     string `json:"book_id"`
	UserID     string `json:"user_id"`
	ReturnDate string `json:"return_date"`
}
