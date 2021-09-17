package main

import (
	"encoding/json"
	"fmt"
	"io"
)

type FileSystemPlayerStore struct {
	database io.ReadSeeker
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
	s.database.Seek(0, 0) // reset reading pointer to beginning for idempotency
	league, _ := NewLeague(s.database)
	return league
}

func (s FileSystemPlayerStore) GetPlayerScore(name string) (int, bool) {
	league := s.GetLeague()
	for _, v := range league {
		if v.Name == name {
			return v.Score, true
		}
	}
	return 0, false
}
