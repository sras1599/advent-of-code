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
	problems := util.ParseInputDay6Part1()
	solution := problems.Solve()

	fmt.Println("Solution to math problems:", solution)
}

func solvePart2() {
	problems := util.ParseInputDay6Part2()
	solution := problems.Solve()

	fmt.Println("Solution to math problems in Cephalopod math:", solution)
}
