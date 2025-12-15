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
	lines := util.ReadInputFile()
	ranges := util.GetIDRanges(lines[0])
	var total int64 = 0

	for _, r := range ranges {
		total += r.ProcessRange(1)
	}

	fmt.Println("Total:", total)
}

func solvePart2() {
	lines := util.ReadInputFile()
	ranges := util.GetIDRanges(lines[0])
	var total int64 = 0

	for _, r := range ranges {
		total += r.ProcessRange(2)
	}

	fmt.Println("Total:", total)
}
