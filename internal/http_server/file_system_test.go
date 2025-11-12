package httpserver

import (
	"log"
	"os"
	"testing"
)

func TestFileSystemStore(t *testing.T) {

	t.Run("league from reader", func(t *testing.T) {
		database, cleanDB := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer cleanDB()

		store, err := NewFileSystemPlayerStore(database)
		assertNoError(t, err)

		got := store.GetLeague()
		want := []Player{
			{"Chris", 33},
			{"Cleo", 10},
		}
		AssertLeague(t, got, want)
		got = store.GetLeague()
		AssertLeague(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		database, cleanDB := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer cleanDB()

		store, err := NewFileSystemPlayerStore(database)
		if err != nil {
			log.Fatalf("problem creating system player store %v", err)
		}
		got := store.GetPlayerScore("Chris")
		want := 33
		assertScoreEqual(t, got, want)
	})

	t.Run("store wins for existing players", func(t *testing.T) {
		database, cleanDB := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer cleanDB()

		store, err := NewFileSystemPlayerStore(database)
		if err != nil {
			log.Fatalf("problem creating system player store %v", err)
		}

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

		store, err := NewFileSystemPlayerStore(database)
		if err != nil {
			log.Fatalf("problem creating system player store %v", err)
		}

		store.RecordWin("Pepper")
		got := store.GetPlayerScore("Pepper")
		want := 1
		assertScoreEqual(t, got, want)
	})

	t.Run("work with empty file", func(t *testing.T) {
		database, cleanDB := createTempFile(t, "")
		defer cleanDB()

		_, err := NewFileSystemPlayerStore(database)
		assertNoError(t, err)
	})
}

func createTempFile(t testing.TB, initial string) (*os.File, func()) {
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

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("didn't expect an error, but found one, %v", err)
	}
}
