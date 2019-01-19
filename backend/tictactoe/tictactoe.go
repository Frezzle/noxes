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
	board      [3][3]Cell
	turnsTaken int
	nextTurn   Cell
	gameOver   bool
	winner     Cell
}

func NewGame() Game {
	return Game{
		board: [3][3]Cell{
			[3]Cell{None, None, None},
			[3]Cell{None, None, None},
			[3]Cell{None, None, None},
		},
		turnsTaken: 0,
		nextTurn:   Cross,
		gameOver:   false,
		winner:     None,
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
	} else if x < 0 || x > 2 || y < 0 || y > 2 {
		return errors.New("invalid location")
	} else if g.IsGameOver() {
		return errors.New("cannot make moves on a finished game")
	}

	g.board[y][x] = cell
	g.turnsTaken++
	g.nextTurn = -g.nextTurn // able to do this as long as Nought == -Cross. // TODO: test this.
	g.updateOutcome(cell)

	return nil
}

func (g *Game) updateOutcome(candidate Cell) {
	if (candidate == g.board[0][0] && candidate == g.board[0][1] && candidate == g.board[0][2]) || // top row
		(candidate == g.board[1][0] && candidate == g.board[1][1] && candidate == g.board[1][2]) || // middle row
		(candidate == g.board[2][0] && candidate == g.board[2][1] && candidate == g.board[2][2]) || // bottom row
		(candidate == g.board[0][0] && candidate == g.board[1][0] && candidate == g.board[2][0]) || // left column
		(candidate == g.board[0][1] && candidate == g.board[1][1] && candidate == g.board[2][1]) || // middle column
		(candidate == g.board[0][2] && candidate == g.board[1][2] && candidate == g.board[2][2]) || // right column
		(candidate == g.board[0][0] && candidate == g.board[1][1] && candidate == g.board[2][2]) || // back-slash diagonal ( \ )
		(candidate == g.board[2][0] && candidate == g.board[1][1] && candidate == g.board[0][2]) { // forward-slash diagonal ( / )
		g.winner = candidate
	}

	if g.turnsTaken == 9 || g.winner != None {
		g.gameOver = true
	}
}

func (g Game) IsGameOver() bool {
	return g.gameOver
}

func (g Game) GetWinner() Cell {
	return g.winner
}

func (g Game) PrintBoard() {
	fmt.Printf("%s | %s | %s\n", g.board[0][0].ShortString(), g.board[0][1].ShortString(), g.board[0][2].ShortString())
	fmt.Println("----------")
	fmt.Printf("%s | %s | %s\n", g.board[1][0].ShortString(), g.board[1][1].ShortString(), g.board[1][2].ShortString())
	fmt.Println("----------")
	fmt.Printf("%s | %s | %s\n", g.board[2][0].ShortString(), g.board[2][1].ShortString(), g.board[2][2].ShortString())
}
