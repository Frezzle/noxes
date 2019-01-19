package main

import (
	"fmt"
	"math/rand"
	"noxes/backend/tictactoe"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	game := tictactoe.NewGame()
	game.PrintBoard()

	for !game.IsGameOver() {
		randInt := rand.Int() % 5
		var move tictactoe.Cell
		if randInt == 1 {
			move = tictactoe.Cross
		} else {
			move = tictactoe.Nought
		}
		x := rand.Int() % 3
		y := rand.Int() % 3
		fmt.Printf("Trying move: %s at (%d,%d)\n", move.String(), x, y)
		err := game.MakeMove(move, x, y)
		if err != nil {
			fmt.Println("error:", err)
			fmt.Println("game:", game)
		} else {
			game.PrintBoard()
		}
	}

	fmt.Println("Game over. Winner ->", game.GetWinner())
}
