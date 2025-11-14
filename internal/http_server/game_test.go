package httpserver

import (
	"bytes"
	"fmt"
	"testing"
	"time"
)

func TestGameFinish(t *testing.T) {
	store := &StubPlayerStore{}
	game := NewGame(dummySpyAlerter, store)
	winner := "Ruth"
	game.Finish(winner)
	AssertPlayerWin(t, store, winner)
}

func TestStart(t *testing.T) {

	t.Run("it schedules printing of blind values", func(t *testing.T) {
		blindAlerter := &SpyBlindAlerter{}
		game := NewGame(blindAlerter, dummyPlayerStore)
		game.Start(5)
		cases := []scheduleAlert{
			{0 * time.Second, 100},
			{10 * time.Minute, 200},
			{20 * time.Minute, 300},
			{30 * time.Minute, 400},
			{40 * time.Minute, 500},
			{50 * time.Minute, 600},
			{60 * time.Minute, 800},
			{70 * time.Minute, 1000},
			{80 * time.Minute, 2000},
			{90 * time.Minute, 4000},
			{100 * time.Minute, 8000},
		}
		checkSchedulingCases(cases, t, blindAlerter)
	})

	t.Run("it prompts the user to enter the number of players", func(t *testing.T) {
		stdOut := &bytes.Buffer{}
		game := NewGame(dummySpyAlerter, dummyPlayerStore)
		cli := NewCLI(dummyStdIn, stdOut, game)
		cli.PlayPoker()
		got := stdOut.String()
		want := PlayerPrompt

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("it prompts the user to enter the number of players", func(t *testing.T) {
		blindAlerter := &SpyBlindAlerter{}
		game := NewGame(blindAlerter, dummyPlayerStore)
		game.Start(7)

		cases := []scheduleAlert{
			{0 * time.Second, 100},
			{12 * time.Minute, 200},
			{24 * time.Minute, 300},
			{36 * time.Minute, 400},
		}
		checkSchedulingCases(cases, t, blindAlerter)
	})
}

func checkSchedulingCases(cases []scheduleAlert, t *testing.T, alerter *SpyBlindAlerter) {
	for i, want := range cases {
		t.Run(fmt.Sprint(want), func(t *testing.T) {
			if len(alerter.alerts) <= i {
				t.Fatalf("alert %d was not scheduled %v", i, alerter.alerts)
			}

			got := alerter.alerts[i]
			assertScheduledAlert(t, got, want)
		})
	}
}
