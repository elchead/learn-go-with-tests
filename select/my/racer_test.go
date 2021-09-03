package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func makeDelayedServer(duration time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(duration)
		w.WriteHeader(http.StatusOK)
	}))
}

func TestRacer(t *testing.T) {

	slowServer := makeDelayedServer(20 * time.Millisecond)
	fastServer := makeDelayedServer(1 * time.Millisecond)
	slowURL := slowServer.URL
	fastURL := fastServer.URL
	want := fastURL
	got := Racer(slowURL, fastURL)

	assert.Equal(t, want, got)
	defer func() {
		slowServer.Close()
		fastServer.Close()
	}()
}
