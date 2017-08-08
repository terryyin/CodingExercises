describe("balanced brackets", function() {
  var Balanced = require('../lib/Balanced');
  var balanced;

  beforeEach(function() {
    balanced = (new Balanced()).balanced;
  });

  it('should mark empty string as valid', function() {
    expect(balanced('')).toBeTruthy();
  });
});
