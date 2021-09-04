package main

import (
	"fmt"
	"net/http"
)

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := make(chan string, 1)
		go func() {
			data <- store.Fetch()
		}()
		select {
		case <-r.Context().Done():
			store.Cancel()
		case d := <-data:
			fmt.Fprint(w, d)
		}
	}
}

type Store interface {
	Fetch() string
	Cancel()
}
