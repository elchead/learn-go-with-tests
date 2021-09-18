package main

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	db, cleanDb := createTempFile(t, "")
	defer cleanDb()
	store := NewFileSystemPlayerStore(db)
	server := NewPlayerServer(store)
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPostRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostRequest(player))

	t.Run("get score", func(t *testing.T) {
		resp := httptest.NewRecorder()
		server.ServeHTTP(resp, newGetScoreRequest(player))
		assert.Equal(t, "3", resp.Body.String())
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetLeagueRequest())
		var got []Player
		err := json.NewDecoder(response.Body).Decode(&got)
		assert.NoError(t, err)
		assert.ElementsMatch(t, []Player{{player, 3}}, got)
	})
}
