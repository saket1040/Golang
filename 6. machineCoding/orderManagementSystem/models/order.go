package models

import "time"

type OrderStatus int

const (
	Created OrderStatus = iota + 1
	Paid
	Shipped
	Delivered
	Cancelled
)

type Order struct {
	Id        string
	OrderNo   int32
	Items     []Item
	Amount    float64
	Status    OrderStatus
	UserID    string
	PaymentID string
	CreatedAt time.Time
	UpdatedAt time.Time
}