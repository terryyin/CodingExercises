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
		if CreateGame(game).Compare() > 0 {
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (c Ranks) compareRanks(other Ranks) int {
	for i := 0; i < min(len(c.ranks), len(other.ranks)); i++ {
		if c.ranks[i].comp(other.ranks[i]) == 0 {
			continue
		}
		return c.ranks[i].comp(other.ranks[i])
	}
	return len(c.ranks) - len(other.ranks)
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

func (h Hand) StraightFlush(f func(ranks Ranks)) {
	h.Straight(func(straight Ranks) {
		h.Flush(func(flush Ranks) {
			f(h.ranks)
		})
	})
}

func (h Hand) Apply(functor func(hand Hand, f func(r Ranks))) Ranks {
	result := Ranks{ranks: []Rank{}}
	functor(h, func(left Ranks) {
		result = left
	})
	return result
}

type Rule func(hand Hand, f func(r Ranks))

type Game struct {
	left  Hand
	right Hand
}

func CreateGame(game string) Game {
	return Game{
		left:  CreateHand(game[0:14]),
		right: CreateHand(game[15:]),
	}
}

func (g Game) Rule(functor Rule) int {
	left := g.left.Apply(functor)
	right := g.right.Apply(functor)
	return left.compareRanks(right)
}

func (g Game) Exec(rules []Rule) int {
	for _, l := range rules {
		result := g.Rule(l)
		if result != 0 {
			return result
		}
	}
	return 0
}

func (g Game) Compare() int {
	return g.Exec([]Rule{
		Hand.StraightFlush,
		Hand.FourOfAKind,
		Hand.FullHouse,
		Hand.Flush,
		Hand.Straight,
		Hand.ThreeOfAKind,
		Hand.TwoPairs,
		Hand.OnePair,
		Hand.All})
}
