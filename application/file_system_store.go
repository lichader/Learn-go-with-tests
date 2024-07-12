package main

import (
	"encoding/json"
	"fmt"
	"io"
)

type FileSystemPlayerStore struct {
	database io.ReadWriteSeeker
}

func (f *FileSystemPlayerStore) GetLeague() League {
	f.database.Seek(0, io.SeekStart)
	league, _ := f.NewLeague()
	return league
}

func (f *FileSystemPlayerStore) NewLeague() (League, error) {
	var league []Player
	err := json.NewDecoder(f.database).Decode(&league)
	if err != nil {
		err = fmt.Errorf("problem parsing league, %v", err)
	}
	return league, err
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	player := f.GetLeague().Find(name)

	if player != nil {
		return player.Wins
	}

	return 0
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
	league := f.GetLeague()
	player := league.Find(name)

	if player != nil {
		player.Wins++
	} else {
		newPlayer := Player{Name: name, Wins: 1}
		league = append(league, newPlayer)
	}

	f.database.Seek(0, io.SeekStart)
	json.NewEncoder(f.database).Encode(league)
}
