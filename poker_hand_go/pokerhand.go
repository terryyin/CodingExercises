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

type Cards struct {
	cards []Card
}

func (c Cards) PairThen(f func(p Rank))  {
	for i := 0; i < 4; i++ {
		if c.cards[i].comp(c.cards[i+1]) == 0 {
			f(c.cards[i].rank)
			break
		}
	}
}

func (c Cards) compHighCards(other Cards) int {
	for i, card := range c.cards {
		if card.comp(other.cards[i]) == 0 {
			continue
		}
		return card.comp(other.cards[i])
	}
	return 0
}

type Hand struct {
	cards1 Cards
}

func CreateHand(cardsString string) Hand {
	cards := []Card{}
	for _, c := range strings.Split(cardsString, " ") {
		cards = append(cards, CreateCard(c))
	}
	sort.Slice(cards, func(i, j int) bool {
		return cards[i].comp(cards[j]) > 0
	})
	return Hand{cards1: Cards{cards: cards}}
}

func (h Hand) Wins(other Hand) bool {
	result := false
	h.cards1.PairThen(func(p Rank) {
		result = true
		other.cards1.PairThen(func(o Rank) {
			result = p.comp(o) > 0
		}) 
	}) 

	if(result) {
		return true
	}

	return h.cards1.compHighCards(other.cards1) > 0
}

func Player1Win(game string) bool {
	h1 := CreateHand(game[0:13])
	h2 := CreateHand(game[15:])
	return h1.Wins(h2)
}
