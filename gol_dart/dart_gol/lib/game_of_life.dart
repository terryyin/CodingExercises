class ALife {
  late final bool result;
  ALife() {
    result = false;
  }

  ALife.having(Neighbours neighbours) {
    result = neighbours.count == 2;
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
