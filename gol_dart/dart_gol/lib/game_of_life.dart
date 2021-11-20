class GameState {
  GameState.withLifeAt(Position pos) {}

  GameState nextState() {
    return this;
  }

  bool alive(Position pos) {
    return false;
  }
}

class Position {
  Position();
}
