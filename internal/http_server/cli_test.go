package httpserver

import (
	"bytes"
	"strings"
	"testing"
)

var dummySpyAlerter = &SpyBlindAlerter{}
var dummyPlayerStore = &StubPlayerStore{}
var dummyStdIn = &bytes.Buffer{}
var dummyStdOut = &bytes.Buffer{}

func TestCli(t *testing.T) {
	t.Run("record chris win from user input", func(t *testing.T) {

		in := strings.NewReader("Chris wins\n")
		playerStore := &StubPlayerStore{}
		game := NewGame(dummySpyAlerter, playerStore)
		cli := NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()
		AssertPlayerWin(t, playerStore, "Chris")
	})

	t.Run("record cleo win from user input", func(t *testing.T) {

		in := strings.NewReader("Cleo wins\n")
		playerStore := &StubPlayerStore{}
		game := NewGame(dummySpyAlerter, playerStore)
		cli := NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()
		AssertPlayerWin(t, playerStore, "Cleo")
	})
	//
	// t.Run("it schedules printing of blind values", func(t *testing.T) {
	// 	in := strings.NewReader("Cleo wins\n")
	// 	playerStore := &StubPlayerStore{}
	// 	blindAlerter := &SpyBlindAlerter{}
	//
	// 	cli := NewCLI(playerStore, in, blindAlerter)
	// 	cli.PlayPoker()
	// 	if len(blindAlerter.alerts) != 1 {
	// 		t.Fatal("expected a blind alert to be scheduled")
	// 	}
	// })

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
