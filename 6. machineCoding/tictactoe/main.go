package main

import (
	"errors"
	"fmt"
	"strings"
)

// Constants for players and game status
const (
	EMPTY = " "
	PLAYER_X = "X"
	PLAYER_O = "O"

	IN_PROGRESS = "IN_PROGRESS"
	DRAW        = "DRAW"
	WIN         = "WIN"
)

type Player struct {
	Id     int
	Symbol string
}

type Cell struct {
	Row   int
	Col   int
	Value string
}

type Board struct {
	Size  int
	Cells [][]*Cell
}

func NewBoard(size int) *Board {
	cells := make([][]*Cell, size)
	for i := 0; i < size; i++ {
		cells[i] = make([]*Cell, size)
		for j := 0; j < size; j++ {
			cells[i][j] = &Cell{Row: i, Col: j, Value: EMPTY}
		}
	}
	return &Board{Size: size, Cells: cells}
}

func (b *Board) PrintBoard() {
	for i := 0; i < b.Size; i++ {
		for j := 0; j < b.Size; j++ {
			fmt.Printf(" %s ", b.Cells[i][j].Value)
			if j < b.Size-1 {
				fmt.Print("|")
			}
		}
		fmt.Println()
		if i < b.Size-1 {
			fmt.Println(strings.Repeat("---+", b.Size-1) + "---")
		}
	}
}

func (b *Board) IsFull() bool {
	for i := 0; i < b.Size; i++ {
		for j := 0; j < b.Size; j++ {
			if b.Cells[i][j].Value == EMPTY {
				return false
			}
		}
	}
	return true
}

func (b *Board) MakeMove(row, col int, symbol string) error {
	if row < 0 || row >= b.Size || col < 0 || col >= b.Size {
		return errors.New("invalid move")
	}
	if b.Cells[row][col].Value != EMPTY {
		return errors.New("cell already filled")
	}
	b.Cells[row][col].Value = symbol
	return nil
}

func (b *Board) CheckWin(symbol string) bool {
	size := b.Size
	// Check rows
	for i := 0; i < size; i++ {
		win := true
		for j := 0; j < size; j++ {
			if b.Cells[i][j].Value != symbol {
				win = false
				break
			}
		}
		if win {
			return true
		}
	}

	// Check columns
	for i := 0; i < size; i++ {
		win := true
		for j := 0; j < size; j++ {
			if b.Cells[j][i].Value != symbol {
				win = false
				break
			}
		}
		if win {
			return true
		}
	}

	// Check diagonal
	win := true
	for i := 0; i < size; i++ {
		if b.Cells[i][i].Value != symbol {
			win = false
			break
		}
	}
	if win {
		return true
	}

	// Check anti-diagonal
	win = true
	for i := 0; i < size; i++ {
		if b.Cells[i][size-i-1].Value != symbol {
			win = false
			break
		}
	}
	return win
}

type Game struct {
	Board       *Board
	Players     [2]*Player
	CurrentTurn int
	Status      string
	Winner      *Player
}

func NewGame() *Game {
	board := NewBoard(3)
	players := [2]*Player{
		{Id: 1, Symbol: PLAYER_X},
		{Id: 2, Symbol: PLAYER_O},
	}
	return &Game{Board: board, Players: players, CurrentTurn: 0, Status: IN_PROGRESS}
}

func (g *Game) Play(row, col int) error {
	if g.Status != IN_PROGRESS {
		return errors.New("game is already over")
	}
	player := g.Players[g.CurrentTurn]
	err := g.Board.MakeMove(row, col, player.Symbol)
	if err != nil {
		return err
	}

	if g.Board.CheckWin(player.Symbol) {
		g.Status = WIN
		g.Winner = player
		return nil
	}

	if g.Board.IsFull() {
		g.Status = DRAW
		return nil
	}

	g.CurrentTurn = 1 - g.CurrentTurn
	return nil
}

func main() {
	game := NewGame()
	var row, col int
	for game.Status == IN_PROGRESS {
		game.Board.PrintBoard()
		fmt.Printf("Player %s, enter your move (row col): ", game.Players[game.CurrentTurn].Symbol)
		fmt.Scanf("%d %d\n", &row, &col)
		err := game.Play(row, col)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}

	game.Board.PrintBoard()
	if game.Status == WIN {
		fmt.Printf("Player %s wins!\n", game.Winner.Symbol)
	} else {
		fmt.Println("It's a draw!")
	}
}
