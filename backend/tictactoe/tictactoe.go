package tictactoe

import (
	"errors"
	"fmt"
)

type Cell int

const (
	None   Cell = 0
	Nought Cell = -1
	Cross  Cell = 1
)

func (c Cell) String() string {
	switch c {
	case Nought:
		return "Nought"
	case Cross:
		return "Cross"
	case None:
		return "None"
	default:
		return "Unknown"
	}
}

func (c Cell) ShortString() string {
	switch c {
	case Nought:
		return "O"
	case Cross:
		return "X"
	case None:
		return "-"
	default:
		return "?"
	}
}

type Game struct {
	board    [3][3]Cell
	winner   Cell
	nextTurn Cell
}

func NewGame() Game {
	return Game{
		board: [3][3]Cell{
			[3]Cell{None, None, None},
			[3]Cell{None, None, None},
			[3]Cell{None, None, None},
		},
		winner:   None,
		nextTurn: Cross,
	}
}

func (g *Game) MakeMove(cell Cell, x int, y int) error {
	if g.board[y][x] != None {
		return errors.New("cell already taken")
	} else if cell == None {
		return errors.New("only Nought or Cross may be played")
	} else if cell != g.nextTurn {
		return fmt.Errorf("%s tried to make a move, but it's %s's turn",
			cell.String(), g.nextTurn.String())
	} else if g.IsGameOver() {
		return errors.New("cannot make moves on a finished game")
	}

	g.board[y][x] = cell
	g.nextTurn = -g.nextTurn // able to do this as long as this is true: Nought == -Cross. // TODO: test this.

	// TODO: check for winner

	return nil
}

func (g Game) IsGameOver() bool {
	return false // TODO
}

func (g Game) GetWinner() Cell {
	return None // TODO
}

func (g Game) PrintBoard() {
	fmt.Printf("%s | %s | %s\n", g.board[0][0].ShortString(), g.board[0][1].ShortString(), g.board[0][2].ShortString())
	fmt.Println("----------")
	fmt.Printf("%s | %s | %s\n", g.board[1][0].ShortString(), g.board[1][1].ShortString(), g.board[1][2].ShortString())
	fmt.Println("----------")
	fmt.Printf("%s | %s | %s\n", g.board[2][0].ShortString(), g.board[2][1].ShortString(), g.board[2][2].ShortString())
}
