import 'package:test/test.dart';
import 'package:dart_gol/game_of_life.dart' show ALife, NeighbourExtension;

void main() {
  group('Game Of Life', () {
    test('', () {
      expect(ALife.having(0.neighbours).survive(), false);
    });

    test('', () {
      expect(ALife.having(1.neighbours).survive(), false);
    });

    test('', () {
      expect(ALife.having(2.neighbours).survive(), true);
    });

    test('', () {
      expect(ALife.having(3.neighbours).survive(), true);
    });

    test('', () {
      expect(ALife.having(4.neighbours).survive(), false);
    });
  });
}
