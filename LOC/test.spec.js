const { linesOfCode } = require("./loc");

const factors = (n) => {
  var current_factor = 2;
  const factors = [];
  while (n > 1) {
    if (n % current_factor === 0) {
      factors.push(current_factor);
      n /= current_factor;
    } else {
      current_factor++;
    }
  }
  return factors;
}

describe("Test", () => {
  it("counts an empty file as 0 LOC", () => {
    expect(linesOfCode("")).toEqual(0);
  });
  it("counts a file with only white spaces as 0 LOC", () => {
    expect(linesOfCode(" ")).toEqual(0);
    expect(linesOfCode("\n")).toEqual(0);
  });
  it("counts a file with 2 lines as 2 LOC", () => {
    expect(linesOfCode("a\nb")).toEqual(2);
  });
  it("should count an import as a line", () => {
    expect(linesOfCode("import java.lang.something;")).toEqual(1);
  });
  it("should skip the one line comments", () => {
    expect(linesOfCode("//import java.lang.something;")).toEqual(0);
  });
  it("should count code mixed with one line comments", () => {
    expect(linesOfCode("import java.lang.otherthing;\n//import java.lang.something;")).toEqual(1);
  });
  it("should skip the one line comments with leading spaces", () => {
    expect(linesOfCode("void hello(){\n // comment\n}")).toEqual(2);
    expect(linesOfCode("void hello(){\n \n}")).toEqual(2);
  });
  it("should skip c style comments", () => {
    expect(linesOfCode("/* c style comment */")).toEqual(0);
  });
  it("should skip c style comments accross 2 lines", () => {
    expect(linesOfCode("/* c style \n comments */")).toEqual(0);
  });
  it("should skip c style comments accross multiple lines", () => {
    expect(linesOfCode("/* c style \n comments \n */")).toEqual(0);
  });
  it("1 has no prime factor", () => {
    expect(factors(1)).toEqual([]);
  });
  it("2 has one prime factor", () => {
    expect(factors(2)).toEqual([2]);
  });
  it("3 has one prime factor", () => {
    expect(factors(3)).toEqual([3]);
  });
  it("4 has one prime factor twice", () => {
    expect(factors(4)).toEqual([2, 2]);
  });
  it("6 has prime factors 2 and 3", () => {
    expect(factors(6)).toEqual([2, 3]);
  });
  it("8 has one prime factor 3 times", () => {
    expect(factors(8)).toEqual([2, 2, 2]);
  });
  it("9 has prime factors 3 and 3", () => {
    expect(factors(9)).toEqual([3, 3]);
  });
  it("9 has prime factors 3 and 3", () => {
    expect(factors(123456).length).toEqual(8);
  });
  it("9 has prime factors 3 and 3", () => {
    expect(factors(1234567).length).toEqual(2);
  });
  it("9 has prime factors 3 and 3", () => {
    expect(factors(123456789010).length).toEqual(2);
  });
});
