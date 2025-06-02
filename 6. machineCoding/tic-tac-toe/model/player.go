package model

import "tictactoe/strategy"

type Player interface {
	GetSymbol() string
	GetName() string
	GetMove(board [][]string) (int, int, error)
}

type BasePlayer struct {
	Name     string
	Symbol   string
	Strategy strategy.MoveStrategy
}

func (p *BasePlayer) GetSymbol() string {
	return p.Symbol
}

func (p *BasePlayer) GetName() string {
	return p.Name
}

func (p *BasePlayer) GetMove(board [][]string) (int, int, error) {
	return p.Strategy.GetMove(board)
}