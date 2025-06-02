package factory

import (
	"errors"
	"tictactoe/model"
	"tictactoe/strategy"
)

func CreatePlayer(name string, symbol string, isBot bool) (model.Player, error) {
	if symbol != "X" && symbol != "O" {
		return nil, errors.New("invalid symbol")
	}

	if isBot {
		return &model.BasePlayer{
			Name:     name,
			Symbol:   symbol,
			Strategy: &strategy.BotStrategy{},
		}, nil
	}

	return &model.BasePlayer{
		Name:     name,
		Symbol:   symbol,
		Strategy: &strategy.HumanStrategy{},
	}, nil
}