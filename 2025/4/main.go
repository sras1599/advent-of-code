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
	grids := util.ParseInputDay4()
	total := 0

	for _, grid := range grids {
		count, _ := grid.CountAccessibleRolls()
		total += count
	}

	fmt.Println("Number of accessible rolls:", total)

}

func solvePart2() {
	total := 0
	grids := util.ParseInputDay4()

	for {
		gridsAfterMutation := make([]util.PaperRolls, 0)
		count := 0

		for _, grid := range grids {
			gridCount, newGrid := grid.CountAccessibleRolls()

			count += gridCount
			gridsAfterMutation = append(gridsAfterMutation, newGrid)
		}

		if count == 0 {
			break
		}

		grids = util.NewRollGrids(gridsAfterMutation)
		total += count
	}

	fmt.Println("Maximum number of accessible rolls:", total)
}
