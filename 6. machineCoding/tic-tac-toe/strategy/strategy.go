package strategy

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MoveStrategy interface {
	GetMove(board [][]string) (int, int, error)
}

// For CLI input
type HumanStrategy struct{}

func (h *HumanStrategy) GetMove(board [][]string) (int, int, error) {
	fmt.Print("Enter your move (row,col): ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	parts := strings.Split(input, ",")

	if len(parts) != 2 {
		return -1, -1, fmt.Errorf("invalid input")
	}

	row, err1 := strconv.Atoi(parts[0])
	col, err2 := strconv.Atoi(parts[1])

	if err1 != nil || err2 != nil {
		return -1, -1, fmt.Errorf("invalid numbers")
	}

	return row, col, nil
}

// Naive bot: pick first available cell
type BotStrategy struct{}

func (b *BotStrategy) GetMove(board [][]string) (int, int, error) {
	for i := range board {
		for j := range board[i] {
			if board[i][j] == "-" {
				return i, j, nil
			}
		}
	}
	return -1, -1, fmt.Errorf("no moves left")
}