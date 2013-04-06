import unittest
from itertools import chain

def _numberOfAliveNeighbors(aliveCells, cell):
    return len(aliveCells.intersection(neighborsOf(cell)))

def tick(aliveCells):
    allNeighborsOfAliveCells = chain(*[neighborsOf(cell) for cell in aliveCells])
    return [cell for cell in aliveCells               if _numberOfAliveNeighbors(aliveCells, cell) == 2] + \
           [cell for cell in allNeighborsOfAliveCells if _numberOfAliveNeighbors(aliveCells, cell) == 3]

def neighborsOf(cell):
    return [(cell[0] + dx, cell[1] + dy) for dx, dy in ((-1, 0), (1, 0), (0, -1), (0, 1), (-1, -1), (-1, 1), (1, -1), (1, 1))]

class TestCell(unittest.TestCase):
    def test_neighbors_of_cell(self):
        cell = x, y = (2, 3)
        self.assertIn((x - 1, y), neighborsOf(cell))
        self.assertIn((x + 1, y), neighborsOf(cell))
        self.assertIn((x, y - 1), neighborsOf(cell))
        self.assertIn((x, y + 1), neighborsOf(cell))
        self.assertIn((x - 1, y - 1), neighborsOf(cell))
        self.assertIn((x - 1, y + 1), neighborsOf(cell))
        self.assertIn((x + 1, y - 1), neighborsOf(cell))
        self.assertIn((x + 1, y + 1), neighborsOf(cell))

class TestTick(unittest.TestCase):
    cell = (2, 4)
    def test_alive_cell_should_be_dead_when_having_0_neighbor(self):
        self.assertNotIn(self.cell, tick(set([self.cell])))
        
    def test_alive_cell_should_be_alive_when_having_3_neighbors(self):
        aliveCells = set([self.cell] + neighborsOf(self.cell)[0:3])
        self.assertIn(self.cell, tick(aliveCells))
 
    def test_two_alive_cells_should_both_be_alive_when_both_having_2_neighbors(self):
        anotherCell = (100, 200)
        aliveCells = set([self.cell] + neighborsOf(self.cell)[0:2] + [anotherCell] + neighborsOf(anotherCell)[0:2])
        self.assertIn(anotherCell, tick(aliveCells))

    def test_dead_cell_should_be_alive_when_having_3_neighbors(self):
        aliveCells = set(neighborsOf(self.cell)[0:3])
        self.assertIn(self.cell, tick(aliveCells))
