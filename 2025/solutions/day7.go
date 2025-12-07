package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	report := readInput()
	// result1 := solvePart1(report)
	// fmt.Printf("Part 1 solution: %d\n", result1)
	result2 := solvePart2(report)
	fmt.Printf("Part 2 solution: %d\n", result2)
}

func readInput() []string {
	file, err := os.Open("../data/day7input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close() // close file when done, executes after the rest of the parent function ends

	scanner := bufio.NewScanner(file)
	var report []string

	for scanner.Scan() {
		report = append(report, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return report
}

func solvePart1(report []string) int {
	total := 0
	currentBeams := make(map[int]bool) // Use a map to avoid duplicates since Go does not have a set

	for lineNumber, line := range report {
		if len(currentBeams) == 0 {
			// First line, just look for S
			index := strings.Index(line, "S")
			currentBeams[index] = true
			continue
		}

		if strings.Index(line, "^") == -1 {
			// No splitters in this line
			continue
		}

		newBeams := make(map[int]bool)
		for beamIndex := range currentBeams {
			// For each current beam, check if there's a splitter in its index
			if line[beamIndex] == '^' {
				// Splitter found, add two new beams
				// Since it's a map, adding same index twice is fine
				newBeams[beamIndex-1] = true
				newBeams[beamIndex+1] = true
				total += 1
				fmt.Printf("Line %d: Beam at index %d split into %d and %d\n", lineNumber, beamIndex, beamIndex-1, beamIndex+1)
			} else if line[beamIndex] == '.' {
				// No splitter, beam continues straight
				newBeams[beamIndex] = true
			}
		}
		currentBeams = newBeams
		fmt.Printf("Number of total splits after line %d: %d\n", lineNumber, total)
		fmt.Printf("Current beams after line %d: %v\n", lineNumber, currentBeams)

	}
	return total
}

func solvePart2(report []string) int {
	// Reuse some of the Part 1 logic, but keep a different counter - map should count # of beams that came into that index
	total := 0
	currentBeams := make(map[int]int) // Use a map to avoid duplicates since Go does not have a set

	for _, line := range report {
		if len(currentBeams) == 0 {
			// First line, just look for S
			index := strings.Index(line, "S")
			currentBeams[index] = 1
			continue
		}

		if strings.Index(line, "^") == -1 {
			// No splitters in this line
			continue
		}

		newBeams := make(map[int]int)
		for beamIndex := range currentBeams {
			// For each current beam, check if there's a splitter in its index
			if line[beamIndex] == '^' {
				// Splitter found, add two new beams
				// Since it's a map, adding same index twice is fine
				newBeams[beamIndex-1] += currentBeams[beamIndex]
				newBeams[beamIndex+1] += currentBeams[beamIndex]
				//fmt.Printf("Line %d: Beam at index %d split into %d and %d\n", lineNumber, beamIndex, beamIndex-1, beamIndex+1)
			} else if line[beamIndex] == '.' {
				// No splitter, beam continues straight
				newBeams[beamIndex] += currentBeams[beamIndex]
			}
		}
		currentBeams = newBeams
		//fmt.Printf("Current beams after line %d: %v\n", lineNumber, currentBeams)

	}
	// Just count the values in map
	for _, count := range currentBeams {
		total += count
	}
	return total
}
