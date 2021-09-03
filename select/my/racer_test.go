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
	t.Run("should return faster url", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(1 * time.Millisecond)
		slowURL := slowServer.URL
		fastURL := fastServer.URL
		want := fastURL
		got, err := Racer(slowURL, fastURL, 1*time.Second)

		assert.Equal(t, want, got)
		assert.NoError(t, err)
		defer func() {
			slowServer.Close()
			fastServer.Close()
		}()
	})
	t.Run("should timeout", func(t *testing.T) {
		slowServer := makeDelayedServer(2 * time.Second)
		fastServer := makeDelayedServer(1 * time.Second)
		slowURL := slowServer.URL
		fastURL := fastServer.URL
		_, err := Racer(slowURL, fastURL, 20*time.Millisecond)

		assert.Error(t, err)
		defer func() {
			slowServer.Close()
			fastServer.Close()
		}()
	})
}
