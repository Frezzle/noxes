package game_manager_test

import (
	"noxes/backend/game_manager"
	"noxes/backend/tictactoe"
	"noxes/backend/utils"
	"testing"
)

func TestCreatingFiveNewGamesResultsInFiveGamesBeingManaged(t *testing.T) {
	gm := game_manager.NewGameManager()
	for i := 0; i < 5; i++ {
		gm.CreateNewGame()
	}

	numGames := len(gm.GetAllGames())

	if numGames != 5 {
		t.Fatalf("Expected 5 games to be created; got %d", numGames)
	}
}

func TestGameIdsShouldBeUnique(t *testing.T) {
	gm := game_manager.NewGameManager()
	ids := make([]int, 10)

	for i := 0; i < 10; i++ {
		ids[i] = gm.CreateNewGame()
	}

	expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	if !utils.SlicesAreEqual(ids, expected) {
		t.Fatal("Game IDs are not unique")
	}
}

func TestGetExistingGameReturnsGameWithSameId(t *testing.T) {
	gm := game_manager.NewGameManager()
	for i := 0; i < 3; i++ {
		gm.CreateNewGame()
	}

	desiredGameId := 1
	game, err := gm.GetGameById(desiredGameId)

	if err != nil {
		t.Fatalf("Failed to get game by ID: %s", err.Error())
	} else if game.GetId() != desiredGameId {
		t.Fatalf("Expected game with ID %d but got %d", desiredGameId, game.GetId())
	}
}

func TestGetNonExistingGameByIdReturnsError(t *testing.T) {
	gm := game_manager.NewGameManager()

	desiredGameId := 100
	_, err := gm.GetGameById(desiredGameId)

	if err == nil {
		t.Fatalf("Found game with ID %d; expected to find nothing and error to return", desiredGameId)
	}
}

func TestMakingAMoveOnExistingGameUpdatesTheGame(t *testing.T) {
	gm := game_manager.NewGameManager()
	gameId := gm.CreateNewGame()

	gm.MakeMove(gameId, tictactoe.Cross, tictactoe.Centre)

	game, _ := gm.GetGameById(gameId)
	board := game.GetBoard()
	expected := "----X----"
	if board != expected {
		t.Fatalf(`Expected board to be "%s" but got "%s"`, expected, board)
	}
}

func TestMakingAMoveOnNonExistingGameReturnsError(t *testing.T) {
	gm := game_manager.NewGameManager()

	err := gm.MakeMove(0, tictactoe.Cross, tictactoe.Centre)

	if err == nil {
		t.Fatal("No error raised when making move on non-existing game")
	}
}

func TestWinningAGameUpdatesWinnerAndGameOver(t *testing.T) {
	gm := game_manager.NewGameManager()
	gameId := gm.CreateNewGame()

	gm.MakeMove(gameId, tictactoe.Cross, tictactoe.Centre)
	gm.MakeMove(gameId, tictactoe.Nought, tictactoe.Left)
	gm.MakeMove(gameId, tictactoe.Cross, tictactoe.Top)
	gm.MakeMove(gameId, tictactoe.Nought, tictactoe.Right)
	gm.MakeMove(gameId, tictactoe.Cross, tictactoe.Bottom)
	// board should look like this:
	// - | X | -
	// ----------
	// O | X | O
	// ----------
	// - | X | -

	game, _ := gm.GetGameById(gameId)
	winner := game.GetWinner()
	expected := tictactoe.Cross
	if winner != expected {
		t.Fatalf(`Expected winner to be "%s" but got "%s"`, expected, winner)
	}
}

func TestWinningAGameEndsTheGame(t *testing.T) {
	gm := game_manager.NewGameManager()
	gameId := gm.CreateNewGame()

	gm.MakeMove(gameId, tictactoe.Cross, tictactoe.Centre)
	gm.MakeMove(gameId, tictactoe.Nought, tictactoe.Left)
	gm.MakeMove(gameId, tictactoe.Cross, tictactoe.Top)
	gm.MakeMove(gameId, tictactoe.Nought, tictactoe.Right)
	gm.MakeMove(gameId, tictactoe.Cross, tictactoe.Bottom)
	// board should look like this:
	// - | X | -
	// ----------
	// O | X | O
	// ----------
	// - | X | -

	game, _ := gm.GetGameById(gameId)
	if !game.IsGameOver() {
		t.Fatal("Game should have ended after player won")
	}
}
