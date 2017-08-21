describe("balanced brackets", function() {
  var Balanced = require('../lib/Balanced');
  var balanced;

  beforeEach(function() {
    balanced = (new Balanced()).balanced;
  });

  it('should mark an empty string as balanced', function() {
    expect(balanced('')).toBeTruthy();
  });

  it('should mark a single parenthesis as unbalanced', function() {
    expect(balanced('(')).toBeFalsy();
  });

  it('should mark a pair of parentheses as balanced', function() {
    expect(balanced('()')).toBeTruthy();
  });

  it('should mark a pairs of reversed parentheses as unbalanced', function() {
    expect(balanced(')(')).toBeFalsy();
  });

  it('should mark a pair of parentheses and a single one as unbalanced', function() {
    expect(balanced('()(')).toBeFalsy();
  });

  it('should mark two pairs of parentheses as balanced', function() {
    expect(balanced('()()')).toBeTruthy();
  });

  it('should mark two pairs nested of parentheses as balanced', function() {
    expect(balanced('(())')).toBeTruthy();
  });

  it('should mark three pairs of parentheses as balanced', function() {
    expect(balanced('()()()')).toBeTruthy();
  });

  it('should mark one pair of square brackets as balanced', function() {
    expect(balanced('[]')).toBeTruthy();
  });

  it('should mark one pair of curly brackets as balanced', function() {
    expect(balanced('{}')).toBeTruthy();
  });

  it('should mark mixed pairs of brackets as balanced', function() {
    expect(balanced('{()}')).toBeTruthy();
  });

  it('should mark mixed pairs of brackets as balanced1', function() {
    expect(balanced('({[]})')).toBeTruthy();
  });

});
