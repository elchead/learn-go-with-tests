package poker

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
)

var PlayerPrompt string = "Please enter the number of players: "

type Gamer interface {
	Start(numberPlayers int)
	Finish(name string)
}

type CLI struct {
	game Gamer
	in   *bufio.Scanner
	out  io.Writer
}

func NewCLI(game Gamer, input io.Reader, output io.Writer) *CLI {
	return &CLI{game, bufio.NewScanner(input), output}
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}

func (c *CLI) PlayPoker() {
	fmt.Fprint(c.out, PlayerPrompt)
	numberPlayers, err := strconv.Atoi(c.readLine())
	if err != nil {
		log.Fatal(err)
	}
	c.game.Start(numberPlayers)
	c.game.Finish(extractWinner(c.readLine()))
}
