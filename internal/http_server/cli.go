package httpserver

import (
	"bufio"
	"io"
	"strings"
)

type CLI struct {
	playerStore PlayerStore
	in          *bufio.Scanner
}

func (c *CLI) PlayPoker() {
	userInput := c.readLine()
	c.playerStore.RecordWin(extractWinner(userInput))
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}

func (c *CLI) readLine() string {
	c.in.Scan()
	return c.in.Text()
}

func NewCLI(store PlayerStore, in io.Reader) *CLI {
	return &CLI{
		playerStore: store,
		in:          bufio.NewScanner(in),
	}
}
