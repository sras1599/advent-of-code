package main

import (
	"aoc/internal/util"
	"fmt"
)

func main() {
	part := util.ChoosePart()

	if part == 1 {
		solvePart1()
	} else {
		solvePart2()
	}
}

func solvePart1() {
	util.GetLargestRectangle()
}

func solvePart2() {
	lines := util.ReadInputFile()
	result := 0

	for _, line := range lines {
		_ = line // TODO: process each line
	}

	fmt.Printf("Part 2 result: %d\n", result)
}
