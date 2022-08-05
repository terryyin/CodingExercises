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

func (c Ranks) All(f func(ranks Ranks)) {
	f(c)
}

func (c Ranks) OnePair(f func(ranks Ranks)) {
	c.nOfAKind(2, f)
}

func (c Ranks) TwoPairs(f func(ranks Ranks)) {
	c.OnePair(func(high Ranks) {
		c.without(high).OnePair(func(low Ranks) {
			f(Ranks{ranks: []Rank{high.ranks[0], low.ranks[0]}})
		})
	})
}

func (c Ranks) ThreeOfAKind(f func(ranks Ranks)) {
	c.nOfAKind(3, f)
}

func (c Ranks) FourOfAKind(f func(ranks Ranks)) {
	c.nOfAKind(4, f)
}

func (c Ranks) Flush(f func(ranks Ranks)) {
	for i := 0; i < len(c.ranks)-1; i++ {
		if c.ranks[i].comp(c.ranks[i+1]) != 1 {
			return
		}
	}
	f(Ranks{ranks: []Rank{1}})
}

func (c Ranks) FullHouse(f func(ranks Ranks)) {
	c.ThreeOfAKind(func(three Ranks) {
		c.without(three).OnePair(func(two Ranks) {
			f(Ranks{ranks: three.ranks})
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

type Result struct {
	result int
}

func (r Result) oneSideRule(left Hand, right Hand, finder func(ranks Ranks, f func(r Ranks))) Result {
	if r.result != 0 {
		return r
	}
	finder(left.cards, func(left Ranks) {
		r.result = 1
		finder(right.cards, func(right Ranks) {
			r.result = left.compareRanks(right)
		})
	})
	return r
}

func (r Result) reverse() Result {
	return Result{result: -r.result}
}

func (r Result) Rule(left Hand, right Hand, finder func(ranks Ranks, f func(r Ranks))) Result {
	next := r.oneSideRule(left, right, finder)
	if next.result != 0 {
		return next
	}
	return r.oneSideRule(right, left, finder).reverse()
}

func (h Hand) Wins(other Hand) bool {
	return Result{result: 0}.
		Rule(h, other, Ranks.FourOfAKind).
		Rule(h, other, Ranks.FullHouse).
		Rule(h, other, Ranks.Flush).
		Rule(h, other, Ranks.ThreeOfAKind).
		Rule(h, other, Ranks.TwoPairs).
		Rule(h, other, Ranks.OnePair).
		Rule(h, other, Ranks.All).result > 0
}

func Player1Win(game string) bool {
	h1 := CreateHand(game[0:13])
	h2 := CreateHand(game[15:])
	return h1.Wins(h2)
}
