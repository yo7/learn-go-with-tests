package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(PlayerServer)
	err := http.ListenAndServe(":5000", handler)
	if err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]

	if player == "Pepper" {
		fmt.Fprintf(w, "20")
	}

	if player == "Floyd" {
		fmt.Fprintf(w, "10")
	}
}
