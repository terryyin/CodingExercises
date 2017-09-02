import React, { Component } from 'react';

class GameOfLifeApp extends Component {
  constructor(props) {
    super(props);
    this.state = {game: props.game};
    this.state.game.randomize();
  }

  componentDidMount() {
    this.timerID = setInterval(
      () => { this.setState({game: this.state.game.next()});},
        300
    );
  }

  componentWillUnmount() {
    clearInterval(this.timerID);
  }

  render() {
    return (
      <div className="GameOfLifeApp">
          <h2>Welcome to the game of life</h2>
          <h2>{this.state.game.cells.length}</h2>
      {[...Array(50)].map((x, i) =>
        <div className="gol-row">
        {[...Array(50)].map((y, j) =>
          <div className={"gol-cell" + (this.state.game.isAliveAt(i, j) ? " alive" : '') }>
          </div>
        )}
        </div>
      )}
      </div>
    );
  }
}

export default GameOfLifeApp;

