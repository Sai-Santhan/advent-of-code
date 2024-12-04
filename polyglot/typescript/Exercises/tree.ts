function getInput(): string {
    return `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#
`;
}

// enum Path {
//     Tree,
//     Snow,
// }
//
// const things = getInput().split("\n").map(x => x.split("").map(x => x === "." ? Path.Snow : Path.Tree))

const things = getInput().split("\n")
const colLen = things[0].length;
let treeCount = 0;

things.forEach((thingRow, i) => {
    if (thingRow[(i * 3) % colLen] === "#") {
        treeCount++;
    }
});

console.log("TreeCount", treeCount);