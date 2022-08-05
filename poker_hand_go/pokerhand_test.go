package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PokerHand(t *testing.T) {
	assertWin := func(game string) {
		assert.Truef(t, Player1Win(game), "game: %v", game)
	}
	assertNotWin := func(game string) {
		assert.Falsef(t, Player1Win(game), "game: %v", game)
	}

	t.Run("win by highest highcard", func(t *testing.T) {
		assertWin(highcardGame("9D", "8D"))
		assertWin(highcardGame("TD", "9D"))
		assertWin(highcardGame("AD", "QD"))
	})

	t.Run("peace", func(t *testing.T) {
		assertNotWin(highcardGame("9D", "9D"))
	})

	t.Run("win by highest highcard unordered", func(t *testing.T) {
		assertWin(highcardGame("9D 6C", "8D 7C"))
	})

	t.Run("lose by highest highcard", func(t *testing.T) {
		assertNotWin(highcardGame("8D", "9D"))
	})

	t.Run("win by 2nd highest highcard", func(t *testing.T) {
		assertWin(highcardGame("8C 9D", "9D"))
	})

	t.Run("win by one pair vs high card", func(t *testing.T) {
		assertWin(game(handWithOnePairOf("2"), leastPowerfulHighCardWithHighest("AD")))
		assertWin(game(handWithOnePairOf("8"), leastPowerfulHighCardWithHighest("AD")))
	})

	t.Run("lose by high card vs one pair", func(t *testing.T) {
		assertNotWin(game( leastPowerfulHighCardWithHighest("AD"), handWithOnePairOf("2")))
	})

	t.Run("one pair vs one pair", func(t *testing.T) {
		assertNotWin(game(handWithOnePairOf("2"), handWithOnePairOf("3")))
		assertWin(game(handWithOnePairOf("3"), handWithOnePairOf("2")))
		assertWin(game("2H 2C 3D 4S AC", "2H 2C 3D 4S 5D"))
	})

	t.Run("two pairs vs one pair", func(t *testing.T) {
		assertWin(game(handWithTwoPairsOf("2", "3"), handWithOnePairOf("4")))
	})

	t.Run("two pairs vs two pairs", func(t *testing.T) {
		assertNotWin(game(handWithTwoPairsOf("2", "3"), handWithTwoPairsOf("3", "4")))
		assertWin(game(handWithOrpanAndTwoPairs("Q", "4", "5"), handWithTwoPairsOf("3", "5")))
	})

}
