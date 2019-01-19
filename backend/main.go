package main

import (
	"fmt"
	"noxes/backend/tictactoe"
)

func main() {
	game := tictactoe.NewGame()
	game.PrintBoard()

	err := game.MakeMove(tictactoe.Cross, 2, 0)
	fmt.Println("error:", err)
	game.PrintBoard()

	err = game.MakeMove(tictactoe.Nought, 1, 0)
	fmt.Println("error:", err)
	game.PrintBoard()

	err = game.MakeMove(tictactoe.Cross, 1, 1)
	fmt.Println("error:", err)
	game.PrintBoard()

	err = game.MakeMove(tictactoe.Nought, 0, 1)
	fmt.Println("error:", err)
	game.PrintBoard()

	err = game.MakeMove(tictactoe.Cross, 0, 2)
	fmt.Println("error:", err)
	game.PrintBoard()

	fmt.Println("winner ->", game.GetWinner())
}
