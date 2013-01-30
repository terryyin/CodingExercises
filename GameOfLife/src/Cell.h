#ifndef CELL_H_
#define CELL_H_
#include <vector>
class Cell {
public:
	Cell(int x, int y);
	bool operator== (const Cell& other) const;
	bool operator< (const Cell& other) const;
	std::vector<Cell> neighbors() const;
private:
	int x_;
	int y_;
};

#endif /* CELL_H_ */
