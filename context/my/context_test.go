package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type SpyStore struct {
	data      string
	cancelled bool
}

func (s SpyStore) Fetch() string {
	time.Sleep(2 * time.Millisecond)
	return s.data
}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}
func TestServer(t *testing.T) {
	t.Run("server returns data", func(t *testing.T) {
		data := "hello"
		store := SpyStore{data: data}
		svr := Server(&store)
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()
		svr.ServeHTTP(response, request)
		assert.Equal(t, data, response.Body.String())
		assert.Equal(t, false, store.cancelled)
	})
	t.Run("store cancels work if request is canceled", func(t *testing.T) {
		data := "hello"
		store := SpyStore{data: data, cancelled: false}
		svr := Server(&store)
		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		response := httptest.NewRecorder()
		svr.ServeHTTP(response, request)

		assert.Equal(t, true, store.cancelled)
	})
}
