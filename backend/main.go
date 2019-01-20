package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/game", getGameHandler)
	// http.HandleFunc("/games", getAllGamesHandler)
	// http.HandleFunc("/game/create", createGameHandler)
	// http.HandleFunc("/game/join", ...)
	// http.HandleFunc("/game/move", ...)
	// http.HandleFunc("/game/resign", ...)

	address := "localhost:9876"
	log.Printf("Listening on %s", address)
	log.Fatal(http.ListenAndServe(address, nil))
}

type Game struct {
	// ID       string `json:"id"`
	Board    string `json:"board"`
	NextTurn string `json:"nextTurn"`
	GameOver string `json:"gameOver"`
	Winner   string `json:"winner"`
}

var game Game = Game{Board: "---------", NextTurn: "X", GameOver: "false", Winner: "-"}

func getGameHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, fmt.Sprintf("Invalid request method: %s", r.Method), http.StatusMethodNotAllowed)
	}

	json.NewEncoder(w).Encode(game)
}

// type Games []Game

// var nextGameId int = 0

// var games Games = make([]Game, 0, 100)

// func getAllGamesHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != "GET" {
// 		http.Error(w, fmt.Sprintf("Invalid request method: %s", r.Method), http.StatusMethodNotAllowed)
// 	}

// 	json.NewEncoder(w).Encode(games)
// }

// type JSONError struct {
// 	Error string `json:"error"`
// }

// func createGameHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != "POST" {
// 		http.Error(w, fmt.Sprintf("Invalid request method: %s", r.Method), http.StatusMethodNotAllowed)
// 		return
// 	}

// 	createNewGame()

// 	// if err != nil {
// 	// 	json.NewEncoder(w).Encode(JSONError{Error: "not implemented"})
// 	// 	return
// 	// }

// 	w.WriteHeader(http.StatusOK)
// }

// func createNewGame() {
// 	id := nextGameId
// 	nextGameId++
// 	games = append(games, Game{ID: strconv.Itoa(id), Board: "---------", NextTurn: "X", GameOver: "false", Winner: "-"})
// }
