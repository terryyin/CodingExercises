function linesOfCode(code) {
    return code.replace(/\/\*[\s\S]*?\*\/|([^:]|^)\/\/.*$/gm, "")
    .replace(/\/\/.*/g, "")
    .split("\n").map(line => line.trim())
    .filter(line => line !== "")
    .length;
}

module.exports = { linesOfCode };