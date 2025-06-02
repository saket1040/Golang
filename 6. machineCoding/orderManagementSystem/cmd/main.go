package main

import (
	"fmt"
	"orderManagementSystem/models"
	"orderManagementSystem/services"
)

func main() {
	orderService := &services.OrderService{OrderIDtoOrder: make(map[string]*models.Order)}
	paymentService := &services.PaymentService{
		Os:                 orderService,
		PaymentIDtoPayment: make(map[string]*models.Payment),
	}

	user := &models.User{
		Id:   "user-123",
		Name: "John Doe",
		Address: models.Address{
			House:   "123 Main St",
			Pincode: "123456",
		},
	}

	items := []models.Item{
		{Id: "item-1", Name: "Laptop", Quantity: 1, Price: 1000.0},
		{Id: "item-2", Name: "Mouse", Quantity: 1, Price: 50.0},
	}

	err := orderService.CreateOrder(user, items)
	if err != nil {
		fmt.Println("Error creating order:", err)
		return
	}

	var orderId string
	for id := range orderService.OrderIDtoOrder {
		orderId = id
		break
	}

	fmt.Println("Order created with ID:", orderId)

	status, err := orderService.TrackStatus(orderId)
	if err != nil {
		fmt.Println("Error tracking order status:", err)
		return
	}
	fmt.Println("Order status:", status)

	err = paymentService.CreatePayment(orderId, 1050.0, models.Cash)
	if err != nil {
		fmt.Println("Error creating payment:", err)
		return
	}
	fmt.Println("Payment created for order ID:", orderId)

	status, err = orderService.TrackStatus(orderId)
	if err != nil {
		fmt.Println("Error tracking order status:", err)
		return
	}
	fmt.Println("Order status after payment:", status)

	err = orderService.CancelOrder(orderId)
	if err != nil {
		fmt.Println("Error cancelling order:", err)
	} else {
		fmt.Println("Order cancelled successfully")
	}
}