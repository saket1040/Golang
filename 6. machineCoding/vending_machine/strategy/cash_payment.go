package strategy

import "fmt"

type CashPayment struct{}

func (c *CashPayment) Pay(inserted int, price int) error {
	if inserted < price {
		return fmt.Errorf("insufficient funds")
	}
	fmt.Printf("Cash accepted. Change: %d\n", inserted-price)
	return nil
}