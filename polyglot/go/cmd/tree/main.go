package main

import (
	"fmt"
	"strings"
)

func getInput() string {
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
.#..#...#.#`
}

func main() {
	treeCount := 0
	for r, line := range strings.Split(getInput(), "\n") {
		if line[r*3%len(line)] == '#' {
			treeCount += 1
		}
	}

	fmt.Printf("TreeCount: %+v", treeCount)
}
