Set.prototype.union = function(setB) {
  for (var elem of setB) {
    this.add(elem);
  }
  return this;
}

Set.prototype.intersectionCount = function(setB) {
  var intersection = 0;
  for (var elem of setB) {
    if (this.has(elem)) {
      intersection ++;
    }
  }
  return intersection;
}

Set.prototype.filter = function(f) {
  return Array.from(this).filter(f);
}

Set.prototype.reduce = function(f, i) {
  return Array.from(this).reduce(f, i);
}

module.exports = { Set: Set };
