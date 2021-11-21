class ALife {
  late final bool result;

  ALife.having(Neighbours neighbours) {
    result = [2, 3].contains(neighbours.count);
  }

  survive() {
    return result;
  }
}

class Neighbours {
  final int count;
  Neighbours(this.count);
}

extension NeighbourExtension on int {
  Neighbours get neighbours {
    return Neighbours(this);
  }
}
