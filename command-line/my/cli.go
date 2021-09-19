package poker

import (
	"bufio"
	"io"
	"strings"
	"time"
)

type CLI struct {
	store PlayerStore
	in    *bufio.Scanner
}

type Alerter interface {
	ScheduleAlertAt(duration time.Duration, amount int)
}

func NewCLI(store PlayerStore, input io.Reader, alerter Alerter) *CLI {
	alerter.ScheduleAlertAt(10*time.Second, 10)
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
