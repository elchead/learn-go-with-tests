package main

import (
	"log"
	"net/http"
)

func main() {
	handler := NewPlayerServer(StubStore{"Floyd": 50, "Bob": 100})
	log.Fatal(http.ListenAndServe(":5000", handler))
}
