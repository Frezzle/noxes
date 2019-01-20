package tictactoe_test

import (
	"noxes/tictactoe"
	"testing"
)

func TestNewGameHasClearBoard(t *testing.T) {
	game := tictactoe.NewGame()

	board := game.GetBoard()

	if board != "---------" {
		t.Fatal("Game board is not initially cleared")
	}
}

func TestCrossAlwaysStarts(t *testing.T) {
	game := tictactoe.NewGame()

	err := game.MakeMove(tictactoe.Cross, tictactoe.Centre)

	if err != nil {
		t.Fatalf("Cross was unable to start the game: %s", err)
	}
}

func TestOneConsecutiveMovePerPlayer(t *testing.T) {
	game := tictactoe.NewGame()

	game.MakeMove(tictactoe.Cross, tictactoe.Centre)
	err := game.MakeMove(tictactoe.Cross, tictactoe.Centre)

	if err == nil {
		t.Fatal("Player was able to take two consecutive turns")
	}
}

func TestDrawnGameHasNoWinner(t *testing.T) {
	game := tictactoe.NewGame()

	game.MakeMove(tictactoe.Cross, tictactoe.Left)
	game.MakeMove(tictactoe.Nought, tictactoe.TopLeft)
	game.MakeMove(tictactoe.Cross, tictactoe.Top)
	game.MakeMove(tictactoe.Nought, tictactoe.Centre)
	game.MakeMove(tictactoe.Cross, tictactoe.TopRight)
	game.MakeMove(tictactoe.Nought, tictactoe.Right)
	game.MakeMove(tictactoe.Cross, tictactoe.BottomRight)
	game.MakeMove(tictactoe.Nought, tictactoe.Bottom)
	game.MakeMove(tictactoe.Cross, tictactoe.BottomLeft)
	// board should look like this:
	// O | X | X
	// ----------
	// X | O | O
	// ----------
	// X | O | X

	if !game.IsGameOver() {
		t.Fatal("Game has not ended after nine turns")
	}
	winner := game.GetWinner()
	if winner != tictactoe.None {
		t.Fatalf("Drawn game has %s as the winner", winner)
	}
}

func TestCrossWins(t *testing.T) {
	game := tictactoe.NewGame()

	game.MakeMove(tictactoe.Cross, tictactoe.TopLeft)
	game.MakeMove(tictactoe.Nought, tictactoe.Centre)
	game.MakeMove(tictactoe.Cross, tictactoe.Left)
	game.MakeMove(tictactoe.Nought, tictactoe.Bottom)
	game.MakeMove(tictactoe.Cross, tictactoe.BottomLeft)
	// board should look like this:
	// X | - | -
	// ----------
	// X | O | -
	// ----------
	// X | O | -

	winner := game.GetWinner()
	if winner != tictactoe.Cross {
		t.Fatal("Cross is not the winner")
	}
	if !game.IsGameOver() {
		t.Fatal("Game has not ended after Cross won")
	}
}

func TestNoughtWins(t *testing.T) {
	game := tictactoe.NewGame()

	game.MakeMove(tictactoe.Cross, tictactoe.TopLeft)
	game.MakeMove(tictactoe.Nought, tictactoe.Centre)
	game.MakeMove(tictactoe.Cross, tictactoe.Left)
	game.MakeMove(tictactoe.Nought, tictactoe.TopRight)
	game.MakeMove(tictactoe.Cross, tictactoe.Right)
	game.MakeMove(tictactoe.Nought, tictactoe.BottomLeft)
	// board should look like this:
	// X | - | O
	// ----------
	// X | O | X
	// ----------
	// O | - | -

	winner := game.GetWinner()
	if winner != tictactoe.Nought {
		t.Fatal("Nought is not the winner")
	}
	if !game.IsGameOver() {
		t.Fatal("Game has not ended after Nought won")
	}
}

func TestCannotMakeMoveOnOccupiedCell(t *testing.T) {
	game := tictactoe.NewGame()

	game.MakeMove(tictactoe.Cross, tictactoe.Left)
	game.MakeMove(tictactoe.Nought, tictactoe.TopLeft)
	game.MakeMove(tictactoe.Cross, tictactoe.Top)
	game.MakeMove(tictactoe.Nought, tictactoe.Centre)
	game.MakeMove(tictactoe.Cross, tictactoe.TopRight)
	game.MakeMove(tictactoe.Nought, tictactoe.Right)
	game.MakeMove(tictactoe.Cross, tictactoe.BottomRight)
	game.MakeMove(tictactoe.Nought, tictactoe.Bottom)
	game.MakeMove(tictactoe.Cross, tictactoe.BottomLeft)
	// board should look like this:
	// O | X | X
	// ----------
	// X | O | O
	// ----------
	// X | O | X

	if !game.IsGameOver() {
		t.Fatal("Game has not ended after nine turns")
	}
	winner := game.GetWinner()
	if winner != tictactoe.None {
		t.Fatalf("Drawn game has %s as the winner", winner)
	}
}

func TestCannotClearOccupiedCell(t *testing.T) {
	game := tictactoe.NewGame()

	game.MakeMove(tictactoe.Cross, tictactoe.Left)
	err := game.MakeMove(tictactoe.None, tictactoe.Left)

	if err == nil {
		t.Fatal("Occupied cell was cleared")
	}
}

func TestCannotMakeMoveOutsideTheBoard(t *testing.T) {
	game := tictactoe.NewGame()

	err := game.MakeMove(tictactoe.Cross, -1)

	if err == nil {
		t.Fatal("A move was allowed outside the board")
	}
}

func TestCannotMakeMoveOnceGameHasEnded(t *testing.T) {
	game := tictactoe.NewGame()

	game.MakeMove(tictactoe.Cross, tictactoe.TopLeft)
	game.MakeMove(tictactoe.Nought, tictactoe.Centre)
	game.MakeMove(tictactoe.Cross, tictactoe.Left)
	game.MakeMove(tictactoe.Nought, tictactoe.TopRight)
	game.MakeMove(tictactoe.Cross, tictactoe.Right)
	game.MakeMove(tictactoe.Nought, tictactoe.BottomLeft)
	// board should look like this:
	// X | - | O
	// ----------
	// X | O | X
	// ----------
	// O | - | -
	err := game.MakeMove(tictactoe.Cross, tictactoe.Top)

	if err == nil {
		t.Fatal("A move was allowed after the game ended")
	}
}
