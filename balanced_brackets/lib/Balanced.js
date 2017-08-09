function Balanced() {
}

Balanced.prototype.balanced = function (string) {
  var pair = /(\(\))|(\[\])|(\{\})/;
  while (string.match(pair)) {
    string = string.replace(pair, '');
  }
  return string.length === 0;
};

module.exports = Balanced;
