beforeEach( ()=> {
  jasmine.addMatchers({
    toBeDeadAt: ()=> {
      return {
        compare: function (game, cell) {
          return {
            pass: !game.isAlive(cell)
          };
        }
      };
    }
  });
});


var GameOfLife = require('../app/GameOfLife').GameOfLife;
var Cell = require('../app/GameOfLife').Cell;

describe('Game Of Life', ()=> {
  var the_cell = new Cell(3, 2);
  var game;

  beforeEach(()=> { game = new GameOfLife();});

  it('should kill a single live cell', ()=> {
    game.setAlive(the_cell);
    expect(game.next()).toBeDeadAt(the_cell);
  });

  it('should retain a live cell with two live neighbhours', ()=> {
    game.setAlive(the_cell);
    game.setAlive(new Cell(2, 2));
    game.setAlive(new Cell(3, 3));
    expect(game.next()).not.toBeDeadAt(the_cell);
  });

});

describe('A cell', ()=> {
  xit('s neighours should include the one to its left', ()=> {
    var cell = new Cell(3, 2);
    expect(cell.neighbours()).toInclude(new Cell(2, 2));
  });
});
