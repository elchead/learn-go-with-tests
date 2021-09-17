package main

import (
	"encoding/json"
	"io"
	"log"
)

type FileSystemPlayerStore struct {
	reader io.Reader
}

func (s FileSystemPlayerStore) GetLeague() []Player {
	var got []Player
	err := json.NewDecoder(s.reader).Decode(&got)
	if err != nil {
		log.Fatal("Could not decode database")
		return nil
	}
	return got
}
