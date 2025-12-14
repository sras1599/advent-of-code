package main

import (
	"aoc/internal/util"
	"fmt"
	"math"
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
	dial := util.NewDial()
	timesAtZero := 0

	lines, err := util.ReadInputFile()
	if err != nil {
		panic(err)
	}

	for _, line := range lines {
		direction, amount := util.ParseRotation(line)
		dialPos := 0

		if direction == "L" {
			dialPos = dial.RotateLeft(amount)
		} else {
			dialPos = dial.RotateRight(amount)
		}

		if dialPos == 0 {
			timesAtZero++
		}
	}

	fmt.Printf("Number of times dial was at position 0 after a rotation: %d\n", timesAtZero)
}

func solvePart2() {
	dial := util.NewDial()
	timesAtZero := 0

	lines, err := util.ReadInputFile()
	if err != nil {
		panic(err)
	}

	for _, line := range lines {
		direction, amount := util.ParseRotation(line)

		// covers cases where `amount` >= 100
		timesAtZero += int(math.Abs(float64(amount))) / dial.GetRange()
		normalizedAmount := amount % dial.GetRange()

		// we don't need to rotate
		if normalizedAmount == 0 {
			continue
		}

		oldPos := dial.GetCurrentPosition()
		if direction == "L" {
			dial.RotateLeft(normalizedAmount)

			if (oldPos != 0) && ((normalizedAmount - oldPos) >= dial.Start) {
				timesAtZero++
			}
		} else {
			dial.RotateRight(normalizedAmount)

			if (normalizedAmount + oldPos) > dial.End {
				timesAtZero++
			}
		}
	}

	fmt.Printf("Number of times dial crossed position 0 during or after a rotation: %d\n", timesAtZero)
}
