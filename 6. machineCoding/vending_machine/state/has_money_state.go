package state

import (
	"fmt"
	"vending_machine/service"
)

type HasMoneyState struct {
	vm *service.VendingMachine
}

func NewHasMoneyState(vm *service.VendingMachine) *HasMoneyState {
	return &HasMoneyState{vm: vm}
}

func (s *HasMoneyState) InsertMoney(amount int) error {
	if s.vm.SelectedItem == nil {
		return fmt.Errorf("no item selected")
	}
	if amount < s.vm.SelectedItem.Price {
		return fmt.Errorf("insufficient amount")
	}
	s.vm.CurrentBalance = amount
	return nil
}

func (s *HasMoneyState) Dispense() error {
	if s.vm.PaymentStrategy == nil {
		return fmt.Errorf("no payment strategy set")
	}
	if err := s.vm.PaymentStrategy.Pay(s.vm.CurrentBalance, s.vm.SelectedItem.Price); err != nil {
		return err
	}

	item := s.vm.SelectedItem
	item.Quantity -= 1
	s.vm.Items[item.Code] = *item
	s.vm.CurrentBalance = 0
	s.vm.SetState(NewIdleState(s.vm))
	fmt.Println("Dispensed:", item.Name)
	return nil
}