package main

import (
	"fmt"
	"net/http"
	"strings"
)

type playerStore map[string]int

func PlayerServer(w http.ResponseWriter, request *http.Request) {
	player := strings.TrimPrefix(request.URL.Path, "/players/")
	if player == "Floyd" {
		fmt.Fprint(w, 50)
		return
	}
	if player == "Bob" {
		fmt.Fprint(w, 100)
	}
}

func createServer(store playerStore) string {
	// http.HandleFunc("/players/Bob", playerHandler)
	http.ListenAndServe(":8080", nil)
	return "http:localhost:8080"
}
