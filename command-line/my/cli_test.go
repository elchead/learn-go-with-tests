package poker_test

import (
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
		assert.Equal(t, 1, len(blindAlerter.alerts))
	})
}

type SpyBlindAlerter struct {
	alerts []int
}

func (s *SpyBlindAlerter) ScheduleAlertAt(at time.Duration, amount int) {
	s.alerts = append(s.alerts, amount)
}
