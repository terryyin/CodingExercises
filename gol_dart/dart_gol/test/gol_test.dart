import 'package:test/test.dart';
import 'package:dart_gol/game_of_life.dart' show GameState;

void main() {
  group('Game Of Life', () {
    test('A life with no neighbour must die', () {
      expect(GameState.withLifeAt(2, 3).nextState().alive(), false);
    });
  });
}
