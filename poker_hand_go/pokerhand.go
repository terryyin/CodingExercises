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

type Rank int

func (r Rank) comp(other Rank) int {
	return int(r - other)
}

type Card struct {
	rank Rank
}

func CreateCard(card string) Card {
	rank := Rank(strings.Index("23456789TJQKA", string(card[0])))
	return Card{rank: rank}
}

func (c Card) comp(other Card) int {
	return c.rank.comp(other.rank)
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

func (h Hand) Pair() Rank {
	for i := 0; i < 4; i++ {
		if h.cards[i].comp(h.cards[i+1]) == 0 {
			return h.cards[i].rank 
		}
	}
	return -1
}

func (h Hand) Wins(other Hand) bool {
	if h.Pair() != -1 {
		if other.Pair() != -1 {
			return h.Pair().comp(other.Pair()) > 0
		}
		return true
	}

	for i := 0; i < 5; i++ {
		if h.cards[i].comp(other.cards[i]) == 0 {
			continue
		}
		return h.cards[i].comp(other.cards[i]) > 0
	}
	return false
}

func Player1Win(game string) bool {
	h1 := CreateHand(game[0:13])
	h2 := CreateHand(game[15:])
	return h1.Wins(h2)
}
