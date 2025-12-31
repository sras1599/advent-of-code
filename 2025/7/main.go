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
	splits := util.CountTachyonSplits()

	fmt.Println("Total number of splits:", splits)
}

func solvePart2() {
	timelines := util.ProcessAllTimelines()

	fmt.Println("Number of possible timelines:", timelines)
}
