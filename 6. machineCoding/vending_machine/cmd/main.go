package main

import (
	"fmt"
	"vending_machine/service"
	"vending_machine/model"
	"vending_machine/strategy"
	"vending_machine/state"
)

func main() {
	vm := service.NewVendingMachine()
	idle := state.NewIdleState(vm)
	vm.SetState(idle)
	vm.SetPaymentStrategy(&strategy.CashPayment{})
	
	item := model.Item{Code: "A1", Name: "Coke", Price: 25, Quantity: 10}
	vm.AddItem(item)

	// Select the item with code "A1"
	if vm.SelectItem("A1") {
		vm.SetState(state.NewHasMoneyState(vm))
	}

	// insert money
	err := vm.InsertMoney(30)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// item dispensed
	err = vm.Dispense()
	if err != nil {
		fmt.Println("Error:", err)
	}
}