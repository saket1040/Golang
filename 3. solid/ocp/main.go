package main

import "fmt"

// Defining a variable inside an interface is not allowed in Go.
// Interfaces can only contain method signatures.
type PaymentProcessor interface {
	Pay(amount float64)
}

type CreditCardProcessor struct {}

func (c *CreditCardProcessor) Pay(amount float64) {
	fmt.Println("Paying through Credit card, amount: ", amount)
}

type UPIProcessor struct {}

func (c *UPIProcessor) Pay(amount float64) {
	fmt.Println("Paying through UPI, amount: ", amount)
}

//create payment through PayPal
type PayPalProcessor struct {}

func (c *PayPalProcessor) Pay(amount float64) {
	fmt.Println("Paying through PayPal, amount: ", amount)
}

func MakePayment(processor PaymentProcessor, amount float64) {
	processor.Pay(amount)
}

func main() {
	fmt.Println("ocp")

	creditCardProcessor := &CreditCardProcessor{}
	upiProcessor := &UPIProcessor{}

	MakePayment(creditCardProcessor, 100.50)
	MakePayment(upiProcessor, 200.75)

	//create a new processor
	payPalProcessor := &PayPalProcessor{}
	MakePayment(payPalProcessor, 100.75)
}