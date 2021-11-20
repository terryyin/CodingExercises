class GameState {
  GameState.withLifeAt(Position pos) {}

  GameState nextState() {
    return this;
  }

  bool alive() {
    return false;
  }
}

class Position {
  Position(int i, int j);
}
