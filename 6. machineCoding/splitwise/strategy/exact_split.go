package strategy

import "splitwise/model"

type ExactSplit struct {
    SplitMap map[string]float64
}

func (s *ExactSplit) Split(amount float64, paidBy string, participants []string) ([]model.Split, error) {
    splits := []model.Split{}
    for _, userID := range participants {
        amt, ok := s.SplitMap[userID]
        if !ok {
            continue
        }
        splits = append(splits, model.Split{UserID: userID, Amount: amt})
    }
    return splits, nil
}