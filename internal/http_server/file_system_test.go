package httpserver

import (
	"io"
	"os"
	"testing"
)

func TestFileSystemStore(t *testing.T) {

	t.Run("league from reader", func(t *testing.T) {
		database, cleanDB := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer cleanDB()

		store := NewFileSystemPlayerStore(database)
		got := store.GetLeague()
		want := []Player{
			{"Cleo", 10},
			{"Chris", 33},
		}
		assertLeague(t, got, want)
		got = store.GetLeague()
		assertLeague(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		database, cleanDB := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer cleanDB()

		store := NewFileSystemPlayerStore(database)
		got := store.GetPlayerScore("Chris")
		want := 33
		assertScoreEqual(t, got, want)
	})

	t.Run("store wins for existing players", func(t *testing.T) {
		database, cleanDB := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer cleanDB()

		store := NewFileSystemPlayerStore(database)

		store.RecordWin("Chris")
		got := store.GetPlayerScore("Chris")
		want := 34
		assertScoreEqual(t, got, want)
	})

	t.Run("store wins for new player", func(t *testing.T) {
		database, cleanDB := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer cleanDB()

		store := NewFileSystemPlayerStore(database)

		store.RecordWin("Pepper")
		got := store.GetPlayerScore("Pepper")
		want := 1
		assertScoreEqual(t, got, want)
	})
}

func createTempFile(t testing.TB, initial string) (io.ReadWriteSeeker, func()) {
	t.Helper()
	tempFile, err := os.CreateTemp("", "db")
	if err != nil {
		t.Fatalf("could not create temp file")
	}

	tempFile.Write([]byte(initial))
	removeFile := func() {
		tempFile.Close()
		os.Remove(tempFile.Name())
	}

	return tempFile, removeFile
}

func assertScoreEqual(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
