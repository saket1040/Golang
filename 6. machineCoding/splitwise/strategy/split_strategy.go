package strategy

import "splitwise/model"

type SplitStrategy interface {
    Split(amount float64, paidBy string, participants []string) ([]model.Split, error)
}