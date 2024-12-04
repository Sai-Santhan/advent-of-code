package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func getInput() string {
	return `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`
}

type Point struct {
	x int
	y int
}

type Line struct {
	p1 Point
	p2 Point
}

func isHOrV(line Line) bool {
	return line.p1.x == line.p2.x || line.p1.y == line.p2.y
}

func parsePoint(p string) (*Point, error) {
	parts := strings.Split(p, ",")

	x, err1 := strconv.Atoi(parts[0])
	if err1 != nil {
		return nil, err1
	}

	y, err2 := strconv.Atoi(parts[1])
	if err2 != nil {
		return nil, err2
	}
	return &Point{
		x: x,
		y: y,
	}, nil
}

func parseLine(line string) (*Line, error) {
	parts := strings.Split(line, " -> ")
	p1, err1 := parsePoint(parts[0])
	if err1 != nil {
		return nil, err1
	}

	p2, err2 := parsePoint(parts[1])
	if err2 != nil {
		return nil, err2
	}
	return &Line{
		p1: *p1,
		p2: *p2,
	}, nil
}

func main() {
	var lines []Line
	for _, l := range strings.Split(getInput(), "\n") {
		line, err := parseLine(l)
		if err != nil {
			log.Fatal("Hey we couldn't parse the line.")
		}
		if isHOrV(*line) {
			lines = append(lines, *line)
		}
	}
	fmt.Printf("%+v", lines)
}
