package main

import (
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
			leastPowerfulHighCardWithHighest("8D")+" "+
				leastPowerfulHighCardWithHighest("7D")))
	})

	t.Run("should return 0 when player1 lose the only game", func(t *testing.T) {
		assert.Equal(t, "0\n", runPokerHandApp(
			leastPowerfulHighCardWithHighest("7D")+" "+
				leastPowerfulHighCardWithHighest("8D")))
	})

	t.Run("should return 2 when player1 win the only two games", func(t *testing.T) {
		assert.Equal(t, "2\n", runPokerHandApp(
			leastPowerfulHighCardWithHighest("8D")+" "+
				leastPowerfulHighCardWithHighest("7D")+"\n"+
				leastPowerfulHighCardWithHighest("8D")+" "+
				leastPowerfulHighCardWithHighest("7D")))
	})
}

func leastPowerfulHighCardWithHighest(card string) string {
	return "2H 3D 4S 5C 7C"[:14-len(card)] + card
}

func highcardGame(p1Highest string, p2Highest string) string {
	return leastPowerfulHighCardWithHighest(p1Highest) + " " + leastPowerfulHighCardWithHighest(p2Highest)
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
