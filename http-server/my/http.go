package main

import (
	"fmt"
	"net/http"
	"strings"
)

type playerStore map[string]int

func GetPlayerScore(player string) int {
	if player == "Floyd" {
		return 50
	}
	if player == "Bob" {
		return 100
	}
	return 0
}

func PlayerServer(w http.ResponseWriter, request *http.Request) {
	player := strings.TrimPrefix(request.URL.Path, "/players/")
	fmt.Fprint(w, GetPlayerScore(player))
}

func createServer(store playerStore) string {
	// http.HandleFunc("/players/Bob", playerHandler)
	http.ListenAndServe(":8080", nil)
	return "http:localhost:8080"
}
