import unittest

def fizzBuzz(n):
    return 'fizzbuzz' if n % 15 == 0 else 'buzz' if n % 5 == 0 else 'fizz' if n % 3 == 0 else n


class TestFizzBuzz(unittest.TestCase):
    
    def test_fizz_buzz(self):
        self.assertEqual(1, fizzBuzz(1))
        self.assertEqual('fizz', fizzBuzz(3))
        self.assertEqual('buzz', fizzBuzz(5))
        self.assertEqual('fizz', fizzBuzz(6))
        self.assertEqual('buzz', fizzBuzz(10))
        self.assertEqual('fizzbuzz', fizzBuzz(15))
