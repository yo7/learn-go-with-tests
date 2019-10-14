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

	fmt.Fprintf(w, GetPlayerScore(player))
}

func GetPlayerScore(p string) string {
	if p == "Pepper" {
		return "20"
	}

	if p == "Floyd" {
		return "10"
	}

	return ""
}
