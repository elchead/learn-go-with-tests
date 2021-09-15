package main

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := StubStore{}
	server := PlayerServer{&store}
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPostRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostRequest(player))

	resp := httptest.NewRecorder()
	server.ServeHTTP(resp, newGetScoreRequest(player))
	assert.Equal(t, "3", resp.Body.String())
}
