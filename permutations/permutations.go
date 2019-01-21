package main

import (
	"fmt"
	"strings"

	goset "github.com/deckarep/golang-set"
)

// TODO: more questions:
// how many times does cross or noughts win? what about draws?
// what's the distribution of number of turns taken in each permutation?

func main() {
	uniquePermutations := goset.NewSet()

	permutationQueue := make([]string, 0)
	permutationQueue = append(permutationQueue, "---------")
	uniquePermutations.Add("---------")

	for len(permutationQueue) != 0 {
		fmt.Printf("Permutations left to check: %d\n", len(permutationQueue))

		board := permutationQueue[0]
		permutationQueue = permutationQueue[1:]

		nextPerms := getPossibleNextPermutations(board)
		for _, perm := range nextPerms {
			uniquePermutations.Add(board) // don't worry about duplicates; set ensures uniqueness

			if !isGameOver(perm) {
				permutationQueue = append(permutationQueue, perm)
			}
		}
	}

	fmt.Println(uniquePermutations)
	fmt.Printf("Possible permutations = %d\n", len(uniquePermutations.ToSlice()))
}

func getPossibleNextPermutations(board string) []string {
	nextPerms := make([]string, 0)
	for i, v := range board {
		if string(v) == "-" {
			nextMove := getNextMoveFromBoard(board)
			nextPerms = append(nextPerms, board[:i]+nextMove+board[i+1:])
		}
	}
	return nextPerms
}

func getNextMoveFromBoard(b string) string {
	emptyCount := strings.Count(b, "-")
	if emptyCount%2 == 1 {
		return "X"
	}
	return "O"
}

var empty byte = "-"[0]

func isGameOver(b string) bool {
	return (b[TopLeft] != empty && b[TopLeft] == b[Top] && b[Top] == b[TopRight]) || // top row
		(b[TopLeft] != empty && b[TopLeft] == b[Left] && b[Left] == b[BottomLeft]) || // left column
		(b[Left] != empty && b[Left] == b[Centre] && b[Centre] == b[Right]) || // middle row
		(b[Top] != empty && b[Top] == b[Centre] && b[Centre] == b[Bottom]) || // middle column
		(b[TopLeft] != empty && b[TopLeft] == b[Centre] && b[Centre] == b[BottomRight]) || // back-slash diagonal ( \ )
		(b[BottomLeft] != empty && b[BottomLeft] == b[Centre] && b[Centre] == b[TopRight]) || // forward-slash diagonal ( / )
		(b[BottomLeft] != empty && b[BottomLeft] == b[Bottom] && b[Bottom] == b[BottomRight]) || // bottom row
		(b[TopRight] != empty && b[TopRight] == b[Right] && b[Right] == b[BottomRight]) || // right column
		strings.Count(b, "-") == 9
}

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

func nextPlayer(currentPlayer *string) {
	if *currentPlayer == "X" {
		*currentPlayer = "O"
	} else {
		*currentPlayer = "X"
	}
}
