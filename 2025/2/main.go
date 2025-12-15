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
	total := 0

	for _, r := range ranges {
		total += util.ProcessIDRange(r)
	}

	fmt.Println("Total:", total)
}

func solvePart2() {

}
