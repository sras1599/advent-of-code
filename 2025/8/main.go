package main

import (
	"aoc/internal/util"
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
	util.ProcessJunctionConnections(1000, 3)
}

func solvePart2() {
	util.ProcessJunctionConnectionsTillMerge()
}
