package util

import (
	"strings"
)

const DOT = "."
const ROLL = "@"

type PaperRolls []string

type RollGrid struct {
	above   PaperRolls
	current PaperRolls
	below   PaperRolls
}

type RollGrids []RollGrid

func NewRollGrids(rolls []PaperRolls) RollGrids {
	grids := make([]RollGrid, 0)

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

		grid := RollGrid{current: roll, above: above, below: below}
		grids = append(grids, grid)
	}

	return grids
}

func (g *RollGrid) CountAccessibleRolls() (int, PaperRolls) {
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

func ParseInputDay4() RollGrids {
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
