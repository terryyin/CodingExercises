#ifndef GAMEOFLIFE_H_
#define GAMEOFLIFE_H_
#include <set>

class Cell;
class GameOfLife {
public:
	bool isAlive(const Cell& cell) const;
	GameOfLife& withAlive(const Cell& cell);
	GameOfLife tick() const;
private:
	int getNeighborCountOf(const Cell& cell) const;
	void tickOfAliveCell(const Cell& cell, GameOfLife& gameOfNextTick) const;
	bool has2Or3AliveNeighbors(const Cell& cell) const;
	void tickOfNeighborsAroundAliveCell(const Cell& cell, GameOfLife& gameOfNextTick) const;
	std::set<Cell> aliveCells_;
};

#endif /* GAMEOFLIFE_H_ */
