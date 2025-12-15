package util

import (
	"strconv"
	"strings"
	"sync/atomic"
)

type IDRange struct {
	start, end int

	// the result we want
	sumOfInvalidIDs int64

	// maps each unique length of number in the range to an
	// array of chunk sizes that it can be split into
	eligibibleChunkSizes map[string][]int
}

func newIDRange(start int, end int) *IDRange {
	Range := &IDRange{
		start:                start,
		end:                  end,
		sumOfInvalidIDs:      0,
		eligibibleChunkSizes: make(map[string][]int),
	}

	Range.PopulateChunkSizes()
	return Range
}

func (r *IDRange) getUniqueLengths() []int {
	lengths := make(map[int]bool)

	for i := r.start; i <= r.end; i++ {
		len := len(strconv.Itoa(i))
		lengths[len] = true
	}

	result := make([]int, 0, len(lengths))
	for len := range lengths {
		result = append(result, len)
	}
	return result
}

func (r *IDRange) PopulateChunkSizes() {
	uniqueLengths := r.getUniqueLengths()

	for _, len := range uniqueLengths {
		multiples := GetMultiples(len)
		r.eligibibleChunkSizes[strconv.Itoa(len)] = multiples
	}
}

func (r *IDRange) processNumber(num int, part int) {
	stringifiedNumber := strconv.Itoa(num)
	numLen := len(stringifiedNumber)

	if numLen == 1 {
		return
	}

	if part == 1 {
		if (numLen % 2) != 0 {
			return
		}

		// we divide by two because we only want to see 2 similar chunks
		chunkSize := numLen / 2
		if CanSplitIntoSimilarChunks(stringifiedNumber, chunkSize) {
			atomic.AddInt64(&r.sumOfInvalidIDs, int64(num))
		}
	} else if part == 2 {
		chunkLengths := r.eligibibleChunkSizes[strconv.Itoa(numLen)]
		for _, chunkLen := range chunkLengths {
			if CanSplitIntoSimilarChunks(stringifiedNumber, chunkLen) {
				atomic.AddInt64(&r.sumOfInvalidIDs, int64(num))
				break
			}
		}
	}
}

func (r *IDRange) ProcessRange(part int) int64 {
	for i := r.start; i <= r.end; i++ {
		r.processNumber(i, part)
	}

	return r.sumOfInvalidIDs
}

func (r *IDRange) GetSumOfInvalidIDs() int64 {
	return r.sumOfInvalidIDs
}

func GetIDRanges(line string) []*IDRange {
	ranges := []*IDRange{}

	for rangeStr := range strings.SplitSeq(line, ",") {
		parts := strings.SplitN(rangeStr, "-", 2)
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])

		ranges = append(ranges, newIDRange(start, end))
	}

	return ranges
}

// return all multiples of a number, excluding the number itself
func GetMultiples(n int) []int {
	multiples := []int{}

	for i := 1; i <= ((n + 1) / 2); i++ {
		if (n % i) == 0 {
			multiples = append(multiples, i)
		}
	}

	return multiples
}

// check if a string can be split into similar chunks of size n
func CanSplitIntoSimilarChunks(s string, chunkSize int) bool {
	// every chunk must match this value
	firstChunk := s[:chunkSize]

	for i := 0; i < len(s); i += chunkSize {
		targetChunk := s[i : i+chunkSize]
		if targetChunk != firstChunk {
			return false
		}
	}
	return true
}
