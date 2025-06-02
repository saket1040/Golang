package service

import (
	"errors"
	"tictactoe"
	"tictactoe/model"
)

type GameService struct {
	Game *model.Game
}
func NewGameService(game *model.Game) *GameService {
	return &GameService{Game: game}
}

func (gs *GameService) MakeMove(row, col int) error {
	player := gs.Game.GetCurrentPlayer()

	if !gs.Game.Board.IsCellEmpty(row, col) {
		return errors.New("cell is already occupied")
	}

	gs.Game.Board.MarkCell(row, col, player.GetSymbol())

	if gs.CheckWin(row, col, player.GetSymbol()) {
		gs.Game.Status = tictactoe.WON
		gs.Game.Winner = player
		return nil
	}

	if gs.IsDraw() {
		gs.Game.Status = tictactoe.DRAW
		return nil
	}

	gs.Game.SwitchTurn()
	return nil
}

func (gs *GameService) CheckWin(row, col int, symbol string) bool {
	board := gs.Game.Board
	size := board.Size
	cells := board.Cells

	// Check row
	win := true
	for c := 0; c < size; c++ {
		if cells[row][c] != symbol {
			win = false
			break
		}
	}
	if win {
		return true
	}

	// Check column
	win = true
	for r := 0; r < size; r++ {
		if cells[r][col] != symbol {
			win = false
			break
		}
	}
	if win {
		return true
	}

	// Check diagonal
	if row == col {
		win = true
		for i := 0; i < size; i++ {
			if cells[i][i] != symbol {
				win = false
				break
			}
		}
		if win {
			return true
		}
	}

	// Check anti-diagonal
	if row+col == size-1 {
		win = true
		for i := 0; i < size; i++ {
			if cells[i][size-1-i] != symbol {
				win = false
				break
			}
		}
		if win {
			return true
		}
	}

	return false
}

func (gs *GameService) IsDraw() bool {
	board := gs.Game.Board
	for i := 0; i < board.Size; i++ {
		for j := 0; j < board.Size; j++ {
			if board.Cells[i][j] == tictactoe.EMPTY {
				return false
			}
		}
	}
	return true
}