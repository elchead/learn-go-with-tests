package poker

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strconv"
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

func (cli *CLI) scheduleBlindAlerts(numberPlayers int) {
	blindIncrement := time.Duration(5+numberPlayers) * time.Minute
	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second
	for _, blind := range blinds {
		cli.alerter.ScheduleAlertAt(blindTime, blind)
		blindTime = blindTime + blindIncrement
	}
}

func (c *CLI) PlayPoker() {
	fmt.Fprint(c.out, PlayerPrompt)
	numberPlayers, err := strconv.Atoi(c.readLine())
	if err != nil {
		log.Fatal(err)
	}
	c.scheduleBlindAlerts(numberPlayers)
	c.store.RecordWin(extractWinner(c.readLine())) // TODO fix other tests
}
