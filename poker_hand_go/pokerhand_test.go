package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PokerHand(t *testing.T) {

	t.Run("win by highest highcard", func(t *testing.T) {
		assert.True(t, Player1Win(highcardGame("9D", "8D")))
		assert.True(t, Player1Win(highcardGame("TD", "9D")))
		assert.True(t, Player1Win(highcardGame("AD", "QD")))
	})

	t.Run("peace", func(t *testing.T) {
		assert.False(t, Player1Win(highcardGame("9D", "9D")))
	})

	t.Run("win by highest highcard unordered", func(t *testing.T) {
		assert.True(t, Player1Win(highcardGame("9D 6C", "8D 7C")))
	})

	t.Run("lose by highest highcard", func(t *testing.T) {
		assert.False(t, Player1Win(highcardGame("8D", "9D")))
	})

	t.Run("win by 2nd highest highcard", func(t *testing.T) {
		assert.True(t, Player1Win(highcardGame("8C 9D", "9D")))
	})

	t.Run("win by one pair vs high card", func(t *testing.T) {
		assert.True(t, Player1Win(game(handWithOnePairOf("2"), leastPowerfulHighCardWithHighest("AD"))))
		assert.True(t, Player1Win(game(handWithOnePairOf("8"), leastPowerfulHighCardWithHighest("AD"))))
	})

	t.Run("one pair vs one pair", func(t *testing.T) {
		assert.False(t, Player1Win(game(handWithOnePairOf("2"), handWithOnePairOf("3"))))
		assert.True(t, Player1Win(game(handWithOnePairOf("3"), handWithOnePairOf("2"))))
	})

}