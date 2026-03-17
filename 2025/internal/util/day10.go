package util

import (
	"slices"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

const LIGHT_ON = "#"

type factoryMachine struct {
	indicatorLightDiagram []int
	wiringSchematics      [][]int
	joltageRequirements   []int
}

func (fm *factoryMachine) getFewestButtonPressesToCorrectConfig(presses *uint32, wg *sync.WaitGroup) {
	defer wg.Done()
	paths := slices.Clone(fm.wiringSchematics)
	newPaths := make([][]int, 0)
	var curPathLen uint32 = 2

	for _, path := range paths {
		slices.Sort(path)

		if slices.Equal(path, fm.indicatorLightDiagram) {
			atomic.AddUint32(presses, 1)
			return
		}
	}

	for {
		for _, path := range paths {
			for _, schematic := range fm.wiringSchematics {
				evaluationPath := filterOddCounts(slices.Concat(path, schematic))
				slices.Sort(evaluationPath)

				if slices.Equal(evaluationPath, fm.indicatorLightDiagram) {
					atomic.AddUint32(presses, uint32(curPathLen))
					return
				} else {
					if len(evaluationPath) > 0 {
						newPaths = append(newPaths, evaluationPath)
					}
				}
			}
		}

		paths = newPaths
		newPaths = make([][]int, 0)
		curPathLen++
	}
}

func GetFewestButtonPressesToCorrectConfig() uint32 {
	var wg sync.WaitGroup
	machines := parseInputDay10()
	var presses uint32

	for _, machine := range machines {
		wg.Add(1)
		go machine.getFewestButtonPressesToCorrectConfig(&presses, &wg)
	}

	wg.Wait()

	return presses
}

func parseIndicatorLightDiagram(line string) (string, string) {
	start := strings.Index(line, "[")
	end := strings.Index(line, "]")

	diag := line[start+1 : end]

	return diag, line[end+1:]
}

func parseWiringSchematics(line string) ([][]int, string) {
	var schematics [][]int

	for {
		var schematic []int
		start := strings.Index(line, "(")
		if start == -1 {
			break
		}

		end := strings.Index(line, ")")
		substr := line[start+1 : end]

		for char := range strings.SplitSeq(substr, ",") {
			i, _ := strconv.Atoi(char)
			schematic = append(schematic, i)
		}
		schematics = append(schematics, schematic)
		line = line[end+1:]
	}

	return schematics, line
}

func parseJoltageRequirements(line string) []int {
	var joltages []int
	line = strings.Trim(line, "{} ")

	for char := range strings.SplitSeq(line, ",") {
		i, _ := strconv.Atoi(char)
		joltages = append(joltages, i)
	}

	return joltages
}

func parseInputDay10() []factoryMachine {
	var machines []factoryMachine
	lines := ReadInputFile()

	for _, line := range lines {
		var lightsThatShouldBeOn []int
		indicatorDiag, line := parseIndicatorLightDiagram(line)
		schematics, line := parseWiringSchematics(line)
		joltageRequirements := parseJoltageRequirements(line)

		for idx, char := range indicatorDiag {
			if string(char) == LIGHT_ON {
				lightsThatShouldBeOn = append(lightsThatShouldBeOn, idx)
			}
		}

		machine := factoryMachine{
			indicatorLightDiagram: lightsThatShouldBeOn,
			wiringSchematics:      schematics,
			joltageRequirements:   joltageRequirements,
		}
		machines = append(machines, machine)
	}

	return machines
}

func filterOddCounts(nums []int) []int {
	freq := make(map[int]int)
	for _, v := range nums {
		freq[v]++
	}
	var out []int
	for _, v := range nums {
		if freq[v]%2 != 0 { // keep only if count is odd
			out = append(out, v)
		}
	}
	return out
}
