package main

import (
	"context"
	"errors"
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

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)
	go func() {
		var result string
		for _, c := range s.data {
			select {
			case <-ctx.Done():
				s.Cancel()
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		data <- result
	}()
	select {
	case <-ctx.Done():
		s.Cancel()
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}

type SpyResponseWriter struct {
	written bool
}

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("not implemented")
}

func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}
func TestServer(t *testing.T) {
	t.Run("server returns data", func(t *testing.T) {
		data := "hello"
		store := &SpyStore{data: data}
		svr := Server(store)
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()
		svr.ServeHTTP(response, request)
		assert.Equal(t, data, response.Body.String())
		assert.Equal(t, false, store.cancelled)
	})
	t.Run("store cancels work if request is canceled", func(t *testing.T) {
		data := "hello"
		store := &SpyStore{data: data, cancelled: false}
		svr := Server(store)
		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		response := httptest.NewRecorder()
		svr.ServeHTTP(response, request)

		assert.Equal(t, true, store.cancelled)
	})
	t.Run("if cancelled, no response is written", func(t *testing.T) {
		data := "hello"
		store := SpyStore{data: data, cancelled: false}
		svr := Server(&store)
		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		response := &SpyResponseWriter{} //httptest.NewRecorder()
		svr.ServeHTTP(response, request)

		assert.Equal(t, false, response.written)
	})
}
