package poker

import (
	"io"
	"time"
)

type Gamer interface {
	Start(numberPlayers int, alertsDestination io.Writer)
	Finish(name string)
}
type TexasHoldem struct {
	store   PlayerStore
	alerter BlindAlerter
}

func NewTexasHoldem(store PlayerStore, blind BlindAlerter) *TexasHoldem {
	return &TexasHoldem{store, blind}
}

func (g *TexasHoldem) Start(numberPlayers int, alertsDestination io.Writer) {
	blindIncrement := time.Duration(5+numberPlayers) * time.Minute
	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second
	for _, blind := range blinds {
		g.alerter.ScheduleAlertAt(blindTime, blind, alertsDestination)
		blindTime = blindTime + blindIncrement
	}
}

func (g *TexasHoldem) Finish(name string) {
	g.store.RecordWin(name)
}
