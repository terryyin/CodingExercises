package main

import (
	"fmt"
	"io"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PokerHandsApp(t *testing.T) {
	t.Run("should return 0 when the input is empty", func(t *testing.T) {
		assert.Equal(t, "0\n", runPokerHandApp(""))
	})

	t.Run("should return 1 when player1 win the only game", func(t *testing.T) {
		assert.Equal(t, "1\n", runPokerHandApp(
			game(hands.highCardEndWith("8D"), hands.highCardEndWith("7D"))))
	})

	t.Run("should return 0 when player1 lose the only game", func(t *testing.T) {
		assert.Equal(t, "0\n", runPokerHandApp(
			game(hands.highCardEndWith("7D"), hands.highCardEndWith("8D"))))
	})

	t.Run("should return 2 when player1 win the only two games", func(t *testing.T) {
		assert.Equal(t, "2\n", runPokerHandApp(
			game(hands.highCardEndWith("8D"), hands.highCardEndWith("7D"))+"\n"+
				game(hands.highCardEndWith("8D"), hands.highCardEndWith("7D"))))
	})
}

type Example int

const hands = Example(0)

func (e Example) highCardEndWith(card string) string {
	return "2H 3D 4S 5C 7C"[:14-len(card)] + card
}

func (e Example) onePairOf(rank string) string {
	return "KH QD 7S " + rank + "C " + rank + "D"
}

func (e Example) threeOfAkind(rank string) string {
	return "KH QD " + rank + "C " + rank + "C " + rank + "D"
}

func (e Example) orpanAndTwoPairs(orphan string, rank1 string, rank2 string) string {
	return orphan + "H " + rank1 + "C " + rank1 + "D " + rank2 + "C " + rank2 + "D"
}

func (e Example) twoPairsOf(rank1 string, rank2 string) string {
	return e.orpanAndTwoPairs("K", rank1, rank2)
}

func (e Example) flushStartWith(rank int) string {
	return fmt.Sprintf("%dH %dD %dS %dC %dD", rank, rank + 1, rank + 2, rank + 3, rank + 4)
}

func (e Example) fullHouse(rank3 int, rank2 int) string {
	return fmt.Sprintf("%dH %dD %dS %dC %dD", rank3, rank3, rank3, rank2, rank2)
}

func (e Example) fourOfAKind(rank int) string {
	return fmt.Sprintf("2H %dD %dS %dC %dD", rank, rank, rank, rank)
}

func game(p1 string, p2 string) string {
	return p1 + " " + p2
}

func runPokerHandApp(input string) string {
	cmd := exec.Command("go", "run", "pokerhand.go")
	stdin, _ := cmd.StdinPipe()

	go func() {
		defer stdin.Close()
		io.WriteString(stdin, input)
	}()

	out, _ := cmd.CombinedOutput()

	actual := string(out)
	return actual
}
