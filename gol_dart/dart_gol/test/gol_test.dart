import 'package:test/test.dart';
import 'package:dart_gol/game_of_life.dart' show ALife, NeighbourExtension;

void main() {
  group('Game Of Life', () {
    test('', () {
      expect(ALife().survive(), false);
    });

    test('', () {
      expect(ALife.having(1.neighbours).survive(), false);
    });

    test('', () {
      expect(ALife.having(2.neighbours).survive(), true);
    });
  });
}
