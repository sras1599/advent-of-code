package util

import (
	"fmt"
)

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
