package poker

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"time"
)

var PlayerPrompt string = "Please enter the number of players: "

type CLI struct {
	store   PlayerStore
	in      *bufio.Scanner
	out     io.Writer
	alerter Alerter
}

type Alerter interface {
	ScheduleAlertAt(duration time.Duration, amount int)
}

func NewCLI(store PlayerStore, input io.Reader, output io.Writer, alerter Alerter) *CLI {
	return &CLI{store, bufio.NewScanner(input), output, alerter}
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
	fmt.Fprint(c.out, PlayerPrompt)
	c.scheduleBlindAlerts()
	c.store.RecordWin(extractWinner(c.readLine()))
}
