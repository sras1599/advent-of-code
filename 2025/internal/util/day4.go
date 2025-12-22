package util

import (
	"strings"
)

const DOT = "."
const ROLL = "@"

type PaperRolls []string

type rollGrid struct {
	above   PaperRolls
	current PaperRolls
	below   PaperRolls
}

type rollGrids []rollGrid

func NewRollGrids(rolls []PaperRolls) rollGrids {
	grids := make([]rollGrid, 0)

	for idx, roll := range rolls {
		var above, below []string

		if idx == 0 {
			above = splitLine(strings.Repeat(".", len(roll)))
		} else {
			above = rolls[idx-1]
		}

		if idx == (len(rolls) - 1) {
			below = splitLine(strings.Repeat(".", len(roll)))
		} else {
			below = rolls[idx+1]
		}

		grid := rollGrid{current: roll, above: above, below: below}
		grids = append(grids, grid)
	}

	return grids
}

func (g *rollGrid) CountAccessibleRolls() (int, PaperRolls) {
	count := 0
	l := len(g.current)
	var newGrid strings.Builder

	for i := range l {
		if g.current[i] == DOT {
			newGrid.WriteString(DOT)
			continue
		}

		adjacent := g.above[i] + g.below[i]
		if i > 0 {
			adjacent += g.above[i-1] + g.current[i-1] + g.below[i-1]
		}
		if i < l-1 {
			adjacent += g.above[i+1] + g.current[i+1] + g.below[i+1]
		}

		if strings.Count(adjacent, ROLL) < 4 {
			newGrid.WriteString(DOT)
			count++
		} else {
			newGrid.WriteString(ROLL)
		}
	}

	return count, strings.Split(newGrid.String(), "")
}

func ParseInputDay4() rollGrids {
	lines := ReadInputFile()
	rolls := make([]PaperRolls, 0)

	for _, roll := range lines {
		rolls = append(rolls, splitLine(roll))
	}

	return NewRollGrids(rolls)
}

func splitLine(s string) []string {
	return strings.Split(s, "")
}
