package httpserver

import (
	"encoding/json"
	"io"
)

/**
 * io.Reader read till the end of the file but can't read again
 * need to implement ReadSeeker to make reading everytime to start over
 *
 */
type FileSystemPlayerStore struct {
	Database io.Writer
	League   League
}

type League []Player

func (f *FileSystemPlayerStore) GetLeague() League {
	return f.League
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	player := f.League.Find(name)

	if player != nil {
		return player.Wins
	}
	return 0
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
	player := f.League.Find(name)

	if player != nil {
		player.Wins++
	} else {
		f.League = append(f.League, Player{name, 1})
	}

	json.NewEncoder(f.Database).Encode(f.League)
}

func (l League) Find(name string) *Player {
	for i, p := range l {
		// use league[i] to get the reference of Player instead of copy using player.Wins
		if p.Name == name {
			return &l[i]
		}
	}

	return nil
}

func NewFileSystemPlayerStore(database io.ReadWriteSeeker) *FileSystemPlayerStore {
	database.Seek(0, io.SeekStart)
	league, _ := NewLeague(database)
	return &FileSystemPlayerStore{
		Database: &tape{database},
		League:   league,
	}
}
