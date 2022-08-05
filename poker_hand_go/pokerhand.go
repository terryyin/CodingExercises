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

func RankOfCard(rankString string) Rank {
	return Rank(strings.Index("23456789TJQKA", rankString))
}

func RanksOfCards(cardsString string) Ranks {
	ranks := Ranks{}
	for _, c := range strings.Split(cardsString, " ") {
		ranks = append(ranks, RankOfCard(c[0:1]))
	}
	sort.Slice(ranks, func(i, j int) bool {
		return ranks[i].comp(ranks[j]) > 0
	})
	return ranks
}

func CreateHand(cardsString string) Hand {
	return Hand{
		ranks: RanksOfCards(cardsString),
		flush: strings.Count(cardsString, cardsString[1:2]) == 5,
	}
}

func CreateGame(game string) Game {
	return Game{
		left:  CreateHand(game[0:14]),
		right: CreateHand(game[15:]),
	}
}

type Rank int

func (r Rank) comp(other Rank) int {
	return int(r - other)
}

type Ranks []Rank

func (ranks Ranks) without(rank Rank) Ranks {
	result := Ranks{}
	for _, r := range ranks {
		if r != rank {
			result = append(result, r)
		}
	}
	return result
}

func (ranks Ranks) nOfAKind(n int, f Ranker) {
	for _, r := range ranks {
		if len(ranks)-len(ranks.without(r)) == n {
			f(Ranks{r})
		}
	}
}

func (ranks Ranks) TwoRepeats(count1 int, count2 int, f Ranker) {
	ranks.nOfAKind(count1, func(first Ranks) {
		ranks.without(first[0]).nOfAKind(count2, func(second Ranks) {
			f(Ranks{first[0], second[0]})
		})
	})
}

func (ranks Ranks) Straight() bool {
	i := 0
	if ranks[0] == RankOfCard("A") && ranks[1] == RankOfCard("5") {
		i = 1
	}
	for ; i < len(ranks)-1; i++ {
		if ranks[i].comp(ranks[i+1]) != 1 {
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

func (ranks Ranks) compareRanks(other Ranks) int {
	for i := 0; i < min(len(ranks), len(other)); i++ {
		if ranks[i].comp(other[i]) == 0 {
			continue
		}
		return ranks[i].comp(other[i])
	}
	return len(ranks) - len(other)
}

type Hand struct {
	ranks Ranks
	flush bool
}

func (h Hand) All(f Ranker)          { f(h.ranks) }
func (h Hand) FourOfAKind(f Ranker)  { h.ranks.nOfAKind(4, f) }
func (h Hand) OnePair(f Ranker)      { h.ranks.nOfAKind(2, f) }
func (h Hand) TwoPairs(f Ranker)     { h.ranks.TwoRepeats(2, 2, f) }
func (h Hand) ThreeOfAKind(f Ranker) { h.ranks.nOfAKind(3, f) }
func (h Hand) FullHouse(f Ranker)    { h.ranks.TwoRepeats(3, 2, f) }

func (h Hand) Straight(f Ranker) {
	if h.ranks.Straight() {
		f(h.ranks)
	}
}

func (h Hand) Flush(f Ranker) {
	if h.flush {
		f(h.ranks)
	}
}

func (h Hand) StraightFlush(f Ranker) {
	h.Straight(func(straight Ranks) {
		h.Flush(func(flush Ranks) {
			f(h.ranks)
		})
	})
}

type Ranker func(r Ranks)
type Rule func(hand Hand, f Ranker)

func (functor Rule) applyToHandDefaultEmpty(hand Hand) Ranks {
	result := Ranks{}
	functor(hand, func(left Ranks) { result = left })
	return result
}

func (functor Rule) Apply(g Game) int {
	left := functor.applyToHandDefaultEmpty(g.left)
	right := functor.applyToHandDefaultEmpty(g.right)
	return left.compareRanks(right)
}

type Game struct {
	left  Hand
	right Hand
}

func (g Game) Exec(rules ...Rule) int {
	for _, rule := range rules {
		result := rule.Apply(g)
		if result != 0 {
			return result
		}
	}
	return 0
}

func (g Game) Compare() int {
	return g.Exec(
		Hand.StraightFlush,
		Hand.FourOfAKind,
		Hand.FullHouse,
		Hand.Flush,
		Hand.Straight,
		Hand.ThreeOfAKind,
		Hand.TwoPairs,
		Hand.OnePair,
		Hand.All,
	)
}
