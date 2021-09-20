package poker_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	"github.com/elchead/poker"
	"github.com/stretchr/testify/assert"
)

type SpyGame struct {
	Name       string
	NumPlayers int
	WasStarted bool
}

func (s *SpyGame) Start(numberOfPlayers int, to io.Writer) {
	s.NumPlayers = numberOfPlayers
	s.WasStarted = true
}

func (s *SpyGame) Finish(name string) {
	s.Name = name
}

func assertMsgSentToUser(t testing.TB, stdout *bytes.Buffer) {
	t.Helper()
	got := stdout.String()
	want := poker.PlayerPrompt
	assert.Equal(t, got, want)

}

func assertFinishedWith(t testing.TB, want int, game *SpyGame) {
	t.Helper()
	assert.Equal(t, want, game.NumPlayers)
}

func TestCLI(t *testing.T) {
	t.Run("start game with 7 players and let 'Chris' win", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("7\nChris wins\n")
		dummyGame := &SpyGame{}
		cli := poker.NewCLI(dummyGame, in, stdout)
		cli.PlayPoker()
		assertMsgSentToUser(t, stdout)
		assertFinishedWith(t, 7, dummyGame)
		assert.Equal(t, "Chris", dummyGame.Name)
	})
	t.Run("start game with 3 players and let 'Leo' win", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("3\nLeo wins\n")
		dummyGame := &SpyGame{}
		cli := poker.NewCLI(dummyGame, in, stdout)
		cli.PlayPoker()
		assertMsgSentToUser(t, stdout)
		assertFinishedWith(t, 3, dummyGame)
		assert.Equal(t, "Leo", dummyGame.Name)
	})
	t.Run("no game when no number of players provided", func(t *testing.T) {
		dummyStdout := &bytes.Buffer{}
		in := strings.NewReader("Hi\n")
		gameSpy := &SpyGame{WasStarted: false}
		cli := poker.NewCLI(gameSpy, in, dummyStdout)
		err := cli.PlayPoker()
		assert.Equal(t, false, gameSpy.WasStarted)
		assert.EqualError(t, err, poker.BadPlayerInputErrMsg)

	})
	t.Run("throw error when winner input is invalid", func(t *testing.T) {
		dummyStdout := &bytes.Buffer{}
		in := strings.NewReader("5\nLloyd is a killer")
		gameSpy := &SpyGame{WasStarted: false}
		cli := poker.NewCLI(gameSpy, in, dummyStdout)
		err := cli.PlayPoker()
		assert.Error(t, err)
	})

}
