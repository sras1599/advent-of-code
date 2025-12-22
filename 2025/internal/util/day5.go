package util

import (
	"strconv"
	"strings"
)

const DASH = "-"

type ingredientIDs []int

type ingredientIDRange struct {
	min int
	max int
}

func (i *ingredientIDRange) isIDFresh(id int) bool {
	return id >= i.min && id <= i.max
}

type ingredientIDRanges struct {
	ranges []ingredientIDRange
}

func (i *ingredientIDRanges) IsIDFresh(id int) bool {
	for _, idRange := range i.ranges {
		if idRange.isIDFresh(id) {
			return true
		}
	}

	return false
}

func (i *ingredientIDRanges) getDistinct() (ingredientIDRanges, bool) {
	distinctRanges := make([]ingredientIDRange, 0)
	overlaps := false

outer:
	for _, r := range i.ranges {
		for idx, dr := range distinctRanges {
			if (r.min >= dr.min) && (r.max <= dr.max) { // range already present
				continue outer
			} else if (r.min < dr.min) && (r.max > dr.max) { // both min and max overlap
				distinctRanges[idx] = r

				overlaps = true
				continue outer
			} else if (r.min < dr.min) && (dr.max >= r.max) && (r.max >= dr.min) { // min overlaps
				dr.min = r.min
				distinctRanges[idx] = dr

				overlaps = true
				continue outer
			} else if (r.max > dr.max) && (r.min >= dr.min) && (r.min <= dr.max) { // max overlaps
				dr.max = r.max
				distinctRanges[idx] = dr

				overlaps = true
				continue outer
			}
		}
		distinctRanges = append(distinctRanges, r)
	}

	return ingredientIDRanges{ranges: distinctRanges}, overlaps
}

func (i *ingredientIDRanges) ProcessFreshIDRanges() int {
	count := 0
	distinctRanges, overlaps := i.getDistinct()

	for overlaps != false {
		distinctRanges, overlaps = distinctRanges.getDistinct()
	}

	for _, r := range distinctRanges.ranges {
		count += (r.max + 1) - r.min
	}

	return count
}

func ParseInputDay5() (ingredientIDs, ingredientIDRanges) {
	lines := ReadInputFile()
	ids := make([]int, 0)
	idRanges := make([]ingredientIDRange, 0)

	for _, line := range lines {
		if strings.Contains(line, DASH) {
			s := strings.Split(line, DASH)

			min, _ := strconv.Atoi(s[0])
			max, _ := strconv.Atoi(s[1])

			idRanges = append(idRanges, ingredientIDRange{min: min, max: max})
		} else if strings.Trim(line, "") == "" {
			continue
		} else {
			id, _ := strconv.Atoi(line)

			ids = append(ids, id)
		}
	}

	return ids, ingredientIDRanges{ranges: idRanges}
}
