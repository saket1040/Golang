package main

import (
	"fmt"
)

type Account struct {
	Name    string
	Balance float64
}

func (a *Account) Deposit(amount float64) {
	fmt.Println("Deposited", amount)
	a.Balance += amount
}

func (a *Account) WithDraw(amount float64) {
	fmt.Println("WithDraw", amount)
	a.Balance -= amount
}

func (a *Account) GetBalance() float64 {
	return  a.Balance
}

func main() {
	fmt.Println("checking bank")
	ac := &Account{
		Name: "Saket",
		Balance: 998.09,
	}
	ac.Deposit(88)
	ac.WithDraw(99)
	fmt.Println(ac.GetBalance())
}