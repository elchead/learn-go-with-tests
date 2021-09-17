package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("get league from reader", func(t *testing.T) {
		database := strings.NewReader(`[
			{"Name": "Cleo", "Score": 10},
			{"Name": "Chris", "Score": 33}
		]`)
		store := FileSystemPlayerStore{database}
		got := store.GetLeague()
		assert.ElementsMatch(t, []Player{{"Cleo", 10}, {"Chris", 33}}, got)
		got = store.GetLeague() // test idempotency
		assert.ElementsMatch(t, []Player{{"Cleo", 10}, {"Chris", 33}}, got)
	})
	t.Run("get player score", func(t *testing.T) {
		database := strings.NewReader(`[
			{"Name": "Cleo", "Score": 10},
			{"Name": "Chris", "Score": 33}
		]`)
		store := FileSystemPlayerStore{database}
		got, _ := store.GetPlayerScore("Cleo")
		assert.Equal(t, 10, got)
	})
}
