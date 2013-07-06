import unittest

class Book:
    pass

def costOfBook(*books):
    if len(books) == 0:
        return 0
    bookSet = set(books)
    cost = [0, 8, 7.6, 7.2, 6.4, 6][len(bookSet)]
    
    return cost * len(bookSet) + costOfBook(*restBooks(books, bookSet))

def restBooks(books, bookSet):
    bookList = list(books)
    for book in bookSet:
        bookList.remove(book)
    return bookList

class Test(unittest.TestCase):

    def testOneBookCost8Dollors(self):
        self.assertEqual(8, costOfBook(Book()))

    def testTwoSameBooksHaveNoDiscount(self):
        book = Book()
        self.assertEqual(8 * 2, costOfBook(book, book))

    def testTwoDifferentBooksHas5PercentDiscount(self):
        self.assertEqual(2 * 8 * 0.95, costOfBook(Book(), Book()))

    def testThreeDifferentBooksHas10PercentDiscount(self):
        self.assertEqual(3 * 8 * 0.90, costOfBook(Book(), Book(), Book()))

    def testFourDifferentBooksHas20PercentDiscount(self):
        self.assertEqual(4 * 8 * 0.80, costOfBook(Book(), Book(), Book(), Book()))

    def testFiveDifferentBooksHas25PercentDiscount(self):
        self.assertEqual(5 * 8 * 0.75, costOfBook(Book(), Book(), Book(), Book(), Book()))

    def testTwoSameBooksAndOneDifferentBookThenOneOfTheSameBookHasNoDiscount(self):
        book = Book()
        self.assertEqual(8 * 2 * 0.95 + 8, costOfBook(book, Book(), book))

    def testTwoSetOfTwoDifferentBooksHas5PercentDiscount(self):
        book1 = Book()
        book2 = Book()
        self.assertEqual(8 * 4 * 0.95, costOfBook(book1, book1, book2, book2))



if __name__ == "__main__":
    #import sys;sys.argv = ['', 'Test.testName']
    unittest.main()