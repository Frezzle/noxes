package tictactoe

import (
	"testing"
)

func TestNewGameHasClearBoard(t *testing.T) {
	game := NewGame()

	board := game.GetBoard()

	if board != "---------" {
		t.Fatal("Game board is not initially cleared")
	}
}

func TestCrossAlwaysStarts(t *testing.T) {
	game := NewGame()

	err := game.MakeMove(Cross, Centre)

	if err != nil {
		t.Fatalf("Cross was unable to start the game: %s", err)
	}
}

func TestOneConsecutiveMovePerPlayer(t *testing.T) {
	game := NewGame()

	game.MakeMove(Cross, Centre)
	err := game.MakeMove(Cross, Centre)

	if err == nil {
		t.Fatal("Player was able to take two consecutive turns")
	}
}

func TestDrawnGameHasNoWinner(t *testing.T) {
	game := NewGame()

	game.MakeMove(Cross, Left)
	game.MakeMove(Nought, TopLeft)
	game.MakeMove(Cross, Top)
	game.MakeMove(Nought, Centre)
	game.MakeMove(Cross, TopRight)
	game.MakeMove(Nought, Right)
	game.MakeMove(Cross, BottomRight)
	game.MakeMove(Nought, Bottom)
	game.MakeMove(Cross, BottomLeft)
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
	if winner != None {
		t.Fatalf("Drawn game has %s as the winner", winner)
	}
}

func TestCrossWins(t *testing.T) {
	game := NewGame()

	game.MakeMove(Cross, TopLeft)
	game.MakeMove(Nought, Centre)
	game.MakeMove(Cross, Left)
	game.MakeMove(Nought, Bottom)
	game.MakeMove(Cross, BottomLeft)
	// board should look like this:
	// X | - | -
	// ----------
	// X | O | -
	// ----------
	// X | O | -

	winner := game.GetWinner()
	if winner != Cross {
		t.Fatal("Cross is not the winner")
	}
	if !game.IsGameOver() {
		t.Fatal("Game has not ended after Cross won")
	}
}

func TestNoughtWins(t *testing.T) {
	game := NewGame()

	game.MakeMove(Cross, TopLeft)
	game.MakeMove(Nought, Centre)
	game.MakeMove(Cross, Left)
	game.MakeMove(Nought, TopRight)
	game.MakeMove(Cross, Right)
	game.MakeMove(Nought, BottomLeft)
	// board should look like this:
	// X | - | O
	// ----------
	// X | O | X
	// ----------
	// O | - | -

	winner := game.GetWinner()
	if winner != Nought {
		t.Fatal("Nought is not the winner")
	}
	if !game.IsGameOver() {
		t.Fatal("Game has not ended after Nought won")
	}
}

func TestCannotMakeMoveOnOccupiedCell(t *testing.T) {
	game := NewGame()

	game.MakeMove(Cross, Left)
	game.MakeMove(Nought, TopLeft)
	game.MakeMove(Cross, Top)
	game.MakeMove(Nought, Centre)
	game.MakeMove(Cross, TopRight)
	game.MakeMove(Nought, Right)
	game.MakeMove(Cross, BottomRight)
	game.MakeMove(Nought, Bottom)
	game.MakeMove(Cross, BottomLeft)
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
	if winner != None {
		t.Fatalf("Drawn game has %s as the winner", winner)
	}
}

func TestCannotClearOccupiedCell(t *testing.T) {
	game := NewGame()

	game.MakeMove(Cross, Left)
	err := game.MakeMove(None, Left)

	if err == nil {
		t.Fatal("Occupied cell was cleared")
	}
}

func TestCannotMakeMoveOutsideTheBoard(t *testing.T) {
	game := NewGame()

	err := game.MakeMove(Cross, -1)

	if err == nil {
		t.Fatal("A move was allowed outside the board")
	}
}

func TestCannotMakeMoveOnceGameHasEnded(t *testing.T) {
	game := NewGame()

	game.MakeMove(Cross, TopLeft)
	game.MakeMove(Nought, Centre)
	game.MakeMove(Cross, Left)
	game.MakeMove(Nought, TopRight)
	game.MakeMove(Cross, Right)
	game.MakeMove(Nought, BottomLeft)
	// board should look like this:
	// X | - | O
	// ----------
	// X | O | X
	// ----------
	// O | - | -
	err := game.MakeMove(Cross, Top)

	if err == nil {
		t.Fatal("A move was allowed after the game ended")
	}
}
