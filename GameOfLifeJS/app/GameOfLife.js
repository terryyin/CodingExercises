function GameOfLife(cells) {
  this.cells = cells || new Array();
}

function Cell() {
}

GameOfLife.prototype.setAlive = function(cell){
  this.cells.push(cell);
};

GameOfLife.prototype.isAlive = function(){
  return this.cells.length > 1;
};

GameOfLife.prototype.next = function(){
  return new GameOfLife(this.cells);
};

Cell.prototype.neighbours = function(){
  return ;
};


module.exports = {
  GameOfLife: GameOfLife,
  Cell: Cell };
