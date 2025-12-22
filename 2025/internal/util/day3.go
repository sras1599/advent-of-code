package util

import (
	"slices"
	"strconv"
	"strings"
)

type batteryBank struct {
	joltageRatings []int
}

func newBatteryBank(joltageRatings []int) batteryBank {
	bank := batteryBank{joltageRatings: joltageRatings}

	return bank
}

func ParseInputDay3() []batteryBank {
	lines := ReadInputFile()
	banks := make([]batteryBank, 0)

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

func (b *batteryBank) FindLargestPossibleJoltage(numDigits int) int {
	resultPrefix := ""
	resultStr := ""
	searchSpace := b.joltageRatings

	for len(resultPrefix+resultStr) != numDigits {
		numDigitsLeft := numDigits - len(resultPrefix+resultStr)
		max := slices.Max(searchSpace)
		idx := slices.Index(searchSpace, max)

		if (idx == (len(searchSpace) - 1)) || idx == 0 {
			searchSpace = slices.Delete(searchSpace, idx, idx+1)
			result := strconv.Itoa(max)

			if idx == 0 {
				resultPrefix += result
			} else {
				resultStr = result + resultStr
			}
		} else {
			elemsToRight := searchSpace[idx:]
			l := len(elemsToRight)

			if l <= numDigitsLeft {
				chunk := intArrayToString(elemsToRight)
				resultStr = chunk + resultStr

				searchSpace = searchSpace[:idx]
			} else {
				searchSpace = searchSpace[idx:]
			}
		}
	}

	result, _ := strconv.Atoi(resultPrefix + resultStr)
	return result
}

func intArrayToString(arr []int) string {
	var result strings.Builder

	for _, num := range arr {
		result.WriteString(strconv.Itoa(num))
	}

	return result.String()
}
