package poker

import "time"

type Game struct {
	store   PlayerStore
	alerter Alerter
}

func NewGame(store PlayerStore, blind Alerter) *Game {
	return &Game{store, blind}
}

func (g *Game) Start(numberPlayers int) {
	blindIncrement := time.Duration(5+numberPlayers) * time.Minute
	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second
	for _, blind := range blinds {
		g.alerter.ScheduleAlertAt(blindTime, blind)
		blindTime = blindTime + blindIncrement
	}
}

func (g *Game) Finish(name string) {
	g.store.RecordWin(name)
}
