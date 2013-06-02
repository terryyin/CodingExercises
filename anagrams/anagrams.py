import unittest
from itertools import permutations

def anagrams(base):
    return set(''.join(x) for x in permutations(base))

class TestAnagrams(unittest.TestCase):

    def test_Anagrams_Of_A_Single_Letter_Is_Itself(self):
        self.assertItemsEqual(['a'], anagrams('a'))

    def test_Anagrams_of_2_letters_are_ab_and_ba(self):
        self.assertItemsEqual(['ab', 'ba'], anagrams('ab'))

    def test_Anagrams_of_2_same_letters_is_aa(self):
        self.assertItemsEqual(['aa'], anagrams('aa'))

if __name__ == "__main__":
    #import sys;sys.argv = ['', 'Test.testName']
    unittest.main()