package poker

import (
	"bufio"
	"io"
	"strings"
)

type CLI struct {
	store PlayerStore
	in    *bufio.Scanner
}

func NewCLI(store PlayerStore, input io.Reader) *CLI {
	return &CLI{store, bufio.NewScanner(input)}
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}

func (c *CLI) PlayPoker() {
	c.store.RecordWin(extractWinner(c.readLine()))
}
