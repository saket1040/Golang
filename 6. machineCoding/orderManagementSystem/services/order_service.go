package services

import (
	"errors"
	"orderManagementSystem/models"
	"time"

	"github.com/google/uuid"
)

type OrderService struct {
	OrderIDtoOrder map[string]*models.Order
}

func (o *OrderService) CreateOrder(user *models.User, items []models.Item) error {
	if user == nil || len(items) == 0 {
		return errors.New("cannot create order")
	}

	id := uuid.NewString()
	var amount float64
	for _, item := range items {
		amount += item.Price
	}
	order := &models.Order{
		Id:        id,
		UserID:    user.Id,
		Items:     items,
		Amount:    amount,
		Status:    models.Created,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	o.OrderIDtoOrder[id] = order
	return nil
}

func (o *OrderService) TrackStatus(orderId string) (models.OrderStatus, error) {
	order, ok := o.OrderIDtoOrder[orderId]
	if !ok {
		return 0, errors.New("invalid order")
	}
	return order.Status, nil
}

func (o *OrderService) CancelOrder(orderId string) error {
	order, ok := o.OrderIDtoOrder[orderId]
	if !ok {
		return errors.New("order doesn't exist")
	}
	if order.Status >= models.Shipped {
		return errors.New("order can't be cancelled")
	}
	order.Status = models.Cancelled
	order.UpdatedAt = time.Now()
	return nil
}