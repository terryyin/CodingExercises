package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

type Card struct {
	rank int
}

func CreateCard(card string) Card {
	rank := strings.Index("23456789TJQKA", string(card[0]))
	return Card{rank: rank}
}

func (c Card) comp(other Card) int {
	return c.rank - other.rank
}

type Hand struct {
	cards []Card
}

func CreateHand(cardsString string) Hand {
	cards := []Card{}
	for _, c := range strings.Split(cardsString, " ") {
		cards = append(cards, CreateCard(c))
	}
	sort.Slice(cards, func(i, j int) bool {
		return cards[i].comp(cards[j]) > 0
	})
	return Hand{cards: cards}
}

func (h Hand) card(index int) Card {
	return h.cards[index]
}

func (h Hand) Wins(other Hand) bool {
	if(h.cards[4].comp(h.cards[3]) == 0) {
		return true
	}
	for i := 0; i < 5; i++ {
		if h.card(i).comp(other.card(i)) == 0 {
			continue
		}
		return h.card(i).comp(other.card(i)) > 0
	}
	return false
}

func Player1Win(game string) bool {
	h1 := CreateHand(game[0:13])
	h2 := CreateHand(game[15:])
	return h1.Wins(h2)
}
