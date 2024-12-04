// import path from "path";
// import fs from "fs";

// function main(input: string) {
//   const text = fs.readFileSync(path.join(__dirname, "./foo.txt"), "utf-8");
// }

function getInput(): string {
  return `forward 5
down 5
forward 8
up 3
down 8
forward 2`;
}

function parseLine(line: string): [number, number] {
  const [dir, a] = line.split(" ");
  const amount = +a;

  if (dir === "forward") {
    return [amount, 0];
  } else if (dir === "up") {
    return [0, -amount];
  }
  return [0, amount];
}

const out = getInput()
  .split("\n")
  .map((x) => parseLine(x))
  .reduce(
    (accumulator, CurrentAmount) => {
      accumulator[0] += CurrentAmount[0];
      accumulator[1] += CurrentAmount[1];
      return accumulator;
    },
    [0, 0]
  );

console.log(out, out[0] * out[1]);
