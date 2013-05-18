'''
Programming Kata LCD
-------------------------------
Problems to solve:
    Your task is to create an LCD string representation of an
    integer value using a 3x3 grid of space, underscore, and 
    pipe characters for each digit. Each digit is shown below 
    (using a dot instead of a space)

    Example: 910

    ._. ... ._.
    |_| ..| |.|
    ..| ..| |_|
'''
digit_displays = '''
    ._.   ...   ._.   ._.   ...   ._.   ._.   ._.   ._.   ._.
    |.|   ..|   ._|   ._|   |_|   |_.   |_.   ..|   |_|   |_|
    |_|   ..|   |_.   ._|   ..|   ._|   |_|   ..|   |_|   ..|

'''

class LCD:
    def __init__(self, digit_displays):
        self.digit_lines = [line.strip().split() for line in digit_displays.splitlines()[1:4]]

    def _get_line(self, line_display, n):
        return ' '.join(line_display[int(d)] for d in str(n))

    def _get_lines(self, n):
        return [self._get_line(line, n) for line in self.digit_lines]

    def display(self, n):
        return '\n'.join(self._get_lines(n)).replace('.', ' ')

import unittest

class TestDigitLCD(unittest.TestCase):

    def test_LCD_display_of_0(self):
        self.assertEqual(
                ' _ \n' +
                '| |\n' +
                '|_|',
                LCD(digit_displays).display(0))

    def test_LCD_display_of_1(self):
        self.assertEqual(
                '   \n' +
                '  |\n' +
                '  |',
                LCD(digit_displays).display(1))

    def test_LCD_display_multiple_digits(self):
        self.assertEqual(
                '     _ \n' +
                '  | |_|\n' +
                '  |   |',
                LCD(digit_displays).display(19))

