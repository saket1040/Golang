package models

import "time"

type PaymentType int

const (
	Cash PaymentType = iota + 1
	UPI
)

type Payment struct {
	Id        string
	Amount    float64
	Type      PaymentType
	CreatedAt time.Time
}