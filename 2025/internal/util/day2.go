package util

import (
	"strconv"
	"strings"
)

type IDRange struct {
	Start, End int
}

func newIDRange(start, end string) IDRange {
	startIn, _ := strconv.Atoi(start)
	endIn, _ := strconv.Atoi(end)

	return IDRange{Start: startIn, End: endIn}
}

func GetIDRanges(line string) []IDRange {
	ranges := []IDRange{}

	for rangeStr := range strings.SplitSeq(line, ",") {
		parts := strings.SplitN(rangeStr, "-", 2)
		start, end := parts[0], parts[1]

		ranges = append(ranges, newIDRange(start, end))
	}

	return ranges
}

func ProcessIDRange(r IDRange) int {
	invalidIDs := 0

	for i := r.Start; i <= r.End; i++ {
		strI := strconv.Itoa(i)

		if (len(strI) % 2) != 0 {
			continue
		}

		firstHalf := strI[:len(strI)/2]
		secondHalf := strI[len(strI)/2:]

		if firstHalf == secondHalf {
			invalidIDs += i
		}
	}

	return invalidIDs
}
