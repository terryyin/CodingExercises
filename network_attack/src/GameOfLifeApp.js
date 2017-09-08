import React, { Component } from 'react';
import GameOfLife from './GameOfLifeCells'

class GameOfLifeApp extends Component {
  constructor(props) {
    super(props);
    this.state = {game: GameOfLife.randomGame(props)};
  }

  componentDidMount() {
    this.timerID = setInterval(
      () => { this.setState({game: this.state.game.next()});},
        100
    );
  }

  componentWillUnmount() {
    clearInterval(this.timerID);
  }

  render() {
    return (
      <div className="GameOfLifeApp">
          <h2>Welcome to the game of life</h2>
          <h2>{this.state.game.aliveCells.size}</h2>
      {[...Array(this.props.rows)].map((x, i) =>
        <div className="gol-row">
        {[...Array(this.props.cols)].map((y, j) =>
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

