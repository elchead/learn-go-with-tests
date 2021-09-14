package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func TestListenAndServe(t *testing.T) {
	t.Run("return Bob", func(t *testing.T) {
		req := newGetScoreRequest("Bob")
		resp := httptest.NewRecorder()
		PlayerServer(resp, req)
		assert.Equal(t, resp.Body.String(), "100")
	})
	t.Run("return Floyd", func(t *testing.T) {
		req := newGetScoreRequest("Floyd")
		resp := httptest.NewRecorder()
		PlayerServer(resp, req)
		assert.Equal(t, resp.Body.String(), "50")
	})
}
