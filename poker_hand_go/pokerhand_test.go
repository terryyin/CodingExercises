package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PokerHand(t *testing.T) {
	assertWin := func(left string, right string) {
		assert.Greaterf(t, CreateGame(game(left, right)).Compare(), 0, "game: %v vs %v", CreateHand(left), CreateHand(right))
	}
	assertNotWin := func(left string, right string) {
		assert.LessOrEqualf(t, CreateGame(game(left, right)).Compare(), 0, "game: %v vs %v", CreateHand(left), CreateHand(right))
	}

	t.Run("win by highest highcard", func(t *testing.T) {
		assertWin(hands.highCardEndWith("9D"), hands.highCardEndWith("8D"))
		assertWin(hands.highCardEndWith("TD"), hands.highCardEndWith("9D"))
		assertWin(hands.highCardEndWith("AD"), hands.highCardEndWith("QD"))
	})

	t.Run("peace", func(t *testing.T) {
		assertNotWin(hands.highCardEndWith("9D"), hands.highCardEndWith("9D"))
	})

	t.Run("win by highest highcard unordered", func(t *testing.T) {
		assertWin(hands.highCardEndWith("9D 6C"), hands.highCardEndWith("8D 7C"))
	})

	t.Run("lose by highest highcard", func(t *testing.T) {
		assertNotWin(hands.highCardEndWith("8D"), hands.highCardEndWith("9D"))
	})

	t.Run("win by 2nd highest highcard", func(t *testing.T) {
		assertWin(hands.highCardEndWith("8C 9D"), hands.highCardEndWith("9D"))
	})

	t.Run("win by one pair vs high card", func(t *testing.T) {
		assertWin(hands.onePairOf("2"), hands.highCardEndWith("KD"))
		assertWin(hands.onePairOf("8"), hands.highCardEndWith("KD"))
	})

	t.Run("lose by high card vs one pair", func(t *testing.T) {
		assertNotWin(hands.highCardEndWith("KD"), hands.onePairOf("2"))
	})

	t.Run("one pair vs one pair", func(t *testing.T) {
		assertNotWin(hands.onePairOf("2"), hands.onePairOf("3"))
		assertWin(hands.onePairOf("3"), hands.onePairOf("2"))
		assertWin("2H 2C 3D 4S AC", "2H 2C 3D 4S 5D")
	})

	t.Run("two pairs vs one pair", func(t *testing.T) {
		assertWin(hands.twoPairsOf("2", "3"), hands.onePairOf("4"))
	})

	t.Run("two pairs vs two pairs", func(t *testing.T) {
		assertNotWin(hands.twoPairsOf("2", "3"), hands.twoPairsOf("3", "4"))
		assertWin(hands.orpanAndTwoPairs("Q", "4", "5"), hands.twoPairsOf("3", "5"))
	})

	t.Run("three of a kind", func(t *testing.T) {
		assertWin(hands.threeOfAkind("2"), hands.twoPairsOf("3", "4"))
	})

	t.Run("straight", func(t *testing.T) {
		assertWin(hands.straightStartWith(2), hands.threeOfAkind("2"))
		assertWin("2D 3H 4H 5H AH", hands.threeOfAkind("2"))
	})

	t.Run("flush", func(t *testing.T) {
		assertWin(hands.flush(), hands.straightStartWith(2))
	})

	t.Run("full house", func(t *testing.T) {
		assertWin(hands.fullHouse(2, 3), hands.straightStartWith(2))
	})

	t.Run("four of a kind", func(t *testing.T) {
		assertWin(hands.fourOfAKind(3), hands.fullHouse(2, 3))
	})

	t.Run("four of a kind", func(t *testing.T) {
		assertWin(hands.straightFlush(), hands.fourOfAKind(3))
	})

}
