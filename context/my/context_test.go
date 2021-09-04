package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestStore struct {
	data string
}

func (s TestStore) Fetch() string {
	return s.data
}
func TestServer(t *testing.T) {
	data := "hello"
	store := TestStore{data: data}
	svr := Server(store)
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()
	svr.ServeHTTP(response, request)
	assert.Equal(t, data, response.Body.String())
}
