package state

import (
	"errors"
	"vending_machine/service"
)

type IdleState struct {
	vm *service.VendingMachine
}

func NewIdleState(vm *service.VendingMachine) *IdleState {
	return &IdleState{vm: vm}
}

func (s *IdleState) InsertMoney(amount int) error {
	return errors.New("select item first")
}

func (s *IdleState) Dispense() error {
	return errors.New("select item and pay first")
}