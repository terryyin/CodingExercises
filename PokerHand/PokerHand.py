class PokerHand:
    ''' This is a solution to the Poker Hand Kata
    '''
    def __init__(self, cardString):
        cards = cardString.split(' ')
        cards.sort(reverse=True)
        self.cards = [int(x[0]) for x in cards]
        isFlush = all([x[1] == cards[0][1] for x in cards])
        self.features = [
                         isFlush and self.getStraight(),
                         self.getFourOfAKind(),
                         self.getFullHouse(),
                         isFlush,
                         self.getStraight(),
                         self.getTriple(),
                         len(self.getPairs()),
                         self.getPairs(),
                         self.cards
                         ]
    
    def getPairsOfList(self, cardList):
        return [pair[0] for pair in zip(cardList[:-1], cardList[1:]) if pair[0] == pair[1]]

    def getPairs(self):
        return self.getPairsOfList(self.cards)
    
    def getTriple(self):
        return self.getPairsOfList(self.getPairsOfList(self.cards))
    
    def getFourOfAKind(self):
        return self.getPairsOfList(self.getPairsOfList(self.getPairsOfList(self.cards)))

    def getStraight(self):
        return all([c[0] == c[1] + 1 for c in zip(self.cards[:-1], self.cards[1:]) ])
    
    def getFullHouse(self):
        if len(self.getPairs()) == 3 and len(self.getTriple()) == 1:
            return [self.getTriple(), self.getPairs()]

    def __cmp__(self, other):
        return reduce(lambda old, new: [old, cmp(*new)][old == 0], zip(self.features, other.features), 0)

import unittest
            
class Hands:
    HighCardHand = PokerHand("2H 3H 6H 7H 9D")
    OnePair = PokerHand("2H 2H 6H 7H 9D")
    OnePairWithBiggerPair = PokerHand("4H 4H 5H 7H 9D")
    OnePairOnDifferentPlace = PokerHand("2H 3H 3H 7H 9D")
    TwoPairs = PokerHand("2H 2H 3H 3H 9D")
    ThreeOfAKind = PokerHand("2H 2H 2H 3H 9D")
    Straight = PokerHand("2H 3H 4H 5H 6D")
    BiggerStraight = PokerHand("3H 4H 5H 6D 7D")
    Flush = PokerHand("2H 3H 6H 7H 9H")
    FullHouse = PokerHand("3H 3H 3H 4D 4D")
    FourOfAKind = PokerHand("3H 3H 3H 3D 4D")
    StraightFlush = PokerHand("3H 4H 5H 6H 7H")

class TestPokerHand(unittest.TestCase):
    def testSameHighCardHandsShouldBeEqual(self):
        self.assertEqual(Hands.HighCardHand, Hands.HighCardHand)

    def testHighCardHandWithHigherHighestCardShouldWin(self):
        self.assertGreater(PokerHand("2H 3D 5S 7C 9D"), PokerHand("2H 3D 5S 7C 8D"))

    def testHighCardHandWithLowerHighestCardShouldLose(self):
        self.assertLess(PokerHand("2H 3D 5S 7C 8D"), PokerHand("2H 3D 5S 7C 9D"))

    def testHighCardHandWithLowerSecondCardShouldLose(self):
        self.assertLess(PokerHand("2H 3D 5S 6C 9D"), PokerHand("2H 3D 5S 7C 9D"))

    def testHighCardHandsShouldIgnoreLowerCardsWhenHigherCardsAreDifferent(self):
        self.assertLess(PokerHand("2H 3D 6S 7C 9D"), PokerHand("2H 3D 5S 8C 9D"))

    def testHighCardHandsShouldIgnoreSuits(self):
        self.assertEqual(Hands.HighCardHand, PokerHand("2D 3D 6D 7D 9H"))
        
    def testOnePairShouldWinHighCard(self):
        self.assertGreater(Hands.OnePair, Hands.HighCardHand)

    def testHighCardShouldLoseToOnePair(self):
        self.assertLess(Hands.HighCardHand, Hands.OnePair)

    def testHighCardWithHigherPairShouldWin(self):
        self.assertGreater(Hands.OnePairWithBiggerPair, Hands.OnePair)

    def testOnePairWithPairOnDifferentPlaces(self):
        self.assertGreater(Hands.OnePairOnDifferentPlace, Hands.HighCardHand)

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
    #import sys;sys.argv = ['', 'Test.testName']
    unittest.main()
