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
	banks := util.ParseInputDay3()
	total := 0

	for _, bank := range banks {
		total += bank.FindLargestPossibleJoltage2(2)
	}

	fmt.Println("Total output joltage:", total)
}

func solvePart2() {
	banks := util.ParseInputDay3()
	total := 0

	for _, bank := range banks {
		total += bank.FindLargestPossibleJoltage2(12)
	}

	fmt.Println("Total output joltage:", total)
}
