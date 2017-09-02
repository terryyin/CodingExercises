import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './App';
import registerServiceWorker from './registerServiceWorker';
var GameOfLife = require('./GameOfLifeCells').GameOfLife;
var Cell = require('./GameOfLifeCells').Cell;

let props = {game: new GameOfLife([new Cell(3, 4)])};
ReactDOM.render(<App {...props}/>, document.getElementById('root'));
registerServiceWorker();
