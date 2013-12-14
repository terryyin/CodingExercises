'''
Convey's Game Of Life exercise.
Restriction applied: no local variable
Result: somehow failed, have to use iterators at least
'''

import unittest, itertools

def nextGeneration(liveCells):

    neighborsOf = lambda cell : [(cell[0] + dx, cell[1] + dy) for (dx, dy) in ((-1, 0), (1, 0), (0, -1), (0, 1), (-1, -1), (1, -1), (-1, 1), (1, 1))]
    liveNeighborCountOf = lambda cell : len(set(liveCells).intersection(neighborsOf(cell)))
    allRelatedCells = lambda : itertools.chain(*[neighborsOf(liveCell) for liveCell in liveCells])

    return [cell for cell in allRelatedCells() 
                    if liveNeighborCountOf(cell) is 3 
                    or (liveNeighborCountOf(cell) is 2 and cell in liveCells)]
            
class TestGameOfLife(unittest.TestCase):
    
    cell = (2, 3)
    neighbors = [(1, 3), (3, 3), (2, 2), (2, 4), (1, 2), (3, 2), (1, 4), (3, 4)]
    
    def test_CellWithNoNeighborShouldDie(self):
        self.assertNotIn(self.cell, nextGeneration([self.cell]))

    def test_CellWith2NeighborsShouldSurvive(self):
        self.assertIn(self.cell, nextGeneration([self.cell] + self.neighbors[0:2]))

    def test_CellWith3NeighborsShouldSurvive(self):
        self.assertIn(self.cell, nextGeneration([self.cell] + self.neighbors[0:3]))

    def test_CellWithMoreThan4NeighborsShouldDie(self):
        self.assertNotIn(self.cell, nextGeneration([self.cell] + self.neighbors[0:4]))

    def test_DeadCellWith3NeighborsShouldComeBackToLife(self):
        self.assertIn(self.cell, nextGeneration(self.neighbors[0:3]))

