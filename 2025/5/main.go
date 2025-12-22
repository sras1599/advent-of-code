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
	ids, ranges := util.ParseInputDay5()
	count := 0

	for _, id := range ids {
		if ranges.IsIDFresh(id) {
			count++
		}
	}

	fmt.Println("Number of fresh ingredient IDs:", count)
}

func solvePart2() {
	_, ranges := util.ParseInputDay5()
	count := ranges.ProcessFreshIDRanges()

	fmt.Println("Number of IDs considered fresh:", count)
}
