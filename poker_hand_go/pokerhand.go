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

func RankOfCard(card string) Rank {
	return Rank(strings.Index("23456789TJQKA", string(card[0])))
}

type Ranks struct {
	ranks []Rank
}

func (c Ranks) pairThen(f func(index int, p Rank)) {
	for i := 0; i < len(c.ranks)-1; i++ {
		if c.ranks[i].comp(c.ranks[i+1]) == 0 {
			f(i, c.ranks[i])
			break
		}
	}
}

func (c Ranks) PairThen(f func(ranks Ranks)) {
	c.pairThen(func(i int, p Rank) { f(Ranks{ranks: []Rank{p}}) })
}

func (c Ranks) TwoPairsThen(f func(ranks Ranks)) {
	c.pairThen(func(i int, high Rank) {
		r := []Rank{}
		r = append(r, c.ranks[:i]...)
		remnent := Ranks{ranks: append(r, c.ranks[i+2:]...)}
		remnent.pairThen(func(_ int, low Rank) {
			f(Ranks{ranks: []Rank{low, high}})
		})
	})
}

func (c Ranks) compareRanks(other Ranks) int {
	for i, card := range c.ranks {
		if card.comp(other.ranks[i]) == 0 {
			continue
		}
		return card.comp(other.ranks[i])
	}
	return 0
}

type Hand struct {
	cards Ranks
}

func CreateHand(cardsString string) Hand {
	ranks := []Rank{}
	for _, c := range strings.Split(cardsString, " ") {
		ranks = append(ranks, RankOfCard(c))
	}
	sort.Slice(ranks, func(i, j int) bool {
		return ranks[i].comp(ranks[j]) > 0
	})
	return Hand{cards: Ranks{ranks: ranks}}
}

func rule(left Hand, right Hand, finder func(ranks Ranks, f func(r Ranks))) int {
	result := 0
	finder(left.cards, func(left Ranks) {
		result = 1
		finder(right.cards, func(right Ranks) {
			result = left.compareRanks(right)
		})
	})

	if result == 0 {
		finder(right.cards, func(right Ranks) {
			result = -1
			finder(left.cards, func(left Ranks) {
				result = left.compareRanks(right)
			})
		})
	}

	return result
}

func (h Hand) Wins(other Hand) bool {
	result := 0

	result = rule(h, other, Ranks.TwoPairsThen)

	if result != 0 {
		return result > 0
	}

	result = rule(h, other, Ranks.PairThen)

	if result != 0 {
		return result > 0
	}

	return h.cards.compareRanks(other.cards) > 0
}

func Player1Win(game string) bool {
	h1 := CreateHand(game[0:13])
	h2 := CreateHand(game[15:])
	return h1.Wins(h2)
}
