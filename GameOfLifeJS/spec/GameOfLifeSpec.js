beforeEach( ()=> {
  jasmine.addMatchers({
    toBeDeadAt: ()=> {
      return {
        compare: function (game, x, y) {
          return {
            pass: !game.isAlive(x, y)
          };
        }
      };
    }
  });
});


describe('Game Of Life', ()=> {

  var GameOfLife = require('../app/GameOfLife');

  it('should kill a single live cell', ()=> {
    var game = new GameOfLife();
    game.setAlive(3, 2);
    expect(game.next()).toBeDeadAt(3, 2);
  });

  xit('should retain a live cell with two live neighbhours', ()=> {
    var game = new GameOfLife();
    game.setAlive(3, 2);
    game.setAlive(2, 2);
    game.setAlive(3, 3);
    expect(game.next()).not.toBeDeadAt(3, 2);
  });

});
