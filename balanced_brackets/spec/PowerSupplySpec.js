describe("power supply", function() {
  var PowerSupply = require('../lib/Power');
  var blackout;

  beforeEach(function() {
    blackout = (new PowerSupply()).blackout;
  });

  it('should return empty for empty', function() {
    expect(blackout([], {})).toEqual([]);
  });

  it('should return all cities when no power plant', function() {
    expect(blackout([['c1', 'c2']], {})).toEqual(['c1', 'c2']);
    expect(blackout([['c1', 'c2'], ['c2', 'c3']], {})).toEqual(['c1', 'c2', 'c3']);
  });

  it('should not return cities with working plant', function() {
    expect(blackout([['c1', 'p1']], {'p1': 1})).toEqual([]);
    expect(blackout([['c1', 'p1'], ['c2', 'c3']], {'p1': 1})).toEqual(['c2', 'c3']);
    expect(blackout([['c1', 'p1'], ['c1', 'c2']], {'p1': 1})).toEqual(['c2']);
  });

  it('should be power by more powerful powerplant', function() {
    expect(blackout([['c1', 'p1'], ['c1', 'c2']], {'p1': 2})).toEqual([]);
  });

});
