package poker

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
)

type FileSystemPlayerStore struct {
	database *json.Encoder
	league   League
}

func NewFileSystemPlayerStore(database *os.File) (*FileSystemPlayerStore, error) {
	err := initialisePlayerDBFile(database)
	if err != nil {
		return nil, fmt.Errorf("problem initialising player db file, %v", err)
	}

	league, err := NewLeague(database)
	if err != nil {
		return nil, fmt.Errorf(
			"problem loading player store from file %s, %v",
			database.Name(),
			err,
		)
	}
	return &FileSystemPlayerStore{
		database: json.NewEncoder(&Tape{database}),
		league:   league,
	}, nil
}

func initialisePlayerDBFile(file *os.File) error {
	file.Seek(0, io.SeekStart)

	info, err := file.Stat()
	if err != nil {
		return fmt.Errorf("problem getting file info from file %s, %v", file.Name(), err)
	}

	if info.Size() == 0 {
		file.Write([]byte("[]"))
		file.Seek(0, io.SeekStart)
	}

	return nil
}

func NewLeague(database io.ReadWriteSeeker) (League, error) {
	database.Seek(0, io.SeekStart)
	var league League
	err := json.NewDecoder(database).Decode(&league)
	if err != nil {
		err = fmt.Errorf("problem parsing league, %v", err)
	}
	return league, err
}

func (f *FileSystemPlayerStore) GetLeague() League {
	sort.Slice(f.league, func(i, j int) bool {
		return f.league[i].Wins > f.league[j].Wins
	})
	return f.league
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	player := f.league.Find(name)

	if player != nil {
		return player.Wins
	}

	return 0
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
	player := f.league.Find(name)

	if player != nil {
		player.Wins++
	} else {
		newPlayer := Player{Name: name, Wins: 1}
		f.league = append(f.league, newPlayer)
	}

	f.database.Encode(f.league)
}
