package httpserver

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
)

/**
 * io.Reader read till the end of the file but can't read again
 * need to implement ReadSeeker to make reading everytime to start over
 *
 */
type FileSystemPlayerStore struct {
	Database *json.Encoder
	League   League
}

type League []Player

func (f *FileSystemPlayerStore) GetLeague() League {
	sort.Slice(f.League, func(i, j int) bool {
		return f.League[i].Wins > f.League[j].Wins
	})
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

	f.Database.Encode(f.League)
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

func NewFileSystemPlayerStore(file *os.File) (*FileSystemPlayerStore, error) {

	err := initialisePLayerDBFile(file)
	if err != nil {
		return nil, fmt.Errorf("problem initialise db file, %v", err)
	}

	league, err := NewLeague(file)
	if err != nil {
		return nil, fmt.Errorf("problem loading player from file %s, %v", file.Name(), err)
	}

	return &FileSystemPlayerStore{
		Database: json.NewEncoder(&tape{file}),
		League:   league,
	}, nil
}

func initialisePLayerDBFile(file *os.File) error {
	file.Seek(0, io.SeekStart)

	info, err := file.Stat()
	if err != nil {
		return fmt.Errorf("problem getting info from file %s, %v", file.Name(), err)
	}

	if info.Size() == 0 {
		file.Write([]byte("[]"))
		file.Seek(0, io.SeekStart)
	}

	return nil
}

func FileSystemPlayerStoreFromFile(path string) (*FileSystemPlayerStore, func(), error) {

	db, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, nil, fmt.Errorf("problem opening %s %v", path, err)
	}

	closeFunc := func() {
		db.Close()
	}

	store, err := NewFileSystemPlayerStore(db)

	if err != nil {
		return nil, nil, fmt.Errorf("problem creating file system player store %s %v", path, err)
	}

	return store, closeFunc, nil
}
