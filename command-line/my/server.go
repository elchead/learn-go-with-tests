package poker

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"

	"github.com/gorilla/websocket"
)

// PlayerStore stores score information about players.
type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
	GetLeague() League
}

// Player stores a name with a number of wins.
type Player struct {
	Name string
	Wins int
}

// PlayerServer is a HTTP interface for player information.
type PlayerServer struct {
	store PlayerStore
	http.Handler
	upgrader websocket.Upgrader
	template *template.Template
	game     Gamer
}

const jsonContentType = "application/json"

// NewPlayerServer creates a PlayerServer with routing configured.
func NewPlayerServer(store PlayerStore, game Gamer) *PlayerServer {
	p := new(PlayerServer)

	tmpl, err := template.ParseFiles("game.html")
	if err != nil {
		log.Fatal("Failed to load template")
	}
	p.template = tmpl
	p.store = store
	p.upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	p.game = game

	router := http.NewServeMux()
	router.Handle("/ws", http.HandlerFunc(p.webSocket))
	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.playersHandler))
	router.Handle("/game", http.HandlerFunc(p.gameHandler))
	p.Handler = router

	return p
}

func (p *PlayerServer) webSocket(w http.ResponseWriter, r *http.Request) {

	conn, _ := p.upgrader.Upgrade(w, r, nil)
	// _, winnerMsg, _ := conn.ReadMessage()
	_, numberOfPlayersMsg, _ := conn.ReadMessage()
	numberOfPlayers, _ := strconv.Atoi(string(numberOfPlayersMsg))
	p.game.Start(numberOfPlayers, w)

	_, winner, _ := conn.ReadMessage()
	p.game.Finish(string(winner))
	p.store.RecordWin(string(winner))
}
func (p *PlayerServer) gameHandler(w http.ResponseWriter, r *http.Request) {
	p.template.Execute(w, nil)
}

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", jsonContentType)
	json.NewEncoder(w).Encode(p.store.GetLeague())
}

func (p *PlayerServer) playersHandler(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	switch r.Method {
	case http.MethodPost:
		p.processWin(w, player)
	case http.MethodGet:
		p.showScore(w, player)
	}
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}
