function GameOfLife(cells) {
  this.cells = cells || new Array();
}

GameOfLife.prototype.setAlives = function(cells){
  this.cells = this.cells.concat(cells);
};

GameOfLife.prototype.isAlive = function(cell){
  return this.cells.some((c)=>{ return c.equals(cell);});
};

GameOfLife.prototype.next = function(){
  return new GameOfLife(
    this.cellsWithNeighbourRange(this.cells, [2, 3]).concat(
    this.cellsWithNeighbourRange(this.allNeighbours(), [3])));
};

GameOfLife.prototype.allNeighbours = function(){
  return this.cells.reduce((a, c)=>{
    return a.concat(c.neighbours()); }, []);
}

GameOfLife.prototype.cellsWithNeighbourRange = function(cells, range){
  return cells.filter((c)=>{
    return range.includes(this.numberOfAliveNeighbours(c)); });
}

GameOfLife.prototype.numberOfAliveNeighbours = function(cell){
  return cell.neighbours().filter((n)=>{
    return this.isAlive(n) }).length;
}

function Cell(x, y) {
  this.x = x;
  this.y = y;
}

Cell.prototype.toString = function(){
  return '(' + [this.x, this.y].join(', ') + ')';
};

Cell.prototype.neighbours = function(){
  return [
    [-1, 0], [1, 0], [0, 1], [0, -1], [1, 1], [-1, 1], [1, -1], [-1, -1]
  ].map((d)=> {return  new Cell(this.x + d[0], this.y + d[1]) });
};

Cell.prototype.equals = function(other){
  return this.x === other.x && this.y === other.y;
};

module.exports = { GameOfLife: GameOfLife, Cell: Cell };

