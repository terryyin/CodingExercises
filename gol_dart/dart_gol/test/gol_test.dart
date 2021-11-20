import 'package:test/test.dart';
import 'package:dart_gol/game_of_life.dart' show GameState, Position;

void main() {
  group('Game Of Life', () {
    final position = Position();
    test('A life with no neighbour must die', () {
    });

    test('A life with two neighbours must survive', () {
      expect(GameState.withLivesAt([position, ...position.neighbours([0, 1])]).nextState().alive(position), true);
    });

  });
}
