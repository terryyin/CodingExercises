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

describe('An alive cell', ()=> {
  var the_cell = new Cell(3, 2);
  var game;

  beforeEach(()=> {
    game = new GameOfLife();
    game.setAlive(the_cell);
  });

  it('should be alive if set alive', ()=> {
    expect(game).not.toBeDeadAt(new Cell(3, 2));
  });

  it('should be killed when no neighbour', ()=> {
    expect(game.next()).toBeDeadAt(the_cell);
  });

  it('should survive with two live neighbhours', ()=> {
    game.setAlive(the_cell.neighbours().left);
    game.setAlive(the_cell.neighbours().right);
    expect(game.next()).not.toBeDeadAt(the_cell);
    expect(game.next()).toBeDeadAt(the_cell.neighbours().right);
    expect(game.next()).toBeDeadAt(the_cell.neighbours().left);
  });

  it('should survive with three live neighbhours', ()=> {
    game.setAlive(the_cell.neighbours().left);
    game.setAlive(the_cell.neighbours().right);
    game.setAlive(the_cell.neighbours().up);
    expect(game.next()).not.toBeDeadAt(the_cell);
  });

  it('should be killed with four live neighbhours', ()=> {
    game.setAlive(the_cell.neighbours().left);
    game.setAlive(the_cell.neighbours().right);
    game.setAlive(the_cell.neighbours().up);
    game.setAlive(the_cell.neighbours().down);
    expect(game.next()).toBeDeadAt(the_cell);
  });
});

describe('A cell', ()=> {
  var cell = new Cell(3, 2);

  xit('s neighours should include the one to its left', ()=> {
    var cell = new Cell(3, 2);
    expect(cell.neighbours()).toInclude(new Cell(2, 2));
  });

});
