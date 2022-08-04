package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	count := 0
	ProcessLinesFromInput(func(game string) {
		if Player1Win(game) {
			count++
		}
	})
	fmt.Println(fmt.Sprintf("%v", count))
}

func ProcessLinesFromInput(processor func(string)) {
	buffer := bufio.NewReader(os.Stdin)
	for {
		input, _ := buffer.ReadString('\n')
		if len(input) == 0 {
			break
		}
		processor(input)
	}
}

type Hand struct {
	cards string
}

func (h Hand) card(index int) int {
	return rank(h.cards[index*3])
}

func (h Hand) Wins(other Hand) bool {
	if(h.card(4) == other.card(4)) {
		return h.card(3) > other.card(3)
	}
	return h.card(4) > other.card(4)
}

func Player1Win(game string) bool {
	h1 := Hand{cards: game[0:13]}
	h2 := Hand{cards: game[15:]}
	return h1.Wins(h2)
}

func rank(rank byte) int {
	return strings.Index("23456789TJQKA", string(rank))
}
