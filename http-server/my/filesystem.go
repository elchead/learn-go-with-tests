package main

import (
	"encoding/json"
	"fmt"
	"io"
)

type FileSystemPlayerStore struct {
	database io.Reader
}

func NewLeague(rdr io.Reader) ([]Player, error) {
	var got []Player
	err := json.NewDecoder(rdr).Decode(&got)
	if err != nil {
		fmt.Errorf("Could not parse league, %v", err)
	}
	return got, err
}

func (s FileSystemPlayerStore) GetLeague() []Player {
	league, _ := NewLeague(s.database)
	return league
}
