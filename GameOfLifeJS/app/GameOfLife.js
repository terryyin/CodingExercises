function GameOfLife(cells) {
  this.cells = cells || new Array();
}

GameOfLife.prototype.setAlive = function(cell){
  this.cells.push(cell);
};

GameOfLife.prototype.isAlive = function(cell){
  return this.cells.some((c)=>{ return c.equals(cell);});
};

GameOfLife.prototype.next = function(){
  return new GameOfLife(this.cells.filter((c)=>{
    return [2, 3].includes(Object.values(c.neighbours()).filter((n)=>{
      return this.isAlive(n)}).length);
  }));
};

function Cell(x, y) {
  this.x = x;
  this.y = y;
}

Cell.prototype.toString = function(){
  return '(' + [this.x, this.y].join(', ') + ')';
};

Cell.prototype.neighbours = function(){
  return {
    right: new Cell(this.x - 1, this.y),
    left: new Cell(this.x + 1, this.y),
    up: new Cell(this.x, this.y + 1),
    down: new Cell(this.x, this.y - 1)
  };
};

Cell.prototype.equals = function(other){
  return this.x === other.x && this.y === other.y;
};

module.exports = {
  GameOfLife: GameOfLife,
  Cell: Cell };
