package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Address struct {
	House   string
	Pincode string
}

type User struct {
	Id      string
	Name    string
	Address Address
}

type Item struct {
	Id       string
	Name     string
	Quantity int32
	Price    float64
	OrderNo  int32
}

type Payment struct {
	Id        string
	Amount    float64
	Type      PaymentType
	CreatedAt time.Time
}

type PaymentType int

const (
	Cash PaymentType = iota + 1
	UPI
)

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

type OrderService struct {
	orderIDtoOrder map[string]*Order
}

func (o *OrderService) CreateOrder(user *User, items []Item) error {
	if user == nil || len(items) == 0 {
		return errors.New("cant create order")
	}

	id := uuid.NewString()
	var amount float64
	for _, item := range items {
		amount += item.Price
	}
	order := &Order{
		Id:        id,
		UserID:    user.Id,
		Items:     items,
		Amount:    amount,
		Status:    Created,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	o.orderIDtoOrder[id] = order
	return nil
}

func (o *OrderService) TrackStatus(orderId string) (OrderStatus, error) {
	if val, ok := o.orderIDtoOrder[orderId]; !ok {
		return 0, errors.New("invalid order")
	} else {
		return val.Status, nil
	}
}

func (o *OrderService) CancelOrder(orderId string) error {
	if val, ok := o.orderIDtoOrder[orderId]; !ok {
		return errors.New("order doesnt exist")
	} else if val.Status >= 3 {
		return errors.New("order cant be cancelled")
	} else {
		val.Status = Cancelled
	}
	return nil
}

type PaymentService struct {
	Os                 *OrderService
	paymentIDtoPayment map[string]*Payment
}

func (p *PaymentService) CreatePayment(orderId string, amount float64, pType PaymentType) error {
	o := p.Os
	if val, ok := o.orderIDtoOrder[orderId]; !ok {
		return errors.New("order doesnt exist")
	} else if val.Status == Paid {
		return errors.New("order already paid")
	} else if val.Status == Cancelled {
		return errors.New("order is cancelled")
	} else if amount != val.Amount {
		return errors.New("payment amount mismatch")
	} else {
		payment := &Payment{
			Id:        uuid.NewString(),
			Type:      pType,
			Amount:    amount,
			CreatedAt: time.Now(),
		}
		val.Status = Paid
		val.PaymentID = payment.Id
		p.paymentIDtoPayment[payment.Id] = payment
	}
	return nil
}

func main() {
	// Initialize services
	orderService := &OrderService{orderIDtoOrder: make(map[string]*Order)}
	paymentService := &PaymentService{
		Os:                 orderService,
		paymentIDtoPayment: make(map[string]*Payment),
	}

	// Create a user
	user := &User{
		Id:   "user-123",
		Name: "John Doe",
		Address: Address{
			House:   "123 Main St",
			Pincode: "123456",
		},
	}

	// Create items
	items := []Item{
		{Id: "item-1", Name: "Laptop", Quantity: 1, Price: 1000.0},
		{Id: "item-2", Name: "Mouse", Quantity: 1, Price: 50.0},
	}

	// Create an order
	err := orderService.CreateOrder(user, items)
	if err != nil {
		fmt.Println("Error creating order:", err)
		return
	}

	// Retrieve the created order ID
	var orderId string
	for id := range orderService.orderIDtoOrder {
		orderId = id
		break
	}

	fmt.Println("Order created with ID:", orderId)

	// Track order status
	status, err := orderService.TrackStatus(orderId)
	if err != nil {
		fmt.Println("Error tracking order status:", err)
		return
	}
	fmt.Println("Order status:", status)

	// Create a payment
	err = paymentService.CreatePayment(orderId, 1050.0, Cash)
	if err != nil {
		fmt.Println("Error creating payment:", err)
		return
	}
	fmt.Println("Payment created for order ID:", orderId)

	// Track order status after payment
	status, err = orderService.TrackStatus(orderId)
	if err != nil {
		fmt.Println("Error tracking order status:", err)
		return
	}
	fmt.Println("Order status after payment:", status)

	// Cancel the order (should fail since it's already paid)
	err = orderService.CancelOrder(orderId)
	if err != nil {
		fmt.Println("Error cancelling order:", err)
	} else {
		fmt.Println("Order cancelled successfully")
	}
}
