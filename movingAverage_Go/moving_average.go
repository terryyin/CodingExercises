package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TradeAdvisor struct {
	firstMovingAverageLength  int
	secondMovingAverageLength int
}

func sum(data []int, size int) int {
	result := 0
	for i := 0; i < size; i++ {
		result += data[i]
	}

	return result
}

func (ta *TradeAdvisor) diff(data []int) int {
	return sum(data, ta.secondMovingAverageLength) - data[ta.secondMovingAverageLength-1]*ta.secondMovingAverageLength
}

func (ta *TradeAdvisor) hasLatestUpwardTrend(data []int) bool {
	dataLength := len(data)
	if dataLength > ta.secondMovingAverageLength {
		return ta.diff(data[dataLength-ta.secondMovingAverageLength:]) < 0
	}
	return false
}

func (ta *TradeAdvisor) hasPreviousMA1LessThanMA2WithoutInverse(data []int) bool {
	dataLength := len(data)
	for i := dataLength - ta.secondMovingAverageLength - 1; i >= 0; i-- {
		diffPrevious := ta.diff(data[i:])
		if diffPrevious > 0 {
			return true
		}
		if diffPrevious < 0 {
			break
		}
	}
	return false
}

func (ta *TradeAdvisor) getAdvice(data []int) string {
	if ta.hasLatestUpwardTrend(data) && ta.hasPreviousMA1LessThanMA2WithoutInverse(data) {
		return "buy"
	}
	return "no trade"
}

func TestWhenThereIsNoEnoughData(t *testing.T) {
	tradeAdvisor := TradeAdvisor{firstMovingAverageLength: 1, secondMovingAverageLength: 2}
	assert.Equal(t, "no trade", tradeAdvisor.getAdvice([]int{1}))
}

func TestShouldNotBuyWhenMa1IsNotPassingMa2(t *testing.T) {
	tradeAdvisor := TradeAdvisor{firstMovingAverageLength: 1, secondMovingAverageLength: 2}
	assert.Equal(t, "no trade", tradeAdvisor.getAdvice([]int{2, 2, 2}))
}

func TestShouldBuyWhenMa1IsPassingMa2(t *testing.T) {
	tradeAdvisor := TradeAdvisor{firstMovingAverageLength: 1, secondMovingAverageLength: 2}
	assert.Equal(t, "buy", tradeAdvisor.getAdvice([]int{2, 1, 2}))
}

func TestShouldNotBuyWhenPassingButNotCrossing(t *testing.T) {
	tradeAdvisor := TradeAdvisor{firstMovingAverageLength: 1, secondMovingAverageLength: 2}
	assert.Equal(t, "no trade", tradeAdvisor.getAdvice([]int{1, 1, 2}))
}

func TestShouldBuyWhenMa1IsPassingMa2With4DataPoints(t *testing.T) {
	tradeAdvisor := TradeAdvisor{firstMovingAverageLength: 1, secondMovingAverageLength: 2}
	assert.Equal(t, "buy", tradeAdvisor.getAdvice([]int{1, 2, 1, 2}))
}

func TestShouldBuyWhenMa1IsPassingMa2With4DataPoints1(t *testing.T) {
	tradeAdvisor := TradeAdvisor{firstMovingAverageLength: 1, secondMovingAverageLength: 2}
	assert.Equal(t, "buy", tradeAdvisor.getAdvice([]int{2, 1, 1, 2}))
}

func TestShouldBuyWhenMa1IsPassingMa2With4DataPoints2(t *testing.T) {
	tradeAdvisor := TradeAdvisor{firstMovingAverageLength: 1, secondMovingAverageLength: 2}
	assert.Equal(t, "buy", tradeAdvisor.getAdvice([]int{2, 1, 1, 1, 2}))
}

func TestShouldNotBuyWhenMa1IsPassingMa2Earlier(t *testing.T) {
	tradeAdvisor := TradeAdvisor{firstMovingAverageLength: 1, secondMovingAverageLength: 2}
	assert.Equal(t, "no trade", tradeAdvisor.getAdvice([]int{2, 1, 2, 2, 3}))
}

func TestShouldNotBuyWhenIncreasingOnly(t *testing.T) {
	tradeAdvisor := TradeAdvisor{firstMovingAverageLength: 1, secondMovingAverageLength: 2}
	assert.Equal(t, "no trade", tradeAdvisor.getAdvice([]int{1, 2, 3}))
}

func TestShouldNotBuyWhenIncreasingOnlyWhen2ndMaIs3(t *testing.T) {
	tradeAdvisor := TradeAdvisor{firstMovingAverageLength: 1, secondMovingAverageLength: 3}
	assert.Equal(t, "no trade", tradeAdvisor.getAdvice([]int{1, 2, 3, 4}))
}

func TestShouldNotBuyWhenPassingOnlyWhen2ndMaIs3(t *testing.T) {
	tradeAdvisor := TradeAdvisor{firstMovingAverageLength: 1, secondMovingAverageLength: 3}
	assert.Equal(t, "buy", tradeAdvisor.getAdvice([]int{4, 3, 2, 3}))
}
