#include "CppUTest/TestHarness.h"
#include "Cell.h"

//todo: add operator<

TEST_GROUP(Cell) {
	bool cellIsANeighborOf(const Cell& neighbor, const Cell& cell) {
		for (int i = 0; i < cell.neighbors().size(); i++)
			if (cell.neighbors()[i] == neighbor)
				return true;

		return false;
	}
};

TEST(Cell, cellsShouldBeEqualWhenHavingSameXandY) {
	CHECK(Cell(1, 2) == Cell(1, 2));
}

TEST(Cell, cellsShouldNotBeEqualWhenHavingDifferentXOrY) {
	CHECK_FALSE(Cell(1, 2) == Cell(2, 2));
	CHECK_FALSE(Cell(1, 2) == Cell(1, 1));
}

TEST(Cell, cellComparisonShouldCompareXFirst) {
	CHECK(Cell(1, 2) < Cell(2, 2));
	CHECK_FALSE(Cell(2, 1) < Cell(1, 2));
}

TEST(Cell, cellComparisonShouldCompareYWhenXsAreTheSame) {
	CHECK(Cell(1, 2) < Cell(1, 3));
}

TEST(Cell, neighborsOfCell) {
	int x = 5, y = 10;
	Cell cell(x, y);
	CHECK(cellIsANeighborOf(Cell(x-1, y), cell));
	CHECK(cellIsANeighborOf(Cell(x+1, y), cell));
	CHECK(cellIsANeighborOf(Cell(x, y-1), cell));
	CHECK(cellIsANeighborOf(Cell(x, y+1), cell));
	CHECK(cellIsANeighborOf(Cell(x-1, y-1), cell));
	CHECK(cellIsANeighborOf(Cell(x+1, y-1), cell));
	CHECK(cellIsANeighborOf(Cell(x-1, y+1), cell));
	CHECK(cellIsANeighborOf(Cell(x+1, y+1), cell));
}

TEST(Cell, cellShouldHave8Neighbors) {
	int x = 5, y = 10;
	Cell cell(x, y);
	LONGS_EQUAL(8, cell.neighbors().size());
}


