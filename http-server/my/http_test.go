package main

import (
	"encoding/json"
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

func newPostRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func TestListenAndServe(t *testing.T) {
	t.Run("return Bob", func(t *testing.T) {
		req := newGetScoreRequest("Bob")
		resp := httptest.NewRecorder()
		sv := NewPlayerServer(StubStore{"Floyd": 50, "Bob": 100})
		sv.ServeHTTP(resp, req)
		assert.Equal(t, resp.Body.String(), "100")
		assert.Equal(t, resp.Code, http.StatusOK)
	})
	t.Run("return Floyd", func(t *testing.T) {
		req := newGetScoreRequest("Floyd")
		resp := httptest.NewRecorder()
		sv := NewPlayerServer(StubStore{"Floyd": 50, "Bob": 100})
		sv.ServeHTTP(resp, req)
		assert.Equal(t, resp.Body.String(), "50")
		assert.Equal(t, resp.Code, http.StatusOK)
	})
	t.Run("missing player 404", func(t *testing.T) {
		req := newGetScoreRequest("Hans")
		resp := httptest.NewRecorder()
		sv := NewPlayerServer(StubStore{"Floyd": 50, "Bob": 100})
		sv.ServeHTTP(resp, req)
		// assert.Equal(t, resp.Body.String(), "0")
		assert.Equal(t, http.StatusNotFound, resp.Code)
	})
}

func TestStorePostScore(t *testing.T) {
	store := StubStore{}
	server := NewPlayerServer(&store)
	t.Run("accept post", func(t *testing.T) {
		rq, _ := http.NewRequest(http.MethodPost, "/players/Pepper", nil)
		rp := httptest.NewRecorder()
		server.ServeHTTP(rp, rq)
		assert.Equal(t, http.StatusOK, rp.Code)
		_, ok := store["Pepper"]
		assert.Equal(t, true, ok)
	})
}

func TestLeague(t *testing.T) {
	mapa := map[string]int{"Bob": 100, "Floyd": 50}
	server := NewPlayerServer(StubStore(mapa))
	rq, _ := http.NewRequest(http.MethodGet, "/league", nil)
	t.Run("returns 200 on /league", func(t *testing.T) {
		rp := httptest.NewRecorder()
		server.ServeHTTP(rp, rq)
		assert.Equal(t, http.StatusOK, rp.Code)
	})
	t.Run("returns player list on /league", func(t *testing.T) {
		rp := httptest.NewRecorder()
		server.ServeHTTP(rp, rq)
		assert.Equal(t, "application/json", rp.Result().Header.Get("Content-Type"))

		var got []Player
		err := json.NewDecoder(rp.Body).Decode(&got)
		assert.NoError(t, err)
		assert.ElementsMatch(t, ConvertMapToPlayers(mapa), got) // use ElementsMatch (sorting) because decoding has NON-DETERMINISTIC array order!
	})
}

func TestGetEndpointName(t *testing.T) {
	path := "/league/s/"
	assert.Equal(t, "league", getEndpointName(path))

}
