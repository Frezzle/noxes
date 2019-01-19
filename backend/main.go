package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", homePage)

	address := "localhost:9876"
	log.Printf("Listening on %s", address)
	log.Fatal(http.ListenAndServe(address, nil))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi!")
}
