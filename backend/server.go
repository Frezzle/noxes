package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"noxes/backend/game_manager"
	"strconv"
)

func main() {
	http.HandleFunc("/game", getGameHandler)
	// http.HandleFunc("/games", getAllGamesHandler)
	http.HandleFunc("/game/create", createGameHandler)
	// http.HandleFunc("/game/join", ...)
	// http.HandleFunc("/game/move", ...)
	// http.HandleFunc("/game/resign", ...)

	address := "localhost:9876"
	log.Printf("Listening on %s", address)
	log.Fatal(http.ListenAndServe(address, nil))
}

type GameJSON struct {
	ID       string `json:"id"`
	Board    string `json:"board"`
	NextTurn string `json:"nextTurn"`
	GameOver string `json:"gameOver"`
	Winner   string `json:"winner"`
}

type IdJSON struct {
	ID string `json:"id"`
}

var gm = game_manager.NewGameManager()

func getGameHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, fmt.Sprintf(`{"error":"Invalid request method: %s"}`, r.Method), http.StatusMethodNotAllowed)
		return
	} else if r.Body == nil {
		http.Error(w, `{"error":"Please send a request body"}`, http.StatusBadRequest)
		return
	}

	var gameIdJSON IdJSON // only care about id from this object
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

// func getAllGamesHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != "GET" {
// 		http.Error(w, fmt.Sprintf("Invalid request method: %s", r.Method), http.StatusMethodNotAllowed)
// 	}

// 	json.NewEncoder(w).Encode(games)
// }

func createGameHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, fmt.Sprintf(`{"error":"Invalid request method: %s"}`, r.Method), http.StatusMethodNotAllowed)
		return
	} else if r.Body == nil {
		http.Error(w, `{"error":"Please send a request body"}`, http.StatusBadRequest)
		return
	}

	gameId := gm.CreateNewGame()

	var result = IdJSON{ID: strconv.FormatInt(int64(gameId), 10)}

	json.NewEncoder(w).Encode(result)
}