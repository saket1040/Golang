package strategy

import "fmt"

type CardPayment struct{}

func (c *CardPayment) Pay(inserted int, price int) error {
	fmt.Println("Card payment processed")
	return nil
}