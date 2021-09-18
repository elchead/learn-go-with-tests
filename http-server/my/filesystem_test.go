package main

import (
	"io"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createTempFile(t testing.TB, initialData string) (io.ReadWriteSeeker, func()) {
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
	t.Run("get league from reader", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
	{"Name": "Cleo", "Score": 10},
	{"Name": "Chris", "Score": 33}
]
`)
		defer cleanDatabase()
		store := NewFileSystemPlayerStore(database)
		got := store.GetLeague()
		assert.ElementsMatch(t, []Player{{"Cleo", 10}, {"Chris", 33}}, got)
		got = store.GetLeague() // test idempotency
		assert.ElementsMatch(t, []Player{{"Cleo", 10}, {"Chris", 33}}, got)
	})
	t.Run("get player score", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			 			{"Name": "Cleo", "Score": 10},
			 			{"Name": "Chris", "Score": 33}
			 		]
		`)
		defer cleanDatabase()
		store := NewFileSystemPlayerStore(database)
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
		store := NewFileSystemPlayerStore(database)
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
		store := NewFileSystemPlayerStore(database)
		store.RecordWin("Adrian")
		got, _ := store.GetPlayerScore("Adrian")
		assert.Equal(t, 1, got)
	})
}
