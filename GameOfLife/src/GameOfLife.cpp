#include "Cell.h"
#include "GameOfLife.h"

using namespace std;

bool GameOfLife::isAlive(const Cell& cell) const {
	return aliveCells_.count(cell) > 0;
}

GameOfLife& GameOfLife::withAlive(const Cell& cell) {
	aliveCells_.insert(cell);
	return *this;
}

int GameOfLife::getNeighborCountOf(const Cell& cell) const {
	int aliveNeighborCount = 0;
	for (int i = 0; i < cell.neighbors().size(); i++)
		if (isAlive(cell.neighbors()[i]))
			aliveNeighborCount++;
	return aliveNeighborCount;
}

bool GameOfLife::has2Or3AliveNeighbors(const Cell& cell) const{
	int aliveNeighborCount = getNeighborCountOf(cell);
	return (aliveNeighborCount == 2 || aliveNeighborCount == 3);
}

void GameOfLife::tickOfAliveCell(const Cell& cell, GameOfLife& gameOfNextTick) const {
	if (has2Or3AliveNeighbors(cell))
		gameOfNextTick.withAlive(cell);
}

void GameOfLife::tickOfNeighborsAroundAliveCell(const Cell& cell, GameOfLife& gameOfNextTick) const {
	for (int i = 0; i < cell.neighbors().size(); i++)
		if (getNeighborCountOf(cell.neighbors()[i]) == 3)
			gameOfNextTick.withAlive(cell.neighbors()[i]);
}

GameOfLife GameOfLife::tick() const {
	GameOfLife gameOfNextTick;
	for(set<Cell>::iterator it = aliveCells_.begin(); it != aliveCells_.end(); it++) {
		tickOfAliveCell(*it, gameOfNextTick);
		tickOfNeighborsAroundAliveCell(*it, gameOfNextTick);
	}

	return gameOfNextTick;
}
