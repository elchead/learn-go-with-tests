package main

import (
	"encoding/json"
	"fmt"
	"io"
)

type FileSystemPlayerStore struct {
	database io.ReadWriteSeeker
}

func NewLeague(rdr io.Reader) (League, error) {
	var got League
	err := json.NewDecoder(rdr).Decode(&got)
	if err != nil {
		fmt.Errorf("Could not parse league, %v", err)
	}
	return got, err
}

func (s FileSystemPlayerStore) GetLeague() League {
	s.database.Seek(0, 0) // reset reading pointer to beginning for idempotency
	league, _ := NewLeague(s.database)
	return league
}

func (s FileSystemPlayerStore) GetPlayerScore(name string) (int, bool) {
	player := s.GetLeague().Find(name)
	if player == nil {
		return 0, false
	}
	return player.Score, true
}

func (s FileSystemPlayerStore) RecordWin(name string) error {
	league := s.GetLeague()
	player := league.Find(name)
	if player == nil {
		league = append(league, Player{name, 1})
	} else {
		player.Score++
	}
	s.database.Seek(0, 0)
	return json.NewEncoder(s.database).Encode(league)
}
