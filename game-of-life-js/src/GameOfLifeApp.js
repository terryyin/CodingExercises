import React, { Component } from 'react';

class GameOfLifeApp extends Component {
  render() {
    return (
      <div className="GameOfLifeApp">
        <div className="App-header">
          <h2>Welcome to React</h2>
        </div>
        <p className="App-intro">
          To get started, edit <code>src/App.js</code> and save to reload.
        </p>
      {[...Array(50)].map((x, i) =>
        <div className="gol-row">
        {[...Array(50)].map((x, i) =>
          <div className="gol-col">
          <h2>Welcome to React</h2>
          </div>
        )}
        </div>
      )}
      </div>
    );
  }
}

export default GameOfLifeApp;

