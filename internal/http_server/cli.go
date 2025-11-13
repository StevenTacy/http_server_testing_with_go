package httpserver

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
)

const PlayerPrompt = "Please enter the number of players: "

type CLI struct {
	in   *bufio.Scanner
	out  io.Writer
	game Game
}

type SpyBlindAlerter struct {
	alerts []scheduleAlert
}

type scheduleAlert struct {
	at     time.Duration
	amount int
}

func (c *CLI) PlayPoker() {
	fmt.Fprint(c.out, PlayerPrompt)
	input := c.readLine()
	numberOfPlayers, _ := strconv.Atoi(strings.Trim(input, "\n"))
	c.game.Start(numberOfPlayers)
	userInput := c.readLine()
	winner := extractWinner(userInput)
	c.game.Finish(winner)
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}

func (c *CLI) readLine() string {
	c.in.Scan()
	return c.in.Text()
}

func NewCLI(in io.Reader, out io.Writer, game Game) *CLI {
	return &CLI{
		in:   bufio.NewScanner(in),
		out:  out,
		game: game,
	}
}

func (s scheduleAlert) String() string {
	return fmt.Sprintf("%d chips at %v", s.amount, s.at)
}

func (s *SpyBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int) {
	s.alerts = append(s.alerts, scheduleAlert{duration, amount})
}
