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
	// var dummyBlindAlerter = &SpyBlindAlerter{}
	var dummyPlayerStore = &poker.StubPlayerStore{}
	// var dummyStdIn = &bytes.Buffer{}
	var dummyStdOut = &bytes.Buffer{}

	in := strings.NewReader("5\nChris wins\n")
	store := &poker.StubPlayerStore{}
	dummyAlerter := &SpyBlindAlerter{}
	cli := poker.NewCLI(store, in, dummyStdOut, dummyAlerter)
	cli.PlayPoker()
	assert.Equal(t, 1, len(store.WinCalls))
	assert.Equal(t, "Chris", store.WinCalls[0])

	t.Run("it schedules printing of blind values", func(t *testing.T) {
		in := strings.NewReader("5\nChris wins\n")
		playerStore := &poker.StubPlayerStore{}
		blindAlerter := &SpyBlindAlerter{}
		cli := poker.NewCLI(playerStore, in, dummyStdOut, blindAlerter)
		assert.Equal(t, 0, len(blindAlerter.alerts))
		cli.PlayPoker()
		cases := []alert{
			{0 * time.Second, 100},
			{10 * time.Minute, 200},
			{20 * time.Minute, 300},
			{30 * time.Minute, 400},
			{40 * time.Minute, 500},
			{50 * time.Minute, 600},
			{60 * time.Minute, 800},
			{70 * time.Minute, 1000},
			{80 * time.Minute, 2000},
			{90 * time.Minute, 4000},
			{100 * time.Minute, 8000},
		}
		for i, want := range cases {
			t.Run(fmt.Sprintf("Value %d scheduled for %v", want.amount, want.time), func(t *testing.T) {
				assert.Equal(t, want.amount, blindAlerter.alerts[i].amount)
				assert.Equal(t, want.time, blindAlerter.alerts[i].time)
			})
		}
	})
	t.Run("it prompts the user to enter the number of players", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("7\n")
		blindAlerter := &SpyBlindAlerter{}
		cli := poker.NewCLI(dummyPlayerStore, in, stdout, blindAlerter)
		cli.PlayPoker()

		got := stdout.String()
		want := poker.PlayerPrompt

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

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

func (s *SpyBlindAlerter) ScheduleAlertAt(at time.Duration, amount int) {
	s.alerts = append(s.alerts, alert{at, amount})
}
