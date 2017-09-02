import React from 'react';
import ReactDOM from 'react-dom';
import GameOfLifeApp from '../src/GameOfLifeApp';
import GameOfLife from '../src/GameOfLifeCells';
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
    props = {seeds: 3, rows: 50, cols: 30};
  });

  it ("has 50 rows, 30 cols", ()=> {
    ReactDOM.render(<GameOfLifeApp {...props}/>, div);
    expect(div.querySelectorAll('.gol-row').length).toEqual(50);
    const first_row = div.querySelectorAll('.gol-row')[0];
    expect(first_row.querySelectorAll('.gol-cell').length).toEqual(30);
  });

  it ("render the lives", ()=> {
    let cell = '2, 3'
    spyOn(GameOfLife, 'randomGame').and.returnValue(new GameOfLife([cell]));
    ReactDOM.render(<GameOfLifeApp {...props}/>, div);
    const row2 = div.querySelectorAll('.gol-row')[2].querySelectorAll('.gol-cell');
    expect(row2[3].className.split(' ')).toContain('alive');
    expect(row2[4].className.split(' ')).not.toContain('alive');
  });

});
