'''
A poker deck contains 52 cards - each card has a suit which is one of clubs, diamonds, hearts, or spades (denoted C, D, H, and S in the input data). Each card also has a value which is one of 2, 3, 4, 5, 6, 7, 8, 9, 10, jack, queen, king, ace (denoted 2, 3, 4, 5, 6, 7, 8, 9, T, J, Q, K, A). For scoring purposes, the suits are unordered while the values are ordered as given above, with 2 being the lowest and ace the highest value.

A poker hand consists of 5 cards dealt from the deck. Poker hands are ranked by the following partial order from lowest to highest.

High Card: Hands which do not fit any higher category are ranked by the value of their highest card. If the highest cards have the same value, the hands are ranked by the next highest, and so on.
Pair: 2 of the 5 cards in the hand have the same value. Hands which both contain a pair are ranked by the value of the cards forming the pair. If these values are the same, the hands are ranked by the values of the cards not forming the pair, in decreasing order.
Two Pairs: The hand contains 2 different pairs. Hands which both contain 2 pairs are ranked by the value of their highest pair. Hands with the same highest pair are ranked by the value of their other pair. If these values are the same the hands are ranked by the value of the remaining card.
Three of a Kind: Three of the cards in the hand have the same value. Hands which both contain three of a kind are ranked by the value of the 3 cards.
Straight: Hand contains 5 cards with consecutive values. Hands which both contain a straight are ranked by their highest card.
Flush: Hand contains 5 cards of the same suit. Hands which are both flushes are ranked using the rules for High Card.
Full House: 3 cards of the same value, with the remaining 2 cards forming a pair. Ranked by the value of the 3 cards.
Four of a kind: 4 cards with the same value. Ranked by the value of the 4 cards.
Straight flush: 5 cards of the same suit with consecutive values. Ranked by the highest card in the hand.
'''
class PokerHand:
    faceRanks = {'2':2, '3':3, '4':4, '5':5, '6':6, '7':7, '8':8, '9':9, '10':10, 'J':11, 'Q':12, 'K':13, 'A':14}
    def __init__(self, cardString):
        cards = cardString.split(' ')
        self.cardRanks = [self.faceRanks[x[:-1]] for x in cards]
        self.cardRanks.sort(reverse=True)
        if self.cardRanks == [14, 5, 4, 3, 2]:
            self.cardRanks = [5, 4, 3, 2, 1]
        self.isFlush = all([x[1] == cards[0][-1] for x in cards])
        
    def _getFeatures(self):
        pairs = self._getNOfAKind(2)
        triple = self._getNOfAKind(3)
        return [
             self.isFlush and self._isStraight(), #Straight Flush
             self._getNOfAKind(4),                #4 of a kind
             self._isFullHouse(triple, pairs),   
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
    
    def _isFullHouse(self, triple, pairs):
        return len(pairs) == 3 and len(triple) == 1

    def __cmp__(self, other):
        return cmp(self._getFeatures(), other._getFeatures())
        

'''
============================
Below are the unit tests
'''
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
    SmallestStraight = PokerHand("AD 2H 3H 4H 5D")
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

    def testStraightWinsThreeOfAKind(self):
        self.assertGreater(Hands.Straight, Hands.ThreeOfAKind)

    def testStraightLoseToBiggerStraight(self):
        self.assertLess(Hands.Straight, Hands.BiggerStraight)

    def testSmallestStraighWinsThreeOfAKind(self):
        self.assertGreater(Hands.SmallestStraight, Hands.ThreeOfAKind)

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
