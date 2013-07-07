import unittest
from itertools import combinations

class Book:
    pass

def costOfBooks(*books):
    if len(books) == 0:
        return 0
    return min(costsOfPossibleCombinations(books))

def costsOfPossibleCombinations(books):
    bookSet = set(books)
    for n in range(1, len(bookSet)+1): 
        for subSet in combinations(bookSet, n) :             
            yield [0, 8, 7.6 * 2, 7.2 * 3, 6.4 * 4, 6 * 5][n]\
                + costOfBooks(*restBooks(books, subSet))

def restBooks(books, bookSet):
    bookList = list(books)
    for book in bookSet:
        bookList.remove(book)
    return bookList

class Test(unittest.TestCase):

    def testOneBookCost8Dollors(self):
        self.assertEqual(8, costOfBooks(Book()))

    def testTwoSameBooksHaveNoDiscount(self):
        book = Book()
        self.assertEqual(8 * 2, costOfBooks(book, book))

    def testTwoDifferentBooksHas5PercentDiscount(self):
        self.assertEqual(2 * 8 * 0.95, costOfBooks(Book(), Book()))

    def testThreeDifferentBooksHas10PercentDiscount(self):
        self.assertEqual(3 * 8 * 0.90, costOfBooks(Book(), Book(), Book()))

    def testFourDifferentBooksHas20PercentDiscount(self):
        self.assertEqual(4 * 8 * 0.80, costOfBooks(Book(), Book(), Book(), Book()))

    def testFiveDifferentBooksHas25PercentDiscount(self):
        self.assertEqual(5 * 8 * 0.75, costOfBooks(Book(), Book(), Book(), Book(), Book()))

    def testTwoSameBooksAndOneDifferentBookThenOneOfTheSameBookHasNoDiscount(self):
        book = Book()
        self.assertEqual(8 * 2 * 0.95 + 8, costOfBooks(book, Book(), book))

    def testTwoSetOfTwoDifferentBooksHas5PercentDiscount(self):
        book1 = Book()
        book2 = Book()
        self.assertEqual(8 * 4 * 0.95, costOfBooks(book1, book1, book2, book2))

    def testUseBiggestDiscount(self):
        book1 = Book()
        book2 = Book()
        book3 = Book()
        book4 = Book()
        book5 = Book()
        self.assertEqual(2 * 4 * 6.4, costOfBooks(book1, book1, book2, book2, book3, book3, book4, book5))



if __name__ == "__main__":
    #import sys;sys.argv = ['', 'Test.testName']
    unittest.main()