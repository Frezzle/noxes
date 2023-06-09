package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Frezzle/noxes/game_manager"
	"github.com/Frezzle/noxes/tictactoe"
)

func main() {
	http.HandleFunc("/game", getGameHandler)
	http.HandleFunc("/games", getAllGamesHandler)
	http.HandleFunc("/game/create", createGameHandler)
	http.HandleFunc("/game/move", makeMoveHandler)
	// TODO: more endpoints can be added to accomodate auth: (including updating above endpoints to require auth)
	// http.HandleFunc("/game/join", ...)
	// http.HandleFunc("/game/quit", ...)
	// http.HandleFunc("/user/register", ...)
	// http.HandleFunc("/user/login", ...)

	address := "localhost:9876"
	log.Printf("Listening on %s", address)
	log.Fatal(http.ListenAndServe(address, nil))
}

// TODO: Update model to use different data types to avoid so many type conversions.
type GameJSON struct {
	ID       string `json:"id"`
	Board    string `json:"board"`
	NextTurn string `json:"nextTurn"`
	GameOver string `json:"gameOver"`
	Winner   string `json:"winner"`
}

type GamesJSON []GameJSON

type IdJSON struct {
	ID string `json:"id"`
}

var gm = game_manager.NewGameManager()

func getGameHandler(w http.ResponseWriter, r *http.Request) { // TODO: GET body should not be used; get ID from somewhere else (e.g. URL, header, other).
	if !verifyRequest("GET", w, r) {
		return
	}

	var gameIdJSON IdJSON
	err := json.NewDecoder(r.Body).Decode(&gameIdJSON)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error":"%s"}`, err.Error()), http.StatusBadRequest)
		return
	}

	gameId, convErr := strconv.ParseInt(gameIdJSON.ID, 10, 32)
	if convErr != nil {
		http.Error(w, `{"error":"bad id provided"}`, http.StatusBadRequest)
		return
	}

	game, findGameErr := gm.GetGameById(int(gameId))
	if findGameErr != nil {
		http.Error(w, `{"error":"game does not exist"}`, http.StatusBadRequest)
		return
	}

	var result = GameJSON{
		ID:       gameIdJSON.ID,
		Board:    game.GetBoard(),
		NextTurn: string(game.GetNextTurn()),
		GameOver: strconv.FormatBool(game.IsGameOver()),
		Winner:   string(game.GetWinner()),
	}

	json.NewEncoder(w).Encode(result)
}

func getAllGamesHandler(w http.ResponseWriter, r *http.Request) {
	if !verifyRequest("GET", w, r) {
		return
	}

	games := gm.GetAllGames()
	gamesJSON := make([]GameJSON, 0, len(games))
	for _, game := range games {
		gamesJSON = append(gamesJSON, GameJSON{ // TODO: refactor into own function.
			ID:       strconv.FormatInt(int64(game.GetId()), 10),
			Board:    game.GetBoard(),
			NextTurn: string(game.GetNextTurn()),
			GameOver: strconv.FormatBool(game.IsGameOver()),
			Winner:   string(game.GetWinner()),
		})
	}

	json.NewEncoder(w).Encode(gamesJSON)
}

func createGameHandler(w http.ResponseWriter, r *http.Request) {
	if !verifyRequest("POST", w, r) {
		return
	}

	gameId := gm.CreateNewGame()

	var result = IdJSON{ID: strconv.FormatInt(int64(gameId), 10)}

	json.NewEncoder(w).Encode(result)
}

type MakeMoveJSON struct {
	GameID   string `json:"gameId"`
	Location string `json:"location"`
	Player   string `json:"player"`
}

func makeMoveHandler(w http.ResponseWriter, r *http.Request) {
	if !verifyRequest("POST", w, r) {
		return
	}

	var makeMoveJSON MakeMoveJSON
	err := json.NewDecoder(r.Body).Decode(&makeMoveJSON)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error":"%s"}`, err.Error()), http.StatusBadRequest)
		return
	} else if makeMoveJSON.GameID == "" || makeMoveJSON.Location == "" || makeMoveJSON.Player == "" { // TODO: refactor into own function to use on all API calls
		http.Error(w, `{"error":"'gameId', 'location' and 'player' fields required"}`, http.StatusBadRequest)
		return
	}

	gameId, convErr := strconv.ParseInt(makeMoveJSON.GameID, 10, 32) // TODO: refactor into own function.
	if convErr != nil {
		http.Error(w, `{"error":"bad gameId provided"}`, http.StatusBadRequest)
		return
	}
	game, findGameErr := gm.GetGameById(int(gameId)) // TODO: refactor into own function.
	if findGameErr != nil {
		http.Error(w, `{"error":"game does not exist"}`, http.StatusBadRequest)
		return
	}
	location, convErr := strconv.ParseInt(makeMoveJSON.Location, 10, 32)
	if convErr != nil {
		http.Error(w, `{"error":"bad location provided"}`, http.StatusBadRequest)
		return
	}

	err = gm.MakeMove(game.GetId(), tictactoe.Cell(makeMoveJSON.Player), tictactoe.Location(location))
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error":"%s"}`, err.Error()), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func verifyRequest(supportedMethod string, w http.ResponseWriter, r *http.Request) bool {
	if r.Method != supportedMethod {
		http.Error(w, fmt.Sprintf(`{"error":"Invalid request method: %s"}`, r.Method), http.StatusMethodNotAllowed)
		return false
	} else if r.Body == nil {
		http.Error(w, `{"error":"Please send a request body"}`, http.StatusBadRequest)
		return false
	}

	return true
}
