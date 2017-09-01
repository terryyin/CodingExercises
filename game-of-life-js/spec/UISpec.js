import React from 'react';
import ReactDOM from 'react-dom';
import GameOfLifeApp from '../src/GameOfLifeApp';
const jsdom = require("jsdom");
const { JSDOM } = jsdom;


describe('In the Game of life, ', ()=> {
  beforeEach(()=> {
    const dom = new JSDOM(`<!DOCTYPE html><p>Hello world</p>`);
    global.window = dom.window;
    global.document = dom.window.document;
  });

  it ("has 50 rows, 50 cols", ()=> {
    const div = document.createElement('div');
    ReactDOM.render(<GameOfLifeApp />, div);
    expect(div.querySelectorAll('.gol-row').length).toEqual(50);
    const first_row = div.querySelectorAll('.gol-row')[0];
    expect(first_row.querySelectorAll('.gol-col').length).toEqual(50);
  });
});
