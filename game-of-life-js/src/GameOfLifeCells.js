Set.prototype.union = function(setB) {
  var union = new Set(this);
  for (var elem of setB) {
    union.add(elem);
  }
  return union;
}

function GameOfLife(cells) {
  this.aliveCells = cells || new Array();
  this.aliveCells = new Set(this.aliveCells);//.filter((v, i, self)=>self.indexOf(v)===i);
}

GameOfLife.prototype.randomize = function(){
  this.aliveCells = new Set(
    [...Array(3080)].map(()=>
       new Cell(
        Math.floor(Math.random() * 200 + 25),
        Math.floor(Math.random() * 200 + 25)).toString()
    ));
};

GameOfLife.prototype.setAlives = function(cells){
  this.aliveCells = this.aliveCells.union(cells.map((x)=>x.toString()));
};

GameOfLife.prototype.isAliveAt = function(x, y){
  return this.isAlive(new Cell(x, y));
};

GameOfLife.prototype.isAlive = function(cell){
  return this.aliveCells.has(cell.toString());
};

GameOfLife.prototype.next = function(){
  return new GameOfLife(
    this.cellsWithNeighbourRange(Array.from(this.aliveCells.values()), [2, 3]).concat(
    this.cellsWithNeighbourRange(this.allNeighbours(), [3])));
};

GameOfLife.prototype.toC = function(s){
  return new Cell(...s.split(', ').map((x)=>parseInt(x)));
}

GameOfLife.prototype.allNeighbours = function(){
  return Array.from(this.aliveCells.values()).reduce((a, c)=>{
    return a.concat(this.toC(c).neighbours()); }, []);
}

GameOfLife.prototype.cellsWithNeighbourRange = function(cells, range){
  return cells.filter((c)=>{
    return range.includes(this.numberOfAliveNeighbours(this.toC(c))); });
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
  return [this.x, this.y].join(', ');
};

Cell.prototype.split = function(){
  return [this.x, this.y];
};

Cell.prototype.neighbours = function(){
  return [
    [-1, 0], [1, 0], [0, 1], [0, -1], [1, 1], [-1, 1], [1, -1], [-1, -1]
  ].map((d)=> {return [this.x + d[0], this.y + d[1]].join(', ') });
};

module.exports = { GameOfLife: GameOfLife, Cell: Cell };

