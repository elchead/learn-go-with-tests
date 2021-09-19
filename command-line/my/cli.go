package poker

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

var PlayerPrompt string = "Please enter the number of players: "

var BadPlayerInputErrMsg string = "Could not parse number of players"

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

func extractWinner(userInput string) (string, error) {
	matched, _ := regexp.MatchString(" wins", userInput)
	if !matched {
		return "", errors.Errorf("Invalid winner input: %s", userInput)
	}
	return strings.Replace(userInput, " wins", "", 1), nil
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}

func (c *CLI) PlayPoker() error {
	fmt.Fprint(c.out, PlayerPrompt)
	numberPlayers, err := strconv.Atoi(c.readLine())
	if err != nil {
		return errors.New(BadPlayerInputErrMsg)
	}
	c.game.Start(numberPlayers)
	winner, err := extractWinner(c.readLine())
	if err != nil {
		return err
	}
	c.game.Finish(winner)
	return nil
}
