package main

import (
	"fmt"
	constants "tictactoe"
	"tictactoe/factory"
	"tictactoe/model"
	"tictactoe/service"
)
// fyi : wrong implementation

func main() {
	fmt.Println("Welcome to Tic Tac Toe")

	player1, err := factory.CreatePlayer("Alice", constants.X, false) // Human
	if err != nil {
		panic(err)
	}

	player2, err := factory.CreatePlayer("BotBob", constants.O, true) // Bot
	if err != nil {
		panic(err)
	}

	game := model.NewGame(player1, player2, 3)
	gameService := service.NewGameService(game)

	for {
		game.Board.Print()

		row, col, err := game.Players[game.CurrentPlayer].GetMove(game.Board.Cells)
		if err != nil {
			fmt.Println("Invalid move:", err)
			continue
		}

		if err = gameService.MakeMove(row, col); err != nil {
			fmt.Println("Cell already occupied. Try again.")
			continue
		}
		
		

		if gameService.CheckWin(row, col, game.Players[game.CurrentPlayer].GetSymbol()){
			game.Board.Print()
			fmt.Printf("Player %s wins!\n", game.Players[game.CurrentPlayer].GetName())
			break
		}

		if gameService.IsDraw() {
			game.Board.Print()
			fmt.Println("It's a draw!")
			break
		}
	}
}
