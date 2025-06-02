package strategy

import (
    "errors"
    "splitwise/model"
)

type EqualSplit struct {}

func (s *EqualSplit) Split(amount float64, paidBy string, participants []string) ([]model.Split, error) {
    if len(participants) == 0 {
        return nil, errors.New("no participants")
    }
    share := amount / float64(len(participants))
    splits := make([]model.Split, len(participants))
    for i, id := range participants {
        splits[i] = model.Split{
            UserID: id,
            Amount: share,
        }
    }
    return splits, nil
}