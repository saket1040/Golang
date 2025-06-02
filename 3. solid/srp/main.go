package main

import "fmt"

type Customer struct {
	Name string
}

type Invoice struct {
	Customer_ Customer
	Amount    float64
}

func (i *Invoice) Print() {
	fmt.Println("Details Customer Name: ", i.Customer_.Name, " Amount: ", i.Amount)
}

type Saver struct {}

func (s *Saver) Save(inv *Invoice) {
	fmt.Println("Saving Invoice", inv.Customer_.Name, "amount", inv.Amount)
}

func main() {
	fmt.Println("lets go solid")

	inv := &Invoice{
		Customer_: Customer{
			Name: "Saket",
		},
		Amount: 990,
	}

	inv.Print()
	saver := &Saver{}
	saver.Save(inv)
}
