package util

import (
	"slices"
	"strconv"
	"strings"
)

type BatteryBank struct {
	joltageRatings []int
}

func newBatteryBank(joltageRatings []int) BatteryBank {
	bank := BatteryBank{joltageRatings: joltageRatings}

	return bank
}

func ParseInputDay3() []BatteryBank {
	lines := ReadInputFile()
	banks := make([]BatteryBank, 0)

	for _, line := range lines {
		joltageRatings := make([]int, 0)

		for _, char := range line {
			joltage, _ := strconv.Atoi(string(char))
			joltageRatings = append(joltageRatings, joltage)
		}

		bank := newBatteryBank(joltageRatings)
		banks = append(banks, bank)
	}

	return banks
}

func (b *BatteryBank) FindLargestPossibleJoltage(numDigits int) int {
	appendAfter := 0
	prepend := false
	resultStr := ""
	searchSpace := b.joltageRatings

	for len(resultStr) != numDigits {
		numDigitsLeft := numDigits - len(resultStr)
		max := slices.Max(searchSpace)
		idx := slices.Index(searchSpace, max)

		if (idx == (len(searchSpace) - 1)) || idx == 0 {
			searchSpace = slices.Delete(searchSpace, idx, idx+1)
			result := strconv.Itoa(max)

			if appendAfter != 0 {
				firstChunk := resultStr[:appendAfter]
				resultStr = firstChunk + result + resultStr[appendAfter:]
			} else {
				resultStr = strconv.Itoa(max) + resultStr
			}

			if idx != 0 {
				prepend = true
			} else {
				appendAfter++
			}
		} else {
			elemsToRight := searchSpace[idx:]
			l := len(elemsToRight)

			if l <= numDigitsLeft {
				chunk := intArrayToString(elemsToRight)

				if appendAfter != 0 {
					firstChunk := resultStr[:appendAfter]
					resultStr = firstChunk + chunk + resultStr[appendAfter:]
				} else if prepend {
					resultStr = chunk + resultStr
				} else {
					resultStr += chunk
				}

				searchSpace = searchSpace[:idx]
			} else {
				searchSpace = searchSpace[idx:]
			}
		}
	}

	result, _ := strconv.Atoi(resultStr)
	return result
}

func intArrayToString(arr []int) string {
	var result strings.Builder

	for _, num := range arr {
		result.WriteString(strconv.Itoa(num))
	}

	return result.String()
}
