package tictactoe

import (
	"errors"
)

type Location int

const (
	TopLeft     Location = 0
	Top         Location = 1
	TopRight    Location = 2
	Left        Location = 3
	Centre      Location = 4
	Right       Location = 5
	BottomLeft  Location = 6
	Bottom      Location = 7
	BottomRight Location = 8
)

type Cell string

const (
	None   Cell = "-"
	Nought Cell = "O"
	Cross  Cell = "X"
)

type Game struct {
	board      string
	turnsTaken int
	nextTurn   Cell
	gameOver   bool
	winner     Cell
}

func NewGame() Game {
	return Game{
		board:      "---------",
		turnsTaken: 0,
		nextTurn:   Cross,
		gameOver:   false,
		winner:     None,
	}
}

func (g Game) IsGameOver() bool {
	return g.gameOver
}

func (g Game) GetWinner() Cell {
	return g.winner
}

func (g Game) GetBoard() string {
	return g.board
}

func (g Game) GetNextTurn() Cell {
	return g.nextTurn
}

func (g *Game) MakeMove(cell Cell, loc Location) error {
	err := g.verifyMove(cell, loc)
	if err != nil {
		return err
	}

	g.board = g.board[:loc] + string(cell) + g.board[loc+1:]
	g.turnsTaken++
	if cell == Cross {
		g.nextTurn = Nought
	} else {
		g.nextTurn = Cross
	}

	g.updateOutcome(cell)

	return nil
}

func (g Game) verifyMove(cell Cell, loc Location) error {
	if loc < 0 || loc > 8 {
		return errors.New("Cannot make a move on an invalid location")
	} else if Cell(g.board[loc]) != None {
		return errors.New("Cannot make a move on an occupied location")
	} else if cell == None {
		return errors.New("Move must be Noughts or Crosses")
	} else if cell != g.nextTurn {
		return errors.New("Can only make a move on your turn")
	} else if g.gameOver {
		return errors.New("Cannot make a move on a finished game")
	}
	return nil
}

func (g *Game) updateOutcome(candidate Cell) {
	c := string(candidate)[0]
	b := g.board
	if (c == b[TopLeft] && b[TopLeft] == b[Top] && b[Top] == b[TopRight]) || // top row
		(c == b[Left] && b[Left] == b[Centre] && b[Centre] == b[Right]) || // middle row
		(c == b[BottomLeft] && b[BottomLeft] == b[Bottom] && b[Bottom] == b[BottomRight]) || // bottom row
		(c == b[TopLeft] && b[TopLeft] == b[Left] && b[Left] == b[BottomLeft]) || // left column
		(c == b[Top] && b[Top] == b[Centre] && b[Centre] == b[Bottom]) || // middle column
		(c == b[TopRight] && b[TopRight] == b[Right] && b[Right] == b[BottomRight]) || // right column
		(c == b[TopLeft] && b[TopLeft] == b[Centre] && b[Centre] == b[BottomRight]) || // back-slash diagonal ( \ )
		(c == b[BottomLeft] && b[BottomLeft] == b[Centre] && b[Centre] == b[TopRight]) { // forward-slash diagonal ( / )
		g.winner = candidate
	}

	if g.turnsTaken == 9 || g.winner != None {
		g.gameOver = true
	}
}
