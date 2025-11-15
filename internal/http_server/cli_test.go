package httpserver

import (
	"bytes"
	"io"
	"strings"
	"testing"
	"time"
)

var dummySpyAlerter = &SpyBlindAlerter{}
var dummyPlayerStore = &StubPlayerStore{}
var dummyStdIn = &bytes.Buffer{}
var dummyStdOut = &bytes.Buffer{}

func TestCli(t *testing.T) {

	t.Run("it prompts the user to enter the number of players and starts the game", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("7\n")
		game := &GameSpy{}
		cli := NewCLI(in, stdout, game)
		cli.PlayPoker()

		gotPrompt := stdout.String()
		wantPrompt := PlayerPrompt

		if gotPrompt != wantPrompt {
			t.Errorf("got %q, want %q", gotPrompt, wantPrompt)
		}

		if game.StartedWith != 7 {
			t.Errorf("wanted start called with 7 but got %d", game.StartedWith)
		}
	})

	t.Run("finish game with 'Chris' as winner", func(t *testing.T) {
		in := strings.NewReader("1\nChris wins\n")
		game := &GameSpy{}
		cli := NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()

		if game.FinishedWith != "Chris" {
			t.Errorf("expected finish called with 'Chris' but got %q", game.FinishedWith)
		}
	})

	t.Run("record 'Cleo' win from user input", func(t *testing.T) {
		in := strings.NewReader("1\nCleo wins\n")
		game := &GameSpy{}
		cli := NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()

		if game.FinishedWith != "Cleo" {
			t.Errorf("expected finish called with 'Cleo' but got %q", game.FinishedWith)
		}
	})

	t.Run("it prints an error when a non numeric value is entered and does not start the game", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("Pies\n")
		game := &GameSpy{}
		cli := NewCLI(in, stdout, game)
		cli.PlayPoker()

		if game.StartedCalled {
			t.Errorf("game should not have started")
		}

		gotPrompt := stdout.String()
		wantPrompt := PlayerPrompt + BadPlayerInputPrompt
		if gotPrompt != wantPrompt {
			t.Errorf("got %q, want %q", gotPrompt, wantPrompt)
		}
	})
}

type failOnEndReader struct {
	t   *testing.T
	rdr io.Reader
}

func assertScheduledAlert(t testing.TB, got, want scheduleAlert) {

	amountGot := got.amount
	if amountGot != want.amount {
		t.Errorf("got amount %d, want %d", amountGot, want.amount)
	}
	gotTime := got.at
	if gotTime != want.at {
		t.Errorf("got amount %v, want %v", gotTime, want.at)
	}
}

func assertGameStartWith(t testing.TB, game *GameSpy, want int) {
	t.Helper()
	passed := retryUntil(500*time.Millisecond, func() bool {
		return game.StartedWith == want
	})

	if !passed {
		t.Errorf("invalid winner got %d, but want %d", game.StartedWith, want)
	}
}

func assertGameFinish(t testing.TB, game *GameSpy, winner string) {
	t.Helper()
	passed := retryUntil(500*time.Millisecond, func() bool {
		return game.FinishedWith == winner
	})
	if !passed {
		t.Errorf("invalid winner got %s, but want %s", game.FinishedWith, winner)
	}
}
