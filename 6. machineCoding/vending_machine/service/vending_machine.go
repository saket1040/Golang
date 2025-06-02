package service

import (
	"fmt"
	"vending_machine/model"
	"vending_machine/strategy"
	"vending_machine/stateifc"
)

type VendingMachine struct {
	Items            map[string]model.Item
	SelectedItem     *model.Item
	CurrentBalance   int
	State            stateifc.State
	PaymentStrategy  strategy.PaymentStrategy
}

func NewVendingMachine() *VendingMachine {
    return &VendingMachine{
        Items: make(map[string]model.Item),
    }
}

func (v *VendingMachine) SetState(s stateifc.State) {
	v.State = s
}

func (v *VendingMachine) SetPaymentStrategy(p strategy.PaymentStrategy) {
	v.PaymentStrategy = p
}

func (v *VendingMachine) AddItem(item model.Item) {
	v.Items[item.Code] = item
}

func (v *VendingMachine) SelectItem(code string) bool {
    item, ok := v.Items[code]
    if !ok || item.Quantity == 0 {
        fmt.Println("Item not available")
        return false
    }
    v.SelectedItem = &item
    fmt.Println("Item Selected:", item.Name)
    return true
}

func (v *VendingMachine) InsertMoney(amount int) error {
	return v.State.InsertMoney(amount)
}

func (v *VendingMachine) Dispense() error {
	return v.State.Dispense()
}