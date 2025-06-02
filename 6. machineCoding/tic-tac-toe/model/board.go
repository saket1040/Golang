package model

import (
	"fmt"
	"tictactoe"
)

type Board struct {
	Size  int
	Cells [][]string
}

func NewBoard(size int) *Board {
	cells := make([][]string, size)
	for i := range cells {
		cells[i] = make([]string, size)
		for j := range cells[i] {
			cells[i][j] = tictactoe.EMPTY
		}
	}
	return &Board{
		Size:  size,
		Cells: cells,
	}
}

func (b *Board) Print() {
	for _, row := range b.Cells {
		for _, cell := range row {
			fmt.Printf("[%s]", cell)
		}
		fmt.Println()
	}
	fmt.Println()
}

func (b *Board) IsCellEmpty(row, col int) bool {
	return b.Cells[row][col] == tictactoe.EMPTY
}

func (b *Board) MarkCell(row, col int, symbol string) {
	b.Cells[row][col] = symbol
}