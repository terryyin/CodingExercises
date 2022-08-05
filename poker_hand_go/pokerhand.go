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

func (c Ranks) withoutRank(rank Rank) Ranks {
	ranks := []Rank{}
	for _, r := range c.ranks {
		if r != rank {
			ranks = append(ranks, r)
		}
	}
	return Ranks{ranks: ranks}
}

func (c Ranks) without(rks Ranks) Ranks {
	return c.withoutRank(rks.ranks[0])
}

func (c Ranks) nOfAKind(n int, f func(ranks Ranks)) {
	for _, r := range c.ranks {
		if len(c.ranks)-len(c.withoutRank(r).ranks) == n {
			f(Ranks{ranks: []Rank{r}})
		}
	}
}

func (c Ranks) Straight() bool {
	for i := 0; i < len(c.ranks)-1; i++ {
		if c.ranks[i].comp(c.ranks[i+1]) != 1 {
			return false
		}
	}
	return true
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
	ranks Ranks
	flush int
}

func CreateHand(cardsString string) Hand {
	ranks := []Rank{}
	for _, c := range strings.Split(cardsString, " ") {
		ranks = append(ranks, RankOfCard(c))
	}
	sort.Slice(ranks, func(i, j int) bool {
		return ranks[i].comp(ranks[j]) > 0
	})
	return Hand{
		ranks: Ranks{ranks: ranks},
		flush: strings.Count(cardsString, cardsString[1:2]),
	}
}

func (h Hand) All(f func(ranks Ranks)) {
	f(h.ranks)
}

func (h Hand) FourOfAKind(f func(ranks Ranks)) {
	h.ranks.nOfAKind(4, f)
}

func (h Hand) Straight(f func(ranks Ranks)) {
	if h.ranks.Straight() {
		f(Ranks{ranks: []Rank{1}})
	}
}

func (h Hand) OnePair(f func(ranks Ranks)) {
	h.ranks.nOfAKind(2, f)
}

func (h Hand) TwoPairs(f func(ranks Ranks)) {
	h.OnePair(func(high Ranks) {
		h.ranks.without(high).nOfAKind(2, func(low Ranks) {
			f(Ranks{ranks: []Rank{high.ranks[0], low.ranks[0]}})
		})
	})
}

func (h Hand) ThreeOfAKind(f func(ranks Ranks)) {
	h.ranks.nOfAKind(3, f)
}

func (h Hand) FullHouse(f func(ranks Ranks)) {
	h.ThreeOfAKind(func(three Ranks) {
		h.ranks.without(three).nOfAKind(2, func(two Ranks) {
			f(Ranks{ranks: three.ranks})
		})
	})
}

func (h Hand) Flush(f func(ranks Ranks)) {
	if h.flush == 5 {
		f(h.ranks)
	}
}

type Result struct {
	result int
}

func (r Result) reverse() Result {
	return Result{result: -r.result}
}

func (r Result) oneSideRule(left Hand, right Hand, functor func(hand Hand, f func(r Ranks))) Result {
	if r.result != 0 {
		return r
	}
	functor(left, func(left Ranks) {
		r.result = 1
		functor(right, func(right Ranks) {
			r.result = left.compareRanks(right)
		})
	})
	return r
}

func (r Result) Rule(left Hand, right Hand, functor func(hand Hand, f func(r Ranks))) Result {
	next := r.oneSideRule(left, right, functor)
	if next.result != 0 {
		return next
	}
	return r.oneSideRule(right, left, functor).reverse()
}

func (h Hand) Wins(other Hand) bool {
	return Result{result: 0}.
		Rule(h, other, Hand.FourOfAKind).
		Rule(h, other, Hand.FullHouse).
		Rule(h, other, Hand.Flush).
		Rule(h, other, Hand.Straight).
		Rule(h, other, Hand.ThreeOfAKind).
		Rule(h, other, Hand.TwoPairs).
		Rule(h, other, Hand.OnePair).
		Rule(h, other, Hand.All).result > 0
}

func Player1Win(game string) bool {
	h1 := CreateHand(game[0:14])
	h2 := CreateHand(game[15:])
	return h1.Wins(h2)
}
