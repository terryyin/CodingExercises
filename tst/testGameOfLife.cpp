#include "CppUTest/TestHarness.h"
#include "Cell.h"
#include "GameOfLife.h"

static const Cell cell(7, 8);
static const Cell anotherCell(1, 2);

TEST_GROUP(GameOfLife) {
	GameOfLife game;
};

TEST(GameOfLife, noCellIsAliveBeforeSet) {
	CHECK_FALSE(game.isAlive(cell));
}

TEST(GameOfLife, cellShouldBeAliveWhenSetAlive) {
	game.withAlive(cell);
	CHECK(game.isAlive(cell));
}

TEST(GameOfLife, cellShouldBeDeadWhenSetAnotherCellAlive) {
	game.withAlive(anotherCell);
	CHECK_FALSE(game.isAlive(cell));
}

TEST(GameOfLife, bothShouldBeAliveWhenSetTwoCellsAlive) {
	game.withAlive(cell).withAlive(anotherCell);
	CHECK(game.isAlive(cell));
	CHECK(game.isAlive(anotherCell));
}

TEST(GameOfLife, gameOfDeadCellShouldStillBeDeadInNextTick) {
	CHECK_FALSE(game.tick().isAlive(cell));
}

TEST(GameOfLife, aliveCellShouldBeDeadInNextTickWhenHavingLessThen2AliveNeighbors) {
	game.withAlive(cell);
	CHECK_FALSE(game.tick().isAlive(cell));
}

TEST(GameOfLife, aliveCellShouldStillBeAliveInNextTickWhenHaving2Or3AliveNeighbors) {
	game.withAlive(cell).withAlive(cell.neighbors()[0]).withAlive(
			cell.neighbors()[1]);
	CHECK(game.tick().isAlive(cell));
	game.withAlive(cell.neighbors()[2]);
	CHECK(game.tick().isAlive(cell));
}

TEST(GameOfLife, deadCellShouldStillBeDeadInNextTickWhenHaving2AliveNeighbors) {
	game.withAlive(cell.neighbors()[0]).withAlive(cell.neighbors()[1]);
	CHECK_FALSE(game.tick().isAlive(cell));
}

TEST(GameOfLife, aliveCellShouldBeDeadInNextTickWhenHavingMoreThan3AliveNeighbors) {
	game.withAlive(cell).withAlive(cell.neighbors()[0]).withAlive(
			cell.neighbors()[1]).withAlive(cell.neighbors()[2]).withAlive(
			cell.neighbors()[3]);
	CHECK_FALSE(game.tick().isAlive(cell));
}

TEST(GameOfLife, deadCellShouldBecomeAliveInNextTickWhenHavingExactly3AliveNeighbors) {
	game.withAlive(cell.neighbors()[0]).withAlive(cell.neighbors()[1]).withAlive(
			cell.neighbors()[2]);
	CHECK(game.tick().isAlive(cell));
}

