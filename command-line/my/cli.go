package poker

import (
	"bufio"
	"io"
	"strings"
	"time"
)

type CLI struct {
	store   PlayerStore
	in      *bufio.Scanner
	alerter Alerter
}

type Alerter interface {
	ScheduleAlertAt(duration time.Duration, amount int)
}

func NewCLI(store PlayerStore, input io.Reader, alerter Alerter) *CLI {
	return &CLI{store, bufio.NewScanner(input), alerter}
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}

func (cli *CLI) scheduleBlindAlerts() {
	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	for i := 0; i < 11; i++ {
		cli.alerter.ScheduleAlertAt(time.Duration(i*10)*time.Minute, blinds[i])
	}
}

func (c *CLI) PlayPoker() {
	c.scheduleBlindAlerts()
	c.store.RecordWin(extractWinner(c.readLine()))
}
