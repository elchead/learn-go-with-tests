package main

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) (int, bool)
	GetPlayers() []string
	PostPlayerWin(name string) error
}

type PlayerServer struct {
	store PlayerStore
}

func (s PlayerServer) showScore(w http.ResponseWriter, player string) {
	if val, ok := s.store.GetPlayerScore(player); ok {
		fmt.Fprint(w, val)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, val)
	}
}

func (s PlayerServer) postPlayer(w http.ResponseWriter, player string) {
	s.store.PostPlayerWin(player)
	fmt.Fprintf(w, "posted")
}

func getEndpointName(path string) string {
	return strings.Split(path, "/")[1]
}

func (s PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	endpoint := getEndpointName(r.URL.Path)
	if endpoint == "league" {
		w.WriteHeader(http.StatusOK)
		players := s.store.GetPlayers()
		fmt.Fprintf(w, "%s, %s", players[0], players[1])
		return
	}
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	switch r.Method {
	case http.MethodGet:
		s.showScore(w, player)
	case http.MethodPost:
		s.postPlayer(w, player)
	}
}

type StubStore map[string]int

func (s StubStore) GetPlayerScore(player string) (val int, ok bool) {
	val, ok = s[player] // map returns 0 if not found..
	return
}

func (s StubStore) PostPlayerWin(name string) error {
	s[name] += 1
	return nil
}

func (s StubStore) GetPlayers() []string {
	keys := make([]string, 0, len(s))
	for k := range s {
		keys = append(keys, k)
	}
	return keys
}
