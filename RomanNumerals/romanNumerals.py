import unittest

def toRoman(number):
    ones = ['', 'I', 'II', 'III', 'IV', 'V', 'VI', 'VII', 'VIII', 'IX']
    tens = ['', 'X', 'XX', 'XXX', 'XL', 'L', 'LX', 'LXX', 'LXXX', 'XC']
    hundreds = ['', 'C', 'CC', 'CCC', 'CD', 'D', 'DC', 'DCC', 'DCCC', 'CM']
    thousands = ['', 'M', 'MM', 'MMM']
    return thousands[number/1000%10] + hundreds[number/100%10] + tens[number/10%10] + ones[number/1%10]

class TestRomanNumerals(unittest.TestCase):
    
    def testRomanNumeralOnes(self):
        self.assertEqual("I", toRoman(1))
        self.assertEqual("II", toRoman(2))
        self.assertEqual("III", toRoman(3))
        self.assertEqual("IV", toRoman(4))
        self.assertEqual("V", toRoman(5))
        self.assertEqual("VI", toRoman(6))
        self.assertEqual("VII", toRoman(7))
        self.assertEqual("VIII", toRoman(8))
        self.assertEqual("IX", toRoman(9))
        
    def testRomanNumeralTens(self):
        self.assertEqual("X", toRoman(10))
        self.assertEqual("XX", toRoman(20))
        self.assertEqual("XXX", toRoman(30))
        self.assertEqual("XL", toRoman(40))
        self.assertEqual("L", toRoman(50))
        self.assertEqual("LX", toRoman(60))
        self.assertEqual("LXX", toRoman(70))
        self.assertEqual("LXXX", toRoman(80))
        self.assertEqual("XC", toRoman(90))
        
    def testRomanNumeralHundreds(self):
        self.assertEqual("C", toRoman(100))
        self.assertEqual("CC", toRoman(200))
        self.assertEqual("CCC", toRoman(300))
        self.assertEqual("CD", toRoman(400))
        self.assertEqual("D", toRoman(500))
        self.assertEqual("DC", toRoman(600))
        self.assertEqual("DCC", toRoman(700))
        self.assertEqual("DCCC", toRoman(800))
        self.assertEqual("CM", toRoman(900))

    def testRomanNumeralThousands(self):
        self.assertEqual("M", toRoman(1000))
        self.assertEqual("MM", toRoman(2000))
        self.assertEqual("MMM", toRoman(3000))
        
    def testRomanNumeralsShouldCombineTensAndOnes(self):
        self.assertEqual("XI", toRoman(11))
    
    def testRomanNumeralsShouldCombineHundredsAndTens(self):
        self.assertEqual("CX", toRoman(110))

    def testRomanNumeralsShouldCombineThousandAndTens(self):
        self.assertEqual("MX", toRoman(1010))


if __name__ == "__main__":
    unittest.main()