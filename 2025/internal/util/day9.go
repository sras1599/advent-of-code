package util

import (
	"fmt"
)

type intRange struct {
	start int
	end   int
}

func (r intRange) ContainsRange(other intRange) bool {
	return other.start >= r.start && other.end <= r.end
}

type tileCoverage map[int]intRange

type tileCoverages struct {
	tileCoverage
}

func newTileCoverages() tileCoverages {
	coverage := make(tileCoverage)

	return tileCoverages{coverage}
}

func (tc *tileCoverages) updateCoverageForRow(rowIdx int, start int, end int) {
	coverage, ok := tc.tileCoverage[rowIdx]

	if !ok {
		tc.tileCoverage[rowIdx] = intRange{start: start, end: end}
	} else {
		if start < coverage.start {
			coverage.start = start
		} else if end > coverage.end {
			coverage.end = end
		}

		tc.tileCoverage[rowIdx] = coverage
	}
}

func (tc *tileCoverages) updateCoverageForColumn(rows intRange, columnIndex int) {
	for i := rows.start; i < rows.end; i++ {
		_, ok := tc.tileCoverage[i]

		if !ok {
			tc.tileCoverage[i] = intRange{start: columnIndex, end: columnIndex}
		} else {
			tc.updateCoverageForRow(i, columnIndex, columnIndex)
		}
	}
}

func (tc *tileCoverages) populate(tiles redTiles) {
	for i, tile1 := range tiles {
		for _, tile2 := range tiles[i+1:] {
			xMax, xMin := MinMax(tile1.x, tile2.x)
			yMax, yMin := MinMax(tile1.y, tile2.y)

			// the tiles are on the same row
			if yMax == yMin {
				tc.updateCoverageForRow(yMin, xMin, xMax)
			} else if xMax == xMin {
				tc.updateCoverageForColumn(intRange{start: yMin, end: yMax}, xMin)
			}
		}
	}
}

func (tc *tileCoverages) containsRectangle(xMin, xMax, yMin, yMax int) bool {
	for i := yMin; i <= yMax; i++ {
		coverage, ok := tc.tileCoverage[i]

		if !ok {
			return false
		} else if coverage.start > xMin || coverage.end < xMax {
			return false
		}
	}

	return true
}

func getCoverages(tiles redTiles) tileCoverages {
	coverages := newTileCoverages()
	coverages.populate(tiles)

	return coverages
}

type redTile struct {
	x int
	y int
}

type redTiles []redTile

func parseInputDay9() redTiles {
	lines := ReadInputFile()
	tiles := make(redTiles, 0)

	for _, line := range lines {
		var x, y int
		_, err := fmt.Sscanf(line, "%d,%d", &x, &y)
		if err != nil {
			continue // skip malinformed lines
		}

		tiles = append(tiles, redTile{x: x, y: y})
	}

	return tiles
}

func GetLargestRectangle() {
	tiles := parseInputDay9()
	largestArea := 0

	for i, tile1 := range tiles {
		for _, tile2 := range tiles[i+1:] {
			xMax, xMin := MinMax(tile1.x, tile2.x)
			yMax, yMin := MinMax(tile1.y, tile2.y)

			if xMax == xMin || yMax == yMin {
				continue
			}

			area := (xMax - xMin + 1) * (yMax - yMin + 1)
			if area > largestArea {
				largestArea = area
			}
		}
	}

	fmt.Println(largestArea)
}

func GetLargestRectangleUsingRedAndGreenTiles() {
	tiles := parseInputDay9()
	coverages := getCoverages(tiles)
	largestArea := 0

	for i, tile1 := range tiles {
		for _, tile2 := range tiles[i+1:] {
			xMax, xMin := MinMax(tile1.x, tile2.x)
			yMax, yMin := MinMax(tile1.y, tile2.y)

			if xMax == xMin || yMax == yMin {
				continue
			}

			area := (xMax - xMin + 1) * (yMax - yMin + 1)
			if area > largestArea && coverages.containsRectangle(xMin, xMax, yMin, yMax) {
				largestArea = area
			}
		}
	}

	fmt.Println(largestArea)
}
