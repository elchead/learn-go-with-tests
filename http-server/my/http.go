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
	store  PlayerStore
	router *http.ServeMux
}

func NewPlayerServer(store PlayerStore) *PlayerServer {
	s := &PlayerServer{store: store, router: http.NewServeMux()}
	s.router.Handle("/league", http.HandlerFunc(s.leagueHandler))
	s.router.Handle("/players/", http.HandlerFunc(s.playerHandler))
	return s
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

func (s PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	players := s.store.GetPlayers()
	fmt.Fprintf(w, "%s, %s", players[0], players[1])
}

func (s PlayerServer) playerHandler(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	switch r.Method {
	case http.MethodGet:
		s.showScore(w, player)
	case http.MethodPost:
		s.postPlayer(w, player)
	}
}

func (s PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
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
