class PokerHand:
    faceRanks = {'2':2, '3':3, '4':4, '5':5, '6':6, '7':7, '8':8, '9':9, '10':10, 'J':11, 'Q':12, 'K':13, 'A':14}
    def __init__(self, cardString):
        cards = cardString.split(' ')
        self.cardRanks = [self.faceRanks[x[:-1]] for x in cards]
        self.cardRanks.sort(reverse=True)
        self.isFlush = all([x[1] == cards[0][-1] for x in cards])
        
    def _getFeatures(self):
        pairs = self._getNOfAKind(2)
        triple = self._getNOfAKind(3)
        return [
             self.isFlush and self._isStraight(), #Straight Flush
             self._getNOfAKind(4),                #4 of a kind
             self._getFullHouse(triple, pairs),   
             self.isFlush,
             self._isStraight(),
             triple,
             len(pairs),
             pairs,
             self.cardRanks
             ]
    
    def _getNOfAKind(self, n):
        reduceSameNeighbors = lambda cards, z: [pair[0] for pair in zip(cards[:-1], cards[1:]) if pair[0] == pair[1]]
        return reduce(reduceSameNeighbors, range(n - 1), self.cardRanks)

    def _isStraight(self):
        return all([neighbor[0] == neighbor[1] + 1 for neighbor in zip(self.cardRanks[:-1], self.cardRanks[1:])])
    
    def _getFullHouse(self, triple, pairs):
        if len(pairs) == 3 and len(triple) == 1:
            return [triple, pairs]

    def __cmp__(self, other):
        return cmp(self._getFeatures(), other._getFeatures())

import unittest
            
class Hands:
    HighCard = PokerHand("2H 4H 6H 7H 9D")
    HighCardWithDifferentSuits = PokerHand("2H 4H 6D 7D 9D")
    BiggerHighCard = PokerHand("2H 3H 7H 8H 9D")
    OnePair = PokerHand("2H 2H 6H 7H 9D")
    OnePairWithBiggerPair = PokerHand("4H 4H 5H 7H 9D")
    TwoPairs = PokerHand("2H 2H 3H 3H 9D")
    ThreeOfAKind = PokerHand("2H 2H 2H 3H 9D")
    Straight = PokerHand("2H 3H 4H 5H 6D")
    BiggerStraight = PokerHand("3H 4H 5H 6D 7D")
    Flush = PokerHand("2H 3H 6H 7H 9H")
    FullHouse = PokerHand("3H 3H 3H 4D 4D")
    FourOfAKind = PokerHand("3H 3H 3H 3D 4D")
    StraightFlush = PokerHand("3H 4H 5H 6H 7H")

class TestPokerHand(unittest.TestCase):

    def testCardRands(self):
        self.assertGreater(PokerHand("JH"), PokerHand("10H"))
        self.assertGreater(PokerHand("QH"), PokerHand("JH"))
        self.assertGreater(PokerHand("KH"), PokerHand("QH"))
        self.assertGreater(PokerHand("AH"), PokerHand("KH"))

    def testHighCardHandWithHigherHighestCardShouldWin(self):
        self.assertGreater(Hands.BiggerHighCard, Hands.HighCard)

    def testHighCardHandsShouldIgnoreSuits(self):
        self.assertEqual(Hands.HighCard, Hands.HighCardWithDifferentSuits)
        
    def testOnePairShouldWinHighCard(self):
        self.assertGreater(Hands.OnePair, Hands.HighCard)

    def testHighCardWithHigherPairShouldWin(self):
        self.assertGreater(Hands.OnePairWithBiggerPair, Hands.OnePair)

    def testTwoPairsWinsOnePair(self):
        self.assertGreater(Hands.TwoPairs, Hands.OnePairWithBiggerPair)

    def testThreeOfAKindWinsTwoPairs(self):
        self.assertGreater(Hands.ThreeOfAKind, Hands.TwoPairs)

    def testStraightWinsTwoPairs(self):
        self.assertGreater(Hands.Straight, Hands.ThreeOfAKind)

    def testStraightLoseToBiggerStraight(self):
        self.assertLess(Hands.Straight, Hands.BiggerStraight)

    def testFlushWinsStraight(self):
        self.assertGreater(Hands.Flush, Hands.Straight)

    def testFullHouseWinsFlush(self):
        self.assertGreater(Hands.FullHouse, Hands.Flush)

    def testFourOfAKindWinsFullHouse(self):
        self.assertGreater(Hands.FourOfAKind, Hands.FullHouse)

    def testStraightFlushWinsFourOfAKind(self):
        self.assertGreater(Hands.StraightFlush, Hands.FourOfAKind)

if __name__ == "__main__":
    unittest.main()
