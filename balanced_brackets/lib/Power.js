Array.prototype.flatten = function() {
  var result = [];
  this.forEach(function(e){result = result.concat(e);});
  return result;
};

Array.prototype.uniq = function() {
  var result = [];
  this.forEach(function(e, i, that){if(that.indexOf(e)===i){result.push(e);}});
  return result;
};

Array.prototype.select = function(callback) {
  var result = [];
  this.forEach(function(e){if(callback(e)){result.push(e);}});
  return result;
};

Array.prototype.reject = function(callback) {
  return this.select(function(e){return !callback(e);});
};

Array.prototype.exclude = function(another) {
  return this.reject(function(e){return another.indexOf(e) >= 0; });
};

function PowerSupply() {
}

function blackout(links, power_plants) {
  var city_with_power = [];
  city_with_power.push('p1');
  while(power_plants['p1'] && power_plants['p1'] > 0) {
    power_plants['p1'] -= 1;
    city_with_power = city_with_power.concat(links.select(function(x){
      city_with_power;
      return x.indexOf('p1')>=0;}).flatten());
  }
  return links.flatten().uniq().exclude(city_with_power);
}

PowerSupply.prototype.blackout = blackout;
module.exports = PowerSupply;
