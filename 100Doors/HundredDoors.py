from unittest import TestCase
from math import sqrt

def doorStateAt(index):
    return int(sqrt(index)) * sqrt(index)  == index

class TestHundredDoors(TestCase):
    
    def test_door1_is_open_after_the_1st_pass(self):
        self.assertTrue(doorStateAt(1))        # 1

    def test_door_with_primer_index_is_closed(self):
        self.assertFalse(doorStateAt(2))       # 1, 2

    def test_door_index_m_by_n_is_closed(self):
        self.assertFalse(doorStateAt(2 * 3))   # 1, 2, 3, 6

    def test_door_with_squre_index_is_open(self):
        self.assertTrue(doorStateAt(2 * 2))    # 1, 2, 4

if __name__ == '__main__':
    print ''.join(["_",str(i)][doorStateAt(i)] for i in range(1, 101))
