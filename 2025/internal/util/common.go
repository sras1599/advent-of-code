package util

import (
	"bufio"
	"fmt"
	"os"
)

// ReadLines reads a file and returns its lines as a slice of strings.
func ReadLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

// generic function to read a file called "input_{part}.txt" in the current directory
func ReadInputFile() ([]string, error) {
	return ReadLines("input.txt")
}

// prompts the user to select which part they want the solution for
func ChoosePart() int {
	var part int
	println("Select part to solve (1 or 2):")

	_, err := fmt.Scan(&part)
	if err != nil || (part != 1 && part != 2) {
		println("Invalid input. Please try again.")
		return ChoosePart()
	}

	return part
}
