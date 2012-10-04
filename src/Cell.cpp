#include "Cell.h"

using namespace std;
Cell::Cell(int x, int y) : x_(x), y_(y){
}

bool Cell::operator== (const Cell& other) const {
	return (x_ == other.x_) && (y_ == other.y_);
}

bool Cell::operator< (const Cell& other) const {
	if (x_ < other.x_)
		return true;
	if (x_ > other.x_)
		return false;
	return (y_ < other.y_);
}

vector<Cell> Cell::neighbors() const {
	vector<Cell> neighborCells;
	neighborCells.push_back(Cell(x_-1, y_));
	neighborCells.push_back(Cell(x_+1, y_));
	neighborCells.push_back(Cell(x_, y_-1));
	neighborCells.push_back(Cell(x_, y_+1));
	neighborCells.push_back(Cell(x_-1, y_-1));
	neighborCells.push_back(Cell(x_+1, y_-1));
	neighborCells.push_back(Cell(x_-1, y_+1));
	neighborCells.push_back(Cell(x_+1, y_+1));
	return neighborCells;
}
