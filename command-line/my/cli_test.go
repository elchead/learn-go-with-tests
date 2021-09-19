package poker_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/elchead/poker"
	"github.com/stretchr/testify/assert"
)

func TestCLI(t *testing.T) {
	in := strings.NewReader("Chris wins\n")
	store := &poker.StubPlayerStore{}
	dummyAlerter := &SpyBlindAlerter{}
	cli := poker.NewCLI(store, in, dummyAlerter)
	cli.PlayPoker()
	assert.Equal(t, 1, len(store.WinCalls))
	assert.Equal(t, "Chris", store.WinCalls[0])

}

func TestTime(t *testing.T) {
	t.Run("it schedules printing of blind values", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		playerStore := &poker.StubPlayerStore{}
		blindAlerter := &SpyBlindAlerter{}
		cli := poker.NewCLI(playerStore, in, blindAlerter)
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
