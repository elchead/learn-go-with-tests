package main

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) (int, bool)
}

type PlayerServer struct {
	store PlayerStore
}

func (s PlayerServer) showScore(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	if val, ok := s.store.GetPlayerScore(player); ok {
		fmt.Fprint(w, val)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, val)
	}
}

func (s PlayerServer) postPlayer(w http.ResponseWriter) {
	fmt.Fprintf(w, "posted")
}

func (s PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.showScore(w, r)
	case http.MethodPost:
		s.postPlayer(w)
	}
}

type StubStore map[string]int

func (s StubStore) GetPlayerScore(player string) (val int, ok bool) {
	val, ok = s[player] // map returns 0 if not found..
	return
}
