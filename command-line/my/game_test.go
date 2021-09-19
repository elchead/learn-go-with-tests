package poker_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/elchead/poker"
	"github.com/stretchr/testify/assert"
)

func TestGame(t *testing.T) {
	store := &poker.StubPlayerStore{}
	blindAlerter := &SpyBlindAlerter{}
	players := 7
	game := poker.NewGame(store, blindAlerter)
	game.Start(players)
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
