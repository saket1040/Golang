package services

import (
	"errors"
	"orderManagementSystem/models"
	"time"

	"github.com/google/uuid"
)

type PaymentService struct {
	Os                  *OrderService
	PaymentIDtoPayment  map[string]*models.Payment
}

func (p *PaymentService) CreatePayment(orderId string, amount float64, pType models.PaymentType) error {
	order, ok := p.Os.OrderIDtoOrder[orderId]
	if !ok {
		return errors.New("order doesn't exist")
	}
	if order.Status == models.Paid {
		return errors.New("order already paid")
	}
	if order.Status == models.Cancelled {
		return errors.New("order is cancelled")
	}
	if amount != order.Amount {
		return errors.New("payment amount mismatch")
	}

	payment := &models.Payment{
		Id:        uuid.NewString(),
		Type:      pType,
		Amount:    amount,
		CreatedAt: time.Now(),
	}
	order.Status = models.Paid
	order.PaymentID = payment.Id
	order.UpdatedAt = time.Now()

	p.PaymentIDtoPayment[payment.Id] = payment
	return nil
}