package main

import (
	"encoding/json"
	"fmt"
	"io"
)

type FileSystemPlayerStore struct {
	database *json.Encoder
	league   League
}

func NewFileSystemPlayerStore(db ReadWriteTruncate) (*FileSystemPlayerStore, error) {
	db.Seek(0, 0)                // reset reading pointer to beginning for idempotency
	league, err := NewLeague(db) // TODO throws error
	if err != nil {
		return nil, fmt.Errorf("Could not load league from file: %v", err)
	}
	return &FileSystemPlayerStore{league: league, database: json.NewEncoder(&tape{db})}, nil
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
	return s.league
}

func (s FileSystemPlayerStore) GetPlayerScore(name string) (int, bool) {
	player := s.GetLeague().Find(name)
	if player == nil {
		return 0, false
	}
	return player.Score, true
}

func (s *FileSystemPlayerStore) RecordWin(name string) error {
	player := s.league.Find(name)
	if player == nil {
		s.league = append(s.league, Player{name, 1})
	} else {
		player.Score++
	}
	return s.database.Encode(s.league)
}
