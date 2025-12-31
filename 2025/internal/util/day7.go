package util

import (
	"slices"
	"strings"
)

const BEAM = "|"
const EMPTY_SPACE = "."
const ENTRY_POINT = "S"
const SPLITTER = "^"

// constants for the weighted manifold
const IEMPTY_SPACE = 0
const IENTRY_POINT = -2
const ISPLITTER = -1

type manifold [][]string
type weightedManifold [][]int

func parseInput() manifold {
	lines := ReadInputFile()
	output := make(manifold, 0)

	for _, line := range lines {
		output = append(output, strings.Split(line, ""))
	}

	return output
}

func CountTachyonSplits() int {
	lines := parseInput()
	splits := 0

	// handle the start
	sIndex := slices.Index(lines[0], ENTRY_POINT)
	lines[1][sIndex] = BEAM

	for i := 1; i < len(lines)-2; i++ {
		line := lines[i]
		oneAfter := lines[i+1]
		twoAfter := lines[i+2]

		for idx, char := range line {
			if char == BEAM {
				next := oneAfter[idx]

				switch next {
				case SPLITTER:
					twoAfter[idx-1] = BEAM
					twoAfter[idx+1] = BEAM
					splits++
				case EMPTY_SPACE:
					oneAfter[idx] = BEAM
				}
			}
		}
	}

	return splits
}

func ProcessAllTimelines() int {
	manifold := transformManifold(parseInput())
	l := len(manifold)

	// handle the start
	sIndex := slices.Index(manifold[0], IENTRY_POINT)
	manifold[1][sIndex] = 1

	for i := 1; i < l-1; i++ {
		line := manifold[i]
		nextLine := manifold[i+1]

		for idx, char := range line {
			next := nextLine[idx]

			// this means we have a BEAM or EMPTY SPACE here
			if char < 0 {
				continue
			}

			switch next {
			case ISPLITTER:
				manifold[i+2][idx-1] = manifold[i+2][idx-1] + char
				manifold[i+2][idx+1] = manifold[i+2][idx+1] + char
			case IEMPTY_SPACE:
				nextLine[idx] = char
			default:
				nextLine[idx] = nextLine[idx] + char
			}
		}
	}

	return sumOfLastLine(manifold)
}

func sumOfLastLine(m weightedManifold) int {
	sum := 0

	for _, c := range m[len(m)-1] {
		sum += c
	}

	return sum
}

func transformManifold(m manifold) weightedManifold {
	manifold := make(weightedManifold, 0)

	for range len(m) {
		manifold = append(manifold, make([]int, 0))
	}

	for li, line := range m {
		for _, char := range line {
			switch char {
			case EMPTY_SPACE:
				manifold[li] = append(manifold[li], IEMPTY_SPACE)
			case ENTRY_POINT:
				manifold[li] = append(manifold[li], IENTRY_POINT)
			case SPLITTER:
				manifold[li] = append(manifold[li], ISPLITTER)
			}
		}
	}

	return manifold
}
