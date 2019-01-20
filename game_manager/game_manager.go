package game_manager

import (
	"fmt"
	"noxes/tictactoe"
)

type Game struct {
	id            int
	tictactoeGame tictactoe.Game
}

func (g Game) GetId() int {
	return g.id
}

func (g Game) GetBoard() string {
	return g.tictactoeGame.GetBoard()
}

func (g Game) GetNextTurn() tictactoe.Cell {
	return g.tictactoeGame.GetNextTurn()
}

func (g Game) IsGameOver() bool {
	return g.tictactoeGame.IsGameOver()
}

func (g Game) GetWinner() tictactoe.Cell {
	return g.tictactoeGame.GetWinner()
}

type GameManager struct {
	games      []Game
	nextGameId int
}

func NewGameManager() GameManager {
	return GameManager{
		games:      make([]Game, 0),
		nextGameId: 0,
	}
}

func (gm *GameManager) CreateNewGame() (gameId int) {
	gameId = gm.nextGameId
	gm.games = append(gm.games, Game{
		id:            gameId,
		tictactoeGame: tictactoe.NewGame(),
	})
	gm.nextGameId++
	return gameId
}

func (gm GameManager) GetAllGames() []Game {
	return gm.games
}

func (gm *GameManager) GetGameById(id int) (Game, error) {
	for i := range gm.games {
		if gm.games[i].id == id {
			return gm.games[i], nil
		}
	}
	return Game{}, fmt.Errorf("Game with ID %d not found", id)
}

func (gm *GameManager) MakeMove(gameId int, cell tictactoe.Cell, loc tictactoe.Location) error {
	for i := range gm.games {
		if gm.games[i].id == gameId {
			return gm.games[i].tictactoeGame.MakeMove(cell, loc)
		}
	}

	return fmt.Errorf("Game with ID %d not found", gameId)
}
