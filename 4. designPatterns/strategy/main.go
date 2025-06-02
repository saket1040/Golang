package main

import "fmt"

// 	Use Factory when: you’re deciding which object to create.
// 	Use Strategy when: you’re deciding how an object should behave.

type PaymentStrategy interface {
	ProcessPayment(amount float64)
}

type CreditCardPayment struct{}

func (c *CreditCardPayment) ProcessPayment(amount float64) {
	fmt.Println("Sending CC Payment to", amount)
}

type UPIPayment struct{}

func (u *UPIPayment) ProcessPayment(amount float64) {
	fmt.Println("Sending UPI Payment to", amount)
}

type NetBankingPayment struct{}

func (n *NetBankingPayment) ProcessPayment(amount float64) {
	fmt.Println("Sending NB Payment to", amount)
}

type PaymentProcessor struct {
	PaymentStrategy
}

func (p *PaymentProcessor) MakePayment(amount float64) {
	p.PaymentStrategy.ProcessPayment(amount)
}

func main() {
	processor := &PaymentProcessor{
		PaymentStrategy: &UPIPayment{},
	}
	processor.MakePayment(100.50)

	processor.PaymentStrategy = &CreditCardPayment{}
	processor.MakePayment(250.75)
}
