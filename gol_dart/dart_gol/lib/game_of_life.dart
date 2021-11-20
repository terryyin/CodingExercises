class GameState {
  final List<Position> lives;

  GameState.withLivesAt(this.lives);

  GameState nextState() {
    return this;
  }

  bool alive(Position pos) {
    return lives.length == 3;
  }
}

class Position {
  Position();

  neighbours(List<int> list) {
    return list.map((e) => Position());
  }
}
