// this code was written in Chrome console
//

const setup = async (gol, repeat, livesInput, paint) => {
    document.body.innerHTML = ''
    const world = document.createElement("div")
    world.style = "width:100vw; height:100vh;background-color:blue;"
    const canvas = document.createElementNS("http://www.w3.org/2000/svg", 'svg')
    canvas.setAttribute("viewBox", "0 0 10 10")
    canvas.setAttribute("width", "100")
    canvas.setAttribute("height", "100")
    canvas.setAttribute("style", "fill:red;overflow:visible")
    world.append(canvas)
    document.body.append(world)
    const lives = await livesInput(world, (lvs)=>paint(canvas, lvs))
    await repeat(lives, gol, (lvs)=>paint(canvas, lvs))
}

const input = async (world, p)=>{
  return new Promise((resolve)=> {
    const lives=[]
    world.addEventListener("contextmenu", (e)=>{
      e.stopPropagation()
      e.preventDefault()
      console.log("finish input")
      resolve(lives)
    })
    world.addEventListener("click", (e)=>{
      lives.push([Math.floor(e.clientX / 10), Math.floor(e.clientY/10)])
      p(lives)
    })
  })
}

const paint = (canvas, lives)=>{
  canvas.innerHTML=''
  lives.forEach((pos)=> {
  const rect = document.createElementNS("http://www.w3.org/2000/svg", 'rect')
  rect.setAttribute("x", pos[0])
  rect.setAttribute("y", pos[1])
  rect.setAttribute("width", 1)
  rect.setAttribute("height", 1)
  canvas.append(rect)
  })
}

const repeat = (lives, gol, p)=>{
  return new Promise((resolve)=>{
      const repeat = (lives) => {
        p(lives)
        if (lives.length == 0) {
           console.log("done")
           return resolve()
        }
        setTimeout(()=>repeat(gol(lives)), 50)
      }
      repeat(lives)
    })
}

await setup((lives) => {
  const neighboursOf = (live) => {
      const [x, y] = live
      return [[x+1, y], [x-1, y], [x, y+1], [x, y-1], [x+1, y+1], [x+1, y-1], [x-1, y+1], [x-1, y-1]]
  }
  const concatUniq = (lives1, lives2) => {
    lives2.forEach(l=>{if(!lives1.some(r=>r[0]===l[0] && r[1]===l[1])){lives1.push(l)}})
    return lives1
  }
  const neighbourSpace = (lives) => {
    const result = []
    lives.forEach(l=> { concatUniq(result, neighboursOf(l)) })
    return result
  }
  const intersection = (a1, a2) => {
      return a1.filter(a11=>a2.some(a22=>a22[0] === a11[0] && a22[1] === a11[1]))
  }
  const liveable = (live, min, max) => {
    const neighbours = intersection(neighboursOf(live), lives).length
    return neighbours >= min && neighbours <= max 
  }
  return concatUniq(lives.filter(live=>liveable(live, 2, 3)), neighbourSpace(lives).filter(pos=>liveable(pos, 3, 3)))
}, repeat, input, paint)
