package poker

import (
	"bufio"
	"io"
	"strings"
)

type CLI struct {
	store PlayerStore
	input io.Reader
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}

func (c *CLI) PlayPoker() {
	reader := bufio.NewScanner(c.input)
	reader.Scan()
	c.store.RecordWin(extractWinner(reader.Text()))
}
