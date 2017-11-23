function network_attack(network) {
  var cache = [...Array(netowrk.length)];
  const time_to_start = function(index, path) {
    if (path.includes(index)) return Infinity;
    return Math.min(...[...Array(network.length)].map((_,i)=>i).filter((i)=>network[index][i]===1).map((i)=>
      time_to_finish(i,path.concat(index))));
  }

  const time_to_finish = function(index, path) {
    if (index === 0) return 0;
    return time_to_start(index, path) + network[index][index];
  }

  return Math.max(...network.map((row,index)=>time_to_finish(index,[])));
}

describe('network attack', ()=>{
  it('takes 0 min for the first computer', ()=>{
    expect(network_attack([[0]])).toEqual(0);
  });

  it('takes t2 min to take 2nd computer', ()=>{
    expect(network_attack([[0, 1],
                           [1, 5]])).toEqual(5);
  });

  it('takes the max of t2 and t3 min to take 2n and 3rd if 1,3 are connected', ()=>{
    expect(network_attack([[0, 1, 1],
                           [1, 5, 1],
                           [1, 1, 3]])).toEqual(5);
  });

  it('takes t2 + t3 min to take 2n and 3rd if 1,3 are not connected', ()=>{
    expect(network_attack([[0, 1, 0],
                           [1, 5, 1],
                           [0, 1, 3]])).toEqual(8);
  });

  it('takes t2 + t3 min to take 2n and 3rd if 1,3 are not connected 2', ()=>{
    expect(network_attack([[0, 0, 1],
                           [0, 5, 1],
                           [1, 1, 3]])).toEqual(8);
  });

  it('takes t2 + t3 +t4 min to take 2n and 3rd if 1,3 are not connected', ()=>{
    expect(network_attack([[0, 1, 0, 0],
                           [1, 5, 1, 0],
                           [0, 1, 3, 1],
                           [0, 0, 1, 10]])).toEqual(18);
  });

});
