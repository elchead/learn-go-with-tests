package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListenAndServe(t *testing.T) {
	t.Run("return Bob", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/players/Bob", nil)
		resp := httptest.NewRecorder()
		PlayerServer(resp, req)
		assert.Equal(t, resp.Body.String(), "100")
	})
	t.Run("return Floyd", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/players/Floyd", nil)
		resp := httptest.NewRecorder()
		PlayerServer(resp, req)
		assert.Equal(t, resp.Body.String(), "50")
	})
}
