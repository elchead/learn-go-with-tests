package main

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createTempFile(t testing.TB, initialData string) (ReadWriteTruncate, func()) {
	t.Helper()

	tmpfile, err := ioutil.TempFile("", "db")

	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	tmpfile.Write([]byte(initialData))

	removeFile := func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}

	return tmpfile, removeFile
}

func TestFileSystemStore(t *testing.T) {
	t.Run("league is sorted", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
	{"Name": "Cleo", "Score": 10},
	{"Name": "Chris", "Score": 33}
]
`)
		defer cleanDatabase()
		store, err := NewFileSystemPlayerStore(database)
		assert.NoError(t, err)
		got := store.GetLeague()
		assert.Equal(t, League{{"Chris", 33}, {"Cleo", 10}}, got)
		// got = store.GetLeague() // test idempotency
		// assert.ElementsMatch(t, []Player{{"Cleo", 10}, {"Chris", 33}}, got)
	})
	t.Run("get player score", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			 			{"Name": "Cleo", "Score": 10},
			 			{"Name": "Chris", "Score": 33}
			 		]
		`)
		defer cleanDatabase()
		store, err := NewFileSystemPlayerStore(database)
		assert.NoError(t, err)
		got, _ := store.GetPlayerScore("Cleo")
		assert.Equal(t, 10, got)

	})
	t.Run("store wins for existing players", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Cleo", "Score": 10},
			{"Name": "Chris", "Score": 33}
		]
`)
		defer cleanDatabase()
		store, err := NewFileSystemPlayerStore(database)
		assert.NoError(t, err)
		store.RecordWin("Chris")
		got, _ := store.GetPlayerScore("Chris")
		assert.Equal(t, 34, got)
	})
	t.Run("store win for new player", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Cleo", "Score": 10},
			{"Name": "Chris", "Score": 33}
		]
`)
		defer cleanDatabase()
		store, err := NewFileSystemPlayerStore(database)
		assert.NoError(t, err)
		store.RecordWin("Adrian")
		got, ok := store.GetPlayerScore("Adrian")
		assert.Equal(t, true, ok)
		assert.Equal(t, 1, got)
	})
	// t.Run("works with an empty file", func(t *testing.T) {
	// 	database, cleanDatabase := createTempFile(t, "")
	// 	defer cleanDatabase()

	// 	_, err := NewFileSystemPlayerStore(database)

	// 	assert.NoError(t, err)
	// })
}
