function GameOfLife() {

}

GameOfLife.prototype.setAlive = ()=>{
};

GameOfLife.prototype.isAlive = ()=>{
  return false;
};

GameOfLife.prototype.next = ()=>{
  return new GameOfLife();
};

module.exports = GameOfLife;
