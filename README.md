# noxes - Noughts and Crosses 

RESTful Web API for managing multiple concurrent TicTacToe games.

## Architecture

The project is split into 3 parts:

1. [TicTacToe library](https://github.com/Frezzle/noxes/blob/master/tictactoe/tictactoe.go): Defines the logic for a single game.
2. [Game Manager](https://github.com/Frezzle/noxes/blob/master/game_manager/game_manager.go): Manages multiple games at a time.
3. [Web API](https://github.com/Frezzle/noxes/blob/master/server.go): Defines the API and launches the web server, handling all HTTP requests and responses.

## Missing features / future improvements

- API testing.
- Authentication: Anyone can make a request to play as either nought or cross on any game. Play nice (for now)!
- Persistent storage: There is no database to store the game data if the server goes down. All is within memory, and resets when server does. Careful!
- GUI: No flashy frontend yet. Please play using your favourite API tool.

## Running locally

1. [Install Go](https://golang.org/doc/install).
2. Clone and run:
```
git clone https://github.com/Frezzle/noxes
go run noxes/server.go
```
3. Play using the [API](#API). The web server listens on http://localhost:9876 by default.

Unit tests can be run with `go test noxes/...`

## API

### Errors

All endpoints, upon some failure, return an error response:

```
{
    "error": "some error message"
}
```

### POST /game/create

Creates a new game instance.

Response:

```
{
    "id": "0"
}
```

- `id` - Uniquely identifies the new game instance.

### GET /game

Gets the state of an existing game, given a game `id`. Unfortunately, the `/game` endpoint currently requires a GET method with a body, so you must use tools which allow providing a body in a GET request (i.e. not [Swagger Inspector](https://inspector.swagger.io/builder)).

Request:

```
{
    "id": "0"
}
```

- `id` - Uniquely identifies an existing game instance.

Response:

```
{
    "id": "0",
    "board": "---------",
    "nextTurn": "X",
    "gameOver": "false",
    "winner": "-"
}
```

- `id` - Uniquely identifies the game instance.
- `board` - 9-character string representing the cells of the board, with `X`, `O`, and `-` being Cross, Nought and empty respectively. Each string index is mapped to a position on the board:
    ```
      0 | 1 | 2
     -----------
      3 | 4 | 5
     -----------
      6 | 7 | 8
    ```
- `nextTurn` - who is next to play; `X` or `O`. `X` always starts.
- `gameOver` - `true` or `false`, signifying if a game is over. The game ends when a player achieves 3-in-a-row or a draw occurs.
- `winner` - who won the game. Starts with noone as the winner (`-`), becoming `X` or `O` if someone wins; stays as `-` if game ends as a draw.

### GET /games

Gets the state of all existing games.

Response is an array of zero or more games:

```
[
    {
        "id": "0",
        "board": "---------",
        "nextTurn": "X",
        "gameOver": "false",
        "winner": "-"
    },
    {
        "id": "1",
        "board": "---X-----",
        "nextTurn": "O",
        "gameOver": "false",
        "winner": "-"
    }
]
```

See [GET /game](#GET-/games) for the meaning of each game's attributes shown above.

### POST /game/move

Play a move on an existing game. 

Request:

```
{
    "gameId": "0",
    "player": "X",
    "location": "4"
}
```

- `gameId` - Identifies the game instance where you'd like to play the move.
- `player` - The side you are playing as; either `X` or `O`.
- `location` - The cell which you'd like to place your move. The board cells are mapped from 0 to 8:
    ```
      0 | 1 | 2
     -----------
      3 | 4 | 5
     -----------
      6 | 7 | 8
    ```

If your move was successful then the response is a `200 OK` status with no body.
