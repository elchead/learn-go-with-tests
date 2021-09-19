package poker_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/elchead/poker"
	"github.com/stretchr/testify/assert"
)

type SpyGame struct {
	Name       string
	NumPlayers int
}

func (s *SpyGame) Start(numberOfPlayers int) {
	s.NumPlayers = numberOfPlayers
}

func (s *SpyGame) Finish(name string) {
	s.Name = name
}

func TestCLI(t *testing.T) {
	var dummyStdOut = &bytes.Buffer{}

	in := strings.NewReader("5\nChris wins\n")
	game := &SpyGame{}
	cli := poker.NewCLI(game, in, dummyStdOut)
	cli.PlayPoker()
	assert.Equal(t, 5, game.NumPlayers)
	assert.Equal(t, "Chris", game.Name)

	t.Run("it prompts the user to enter the number of players", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("7\n")
		dummyGame := &SpyGame{}
		cli := poker.NewCLI(dummyGame, in, stdout)
		cli.PlayPoker()

		got := stdout.String()
		want := poker.PlayerPrompt

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

}
