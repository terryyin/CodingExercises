class GameState {
  GameState.withLifeAt(int i, int j) {}

  GameState nextState() {
    return this;
  }

  bool alive() {
    return true;
  }
}
