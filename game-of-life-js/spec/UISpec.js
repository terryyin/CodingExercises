import React from 'react';
import ReactDOM from 'react-dom';
import GameOfLifeApp from '../src/GameOfLifeApp';
var GameOfLife = require('../src/GameOfLifeCells').GameOfLife;
var Cell = require('../src/GameOfLifeCells').Cell;
const jsdom = require("jsdom");
const { JSDOM } = jsdom;


describe('In the Game of life, ', ()=> {
  var div;
  var props;
  beforeEach(()=> {
    const dom = new JSDOM(`<!DOCTYPE html><p>Hello world</p>`);
    global.window = dom.window;
    global.document = dom.window.document;
    div = document.createElement('div');
    props = {game: new GameOfLife()};
  });

  it ("has 50 rows, 50 cols", ()=> {
    ReactDOM.render(<GameOfLifeApp {...props}/>, div);
    expect(div.querySelectorAll('.gol-row').length).toEqual(50);
    const first_row = div.querySelectorAll('.gol-row')[0];
    expect(first_row.querySelectorAll('.gol-cell').length).toEqual(50);
  });

  it ("render the lives", ()=> {
    let cell = new Cell(2, 3);
    let anothr_cell = new Cell(2, 4);
    props.game.setAlives([cell]);
    ReactDOM.render(<GameOfLifeApp {...props}/>, div);
    const row2 = div.querySelectorAll('.gol-row')[2].querySelectorAll('.gol-cell');
    expect(row2[3].className.split(' ')).toContain('alive');
    expect(row2[4].className.split(' ')).not.toContain('alive');
  });

  it ("should randomize the game", ()=> {
    spyOn(props.game, 'randomize');
    ReactDOM.render(<GameOfLifeApp {...props}/>, div);
    expect(props.game.randomize).toHaveBeenCalled();
  });

});
