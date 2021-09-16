package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Player struct {
	Name  string
	Score int
}

type PlayerStore interface {
	GetPlayerScore(name string) (int, bool)
	GetLeague() []Player
	PostPlayerWin(name string) error
}

type PlayerServer struct {
	store PlayerStore
	http.Handler
}

func NewPlayerServer(store PlayerStore) *PlayerServer {
	s := &PlayerServer{store: store}
	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(s.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(s.playerHandler))
	s.Handler = router
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
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(s.store.GetLeague()); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
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

type StubStore map[string]int

func (s StubStore) GetPlayerScore(player string) (val int, ok bool) {
	val, ok = s[player] // map returns 0 if not found..
	return
}

func (s StubStore) PostPlayerWin(name string) error {
	s[name] += 1
	return nil
}

func ConvertMapToPlayers(m map[string]int) (players []Player) {
	keys := make([]Player, 0, len(m))
	for k, v := range m {
		keys = append(keys, Player{k, v})
	}
	return keys
}

func (s StubStore) GetLeague() []Player {
	return ConvertMapToPlayers(s)
}
