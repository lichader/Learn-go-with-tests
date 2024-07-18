package poker

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const PlayerPrompt = "Please enter the number of players: "

type CLI struct {
	In   *bufio.Scanner
	out  io.Writer
	game Game
}

func (cli *CLI) PlayPoker() {
	fmt.Fprint(cli.out, PlayerPrompt)

	numberOfPlayersInput := cli.readLine()
	numberOfPlayers, err := strconv.Atoi(strings.Trim(numberOfPlayersInput, "\n"))
	if err != nil {
		return
	}

	cli.game.Start(numberOfPlayers)

	winnerInput := cli.readLine()
	winner := extractWinner(winnerInput)

	cli.game.Finish(winner)
}

func NewCLI(in io.Reader, out io.Writer, game Game) *CLI {
	return &CLI{
		In:   bufio.NewScanner(in),
		out:  out,
		game: game,
	}
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}

func (cli *CLI) readLine() string {
	cli.In.Scan()
	return cli.In.Text()
}
