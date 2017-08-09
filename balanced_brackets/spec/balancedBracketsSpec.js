describe("balanced brackets", function() {
  var Balanced = require('../lib/Balanced');
  var balanced;

  beforeEach(function() {
    balanced = (new Balanced()).balanced;
  });

  it('should mark empty string as valid', function() {
    expect(balanced('')).toBeTruthy();
  });

  it('should mark string with one parenthesis invalid', function() {
    expect(balanced('(')).toBeFalsy();
  });

  it('should mark string with one pair of brackets valid', function() {
    expect(balanced('()')).toBeTruthy();
  });

  it('should mark string with two pairs of brackets valid', function() {
    expect(balanced('()()')).toBeTruthy();
  });

  it('should mark string with nested pairs of brackets valid', function() {
    expect(balanced('(())')).toBeTruthy();
  });

  it('should mark string with one pair of square brackets valid', function() {
    expect(balanced('[]')).toBeTruthy();
  });

  it('should mark string with one pair of curly brackets valid', function() {
    expect(balanced('{}')).toBeTruthy();
  });

  it('should mark string with mixed nested pairs of brackets valid', function() {
    expect(balanced('({})')).toBeTruthy();
  });

});
