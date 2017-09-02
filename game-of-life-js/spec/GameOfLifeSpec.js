var GameOfLife = require('../src/GameOfLifeCells').GameOfLife;
var Cell = require('../src/GameOfLifeCells').Cell;

describe('In the Game of life, ', ()=> {
  var the_cell = new Cell(8, 2);
  var neighbours = Object.values(the_cell.neighbours());
  var game;

  beforeEach(()=> { game = new GameOfLife(); });

  describe('An alive cell', ()=> {
    beforeEach(()=> { game.setAlives([the_cell]); });

    [
      [0, false],
      [1, false],
      [2, true],
      [3, true],
      [4, false]
    ].forEach((param)=> {
      it('should '+(param[1] ? '' : 'not') +' be alive with ' + param[0]+' neighbours', ()=> {
        game.setAlives(neighbours.slice(0, param[0]));
        expect(game.next().isAlive(the_cell)).toBe(param[1]);
      });
    });
  });

  describe('A dead cell', ()=> {
    it('should revive with 3 neighbours', ()=> {
      game.setAlives(neighbours.slice(0, 3));
      expect(game.next().isAlive(the_cell)).toBe(true);
    });
  });

  describe('#randomize', ()=> {
    it('should set random place to be alive', ()=> {
      spyOn(Math, 'random').and.returnValue(0.04444);;
      game.randomize();
      expect(game.isAlive(new Cell(2, 2))).toBe(true);
    });
  });

  describe('no duplicated cells', ()=> {
    it('should not keep duplicated cell', ()=> {
      game = new GameOfLife([new Cell(1, 2), new Cell(1, 2)]);
      expect(game.cells.length).toBe(1);
    });
  });

});

describe('A cell', ()=> {
  var cell = new Cell(3, 2);

  it('has neighbours', ()=> {
    expect(cell.neighbours()[4].equals(new Cell(4, 3))).toBe(true);
    expect(cell.neighbours()[5].equals(new Cell(2, 3))).toBe(true);
    expect(cell.neighbours()[6].equals(new Cell(4, 1))).toBe(true);
    expect(cell.neighbours()[7].equals(new Cell(2, 1))).toBe(true);
  });
})

