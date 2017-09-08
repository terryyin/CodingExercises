import { Set } from './Set'

function GameOfLife(cells) {
  this.aliveCells = new Set(cells || []);
}

GameOfLife.randomGame = function(props){
  return new GameOfLife(
    [...Array(props.seeds)].map(()=>
      [props.rows, props.cols].map((n)=> Math.floor(Math.random() * n)).join(', ')
    ));
};

GameOfLife.prototype.isAliveAt = function(x, y){
  return this.isAlive([x, y].join(', '));
};

GameOfLife.prototype.isAlive = function(cell){
  return this.aliveCells.has(cell);
};

GameOfLife.prototype.next = function(){
  return new GameOfLife(
    this.cellsWithNeighbourRange(this.aliveCells,      [2, 3]).concat(
    this.cellsWithNeighbourRange(this.allNeighbours(), [3])));
};

GameOfLife.prototype.allNeighbours = function(){
  return this.aliveCells.reduce((a, c)=> a.union(GameOfLife.neighboursOf(c).filter((x)=>!this.isAlive(x))), new Set());
}

GameOfLife.prototype.cellsWithNeighbourRange = function(cells, range){
  return cells.filter((c)=> range.includes(this.numberOfAliveNeighbours(c)));
}

GameOfLife.prototype.numberOfAliveNeighbours = function(cell){
  return this.aliveCells.intersectionCount(GameOfLife.neighboursOf(cell));
}

GameOfLife.neighboursCache = {}
GameOfLife.neighboursOf = function(cell){
  if(! GameOfLife.neighboursCache[cell]) {
    let [x, y] = cell.split(', ').map((x)=>parseInt(x));
    GameOfLife.neighboursCache[cell] = [[-1, 0], [1, 0], [0, 1], [0, -1], [1, 1], [-1, 1], [1, -1], [-1, -1]
      ].map((d)=> [x + d[0], y + d[1]].join(', '));
  }
  return GameOfLife.neighboursCache[cell];
};

export default GameOfLife;
