package poker_test

import (
	"fmt"
	"io"
	"os"
	"testing"
	"time"

	"github.com/elchead/poker"
	"github.com/stretchr/testify/assert"
)

func TestGame(t *testing.T) {
	store := &poker.StubPlayerStore{}
	blindAlerter := &SpyBlindAlerter{}
	game := poker.NewTexasHoldem(store, blindAlerter)
	players := 7
	game.Start(players, os.Stdout)
	cases := []alert{
		{0 * time.Second, 100},
		{12 * time.Minute, 200},
		{24 * time.Minute, 300},
		{36 * time.Minute, 400},
	}

	for i, want := range cases {
		t.Run(fmt.Sprint(want), func(t *testing.T) {

			if len(blindAlerter.alerts) <= i {
				t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.alerts)
			}

			got := blindAlerter.alerts[i]
			assertScheduledAlert(t, got, want)
		})
	}

	t.Run("record win", func(t *testing.T) {
		game.Finish("Chris")
		fmt.Println(store.WinCalls)
		assert.Equal(t, 1, len(store.WinCalls))
	})

}

func assertScheduledAlert(t testing.TB, got, want alert) {
	t.Helper()
	if got != want {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

type alert struct {
	time   time.Duration
	amount int
}

func (s alert) String() string {
	return fmt.Sprintf("%d chips at %v", s.amount, s.time)
}

type SpyBlindAlerter struct {
	alerts []alert
}

func (s *SpyBlindAlerter) ScheduleAlertAt(at time.Duration, amount int, to io.Writer) {
	s.alerts = append(s.alerts, alert{at, amount})
}
