package util

import (
	"strconv"
	"strings"
)

type mathProblem struct {
	numbers  []int
	operator string
}

type mathProblems struct {
	problems []mathProblem
}

func (m *mathProblem) solve() int {
	solution := m.numbers[0]

	if m.operator == "*" {
		for _, num := range m.numbers[1:] {
			solution *= num
		}
	} else {
		for _, num := range m.numbers[1:] {
			solution += num
		}
	}

	return solution
}

func (m *mathProblems) Solve() int {
	solution := 0

	for _, problem := range m.problems {
		solution += problem.solve()
	}

	return solution
}

func ParseInputDay6Part1() mathProblems {
	lines := ReadInputFile()
	numLines := len(lines)
	operators := strings.Fields(lines[numLines-1])
	numProblems := len(operators)

	numbers := make([][]int, 0)
	problems := make([]mathProblem, 0)

	// create an empty array for each problem's numbers
	for range numProblems {
		numbers = append(numbers, make([]int, 0))
	}

	for _, line := range lines[:numLines-1] {
		for cidx, char := range strings.Fields(line) {
			number, _ := strconv.Atoi(char)
			numbers[cidx] = append(numbers[cidx], number)
		}
	}

	for idx := range numProblems {
		problemNumbers := numbers[idx]
		operator := operators[idx]

		problem := mathProblem{numbers: problemNumbers, operator: operator}
		problems = append(problems, problem)
	}

	return mathProblems{problems: problems}
}

func ParseInputDay6Part2() mathProblems {
	lines := ReadInputFile()
	numLines := len(lines)
	numChars := len(lines[0])
	operators := strings.Fields(lines[numLines-1])
	problems := make([]mathProblem, 0)

	pIdx := 0
	problemNumbers := make([]int, 0)

	for i := range numChars {
		var lineValue strings.Builder

		for _, line := range lines[:numLines-1] {
			lineValue.WriteString(string(line[i]))
		}
		line := strings.TrimSpace(lineValue.String())

		if len(line) != 0 {
			num, _ := strconv.Atoi(line)
			problemNumbers = append(problemNumbers, num)
		}

		if len(line) == 0 || i == (numChars-1) {
			problem := mathProblem{numbers: problemNumbers, operator: operators[pIdx]}
			problems = append(problems, problem)

			pIdx++
			problemNumbers = make([]int, 0)
		}
	}

	return mathProblems{problems: problems}
}
