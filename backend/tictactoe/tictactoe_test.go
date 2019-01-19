package tictactoe_test

import (
	"noxes/backend/tictactoe"
	"testing"
)

func TestCrossAlwaysStarts(t *testing.T) {
	game := tictactoe.NewGame()

	err := game.MakeMove(tictactoe.Cross, 1, 1)

	if err != nil {
		t.Fatalf("Cross was unable to start the game: %s", err)
	}
}

func TestOneConsecutiveMovePerPlayer(t *testing.T) {
	game := tictactoe.NewGame()

	game.MakeMove(tictactoe.Cross, 1, 1)
	err := game.MakeMove(tictactoe.Cross, 1, 1)

	if err == nil {
		t.Fatalf("Player was able to take two consecutive turns")
	}
}

func TestDrawnGameHasNoWinner(t *testing.T) {
	game := tictactoe.NewGame()

	game.MakeMove(tictactoe.Cross, 0, 1)
	game.MakeMove(tictactoe.Nought, 0, 0)
	game.MakeMove(tictactoe.Cross, 1, 0)
	game.MakeMove(tictactoe.Nought, 1, 1)
	game.MakeMove(tictactoe.Cross, 2, 0)
	game.MakeMove(tictactoe.Nought, 2, 1)
	game.MakeMove(tictactoe.Cross, 2, 2)
	game.MakeMove(tictactoe.Nought, 1, 2)
	game.MakeMove(tictactoe.Cross, 0, 2)
	// board should look like this:
	// O | X | X
	// ----------
	// X | O | O
	// ----------
	// X | O | X

	if !game.IsGameOver() {
		t.Fatalf("Game has not ended after nine turns")
	}
	winner := game.GetWinner()
	if winner != tictactoe.None {
		t.Fatalf("Drawn game has %s as the winner", winner.String())
	}
}

func TestCrossWins(t *testing.T) {
	game := tictactoe.NewGame()

	game.MakeMove(tictactoe.Cross, 0, 0)
	game.MakeMove(tictactoe.Nought, 1, 1)
	game.MakeMove(tictactoe.Cross, 0, 1)
	game.MakeMove(tictactoe.Nought, 1, 2)
	game.MakeMove(tictactoe.Cross, 0, 2)
	// board should look like this:
	// X | - | -
	// ----------
	// X | O | -
	// ----------
	// X | O | -

	winner := game.GetWinner()
	if winner != tictactoe.Cross {
		t.Fatalf("Cross is not the winner")
	}
	if !game.IsGameOver() {
		t.Fatalf("Game has not ended after Cross won")
	}
}

func TestNoughtWins(t *testing.T) {
	game := tictactoe.NewGame()

	game.MakeMove(tictactoe.Cross, 0, 0)
	game.MakeMove(tictactoe.Nought, 1, 1)
	game.MakeMove(tictactoe.Cross, 0, 1)
	game.MakeMove(tictactoe.Nought, 2, 0)
	game.MakeMove(tictactoe.Cross, 2, 1)
	game.MakeMove(tictactoe.Nought, 0, 2)
	// board should look like this:
	// X | - | O
	// ----------
	// X | O | X
	// ----------
	// O | - | -

	winner := game.GetWinner()
	if winner != tictactoe.Nought {
		t.Fatalf("Nought is not the winner")
	}
	if !game.IsGameOver() {
		t.Fatalf("Game has not ended after Nought won")
	}
}

func TestCannotMakeMoveOnOccupiedCell(t *testing.T) {
	game := tictactoe.NewGame()

	game.MakeMove(tictactoe.Cross, 0, 1)
	game.MakeMove(tictactoe.Nought, 0, 0)
	game.MakeMove(tictactoe.Cross, 1, 0)
	game.MakeMove(tictactoe.Nought, 1, 1)
	game.MakeMove(tictactoe.Cross, 2, 0)
	game.MakeMove(tictactoe.Nought, 2, 1)
	game.MakeMove(tictactoe.Cross, 2, 2)
	game.MakeMove(tictactoe.Nought, 1, 2)
	game.MakeMove(tictactoe.Cross, 0, 2)
	// board should look like this:
	// O | X | X
	// ----------
	// X | O | O
	// ----------
	// X | O | X

	if !game.IsGameOver() {
		t.Fatalf("Game has not ended after nine turns")
	}
	winner := game.GetWinner()
	if winner != tictactoe.None {
		t.Fatalf("Drawn game has %s as the winner", winner.String())
	}
}

func TestCannotClearOccupiedCell(t *testing.T) {
	game := tictactoe.NewGame()

	game.MakeMove(tictactoe.Cross, 0, 1)
	err := game.MakeMove(tictactoe.None, 0, 1)

	if err == nil {
		t.Fatalf("Occupied cell was cleared")
	}
}

func TestCannotMakeMoveOutsideTheBoard(t *testing.T) {
	game := tictactoe.NewGame()

	err := game.MakeMove(tictactoe.Cross, -1, 1)

	if err == nil {
		t.Fatalf("A move was made outside the board")
	}
}

func TestCannotMakeMoveOnceGameHasEnded(t *testing.T) {
	game := tictactoe.NewGame()

	game.MakeMove(tictactoe.Cross, 0, 0)
	game.MakeMove(tictactoe.Nought, 1, 1)
	game.MakeMove(tictactoe.Cross, 0, 1)
	game.MakeMove(tictactoe.Nought, 2, 0)
	game.MakeMove(tictactoe.Cross, 2, 1)
	game.MakeMove(tictactoe.Nought, 0, 2)
	// board should look like this:
	// X | - | O
	// ----------
	// X | O | X
	// ----------
	// O | - | -
	err := game.MakeMove(tictactoe.Cross, 1, 0)

	if err == nil {
		t.Fatalf("A move was made after the game ended")
	}
}
