package service

import (
	"splitwise/model"
	"splitwise/storage"
	"splitwise/strategy"

	"github.com/google/uuid"
)

type ExpenseService struct {
	store    storage.Storage
	splitter strategy.SplitStrategy
}

func NewExpenseService(store storage.Storage, splitter strategy.SplitStrategy) *ExpenseService {
	return &ExpenseService{store: store, splitter: splitter}
}

func (e *ExpenseService) AddExpense(paidBy, groupID string, amount float64, description string) error {
	group, ok := e.store.GetGroup(groupID)
	if !ok {
		return nil
	}
	splits, err := e.splitter.Split(amount, paidBy, group.Members)
	if err != nil {
		return err
	}
	expense := model.Expense{
		ID:          uuid.New().String(),
		PaidBy:      paidBy,
		Amount:      amount,
		Splits:      splits,
		Description: description,
	}
	e.store.AddExpense(expense)
	group.Expenses = append(group.Expenses, expense.ID)
	e.store.UpdateGroup(group)
	return nil
}
