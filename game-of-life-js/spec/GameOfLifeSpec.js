import GameOfLife from '../src/GameOfLifeCells'

describe('In the Game of life, ', ()=> {
  var the_cell = "8, 2";
  var neighbours = GameOfLife.neighboursOf(the_cell);;


  describe('An alive cell', ()=> {
    [
      [0, false],
      [1, false],
      [2, true],
      [3, true],
      [4, false]
    ].forEach((param)=> {
      it('should '+(param[1] ? '' : 'not') +' be alive with ' + param[0]+' neighbours', ()=> {
        let game = new GameOfLife(neighbours.slice(0, param[0]).concat([the_cell]));
        expect(game.next().isAlive(the_cell)).toBe(param[1]);
      });
    });
  });

  describe('A dead cell', ()=> {
    it('should revive with 3 neighbours', ()=> {
      let game = new GameOfLife(neighbours.slice(0, 3));
      expect(game.next().isAlive(the_cell)).toBe(true);
    });
  });

  describe('#randomize', ()=> {
    it('should set random place to be alive', ()=> {
      spyOn(Math, 'random').and.returnValue(0.04444);;
      let game = GameOfLife.randomGame({rows: 50, cols: 50, seeds: 1});
      expect(game.isAliveAt(2, 2)).toBe(true);
    });
  });

  describe('no duplicated cells', ()=> {
    it('should not keep duplicated cell', ()=> {
      let game = new GameOfLife(['1, 2', '1, 2']);
      expect(game.aliveCells.size).toBe(1);
    });
  });

  describe('A cell', ()=> {
    var cell = '3, 2';
    it('has neighbours', ()=> {
      expect(GameOfLife.neighboursOf(cell)[4]).toEqual('4, 3');
      expect(GameOfLife.neighboursOf(cell)[5]).toEqual('2, 3');
      expect(GameOfLife.neighboursOf(cell)[6]).toEqual('4, 1');
      expect(GameOfLife.neighboursOf(cell)[7]).toEqual('2, 1');
    });
})

});

