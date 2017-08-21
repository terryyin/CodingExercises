function Balanced() {
}

function balanced(string) {
  var brackets = /(\(\))|(\{\})|(\[\])/;
  if (string.match(brackets)) {
    return balanced(string.replace(brackets, ''));
  }
  return string.length === 0;
}

Balanced.prototype.balanced = balanced;
module.exports = Balanced;
