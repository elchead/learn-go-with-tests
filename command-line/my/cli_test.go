package poker_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/elchead/poker"
	"github.com/stretchr/testify/assert"
)

func TestCLI(t *testing.T) {
	var dummyStdOut = &bytes.Buffer{}

	in := strings.NewReader("5\nChris wins\n")
	store := &poker.StubPlayerStore{}
	dummyAlerter := &SpyBlindAlerter{}
	game := poker.NewGame(store, dummyAlerter)
	cli := poker.NewCLI(game, in, dummyStdOut)
	cli.PlayPoker()
	assert.Equal(t, 1, len(store.WinCalls))
	assert.Equal(t, "Chris", store.WinCalls[0])

	// t.Run("it prompts the user to enter the number of players", func(t *testing.T) {
	// 	stdout := &bytes.Buffer{}
	// 	in := strings.NewReader("7\n")
	// 	blindAlerter := &SpyBlindAlerter{}
	// 	cli := poker.NewCLI(dummyPlayerStore, in, stdout, blindAlerter)
	// 	cli.PlayPoker()

	// 	got := stdout.String()
	// 	want := poker.PlayerPrompt

	// 	if got != want {
	// 		t.Errorf("got %q, want %q", got, want)
	// 	}
	// })

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

func (s *SpyBlindAlerter) ScheduleAlertAt(at time.Duration, amount int) {
	s.alerts = append(s.alerts, alert{at, amount})
}
