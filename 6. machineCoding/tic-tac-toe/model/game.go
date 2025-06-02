package model

import (
	"tictactoe"
)

type Game struct {
	Board          *Board
	Players        [2]Player
	CurrentPlayer  int
	Status         string
	Winner         Player
}

func NewGame(player1, player2 Player, boardSize int) *Game {
	return &Game{
		Board:         NewBoard(boardSize),
		Players:       [2]Player{player1, player2},
		CurrentPlayer: 0,
		Status:        tictactoe.IN_PROGRESS,
	}
}

func (g *Game) GetCurrentPlayer() Player {
	return g.Players[g.CurrentPlayer]
}

func (g *Game) SwitchTurn() {
	g.CurrentPlayer = 1 - g.CurrentPlayer
}