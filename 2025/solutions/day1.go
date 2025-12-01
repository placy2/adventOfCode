// Describe problem in shorthand

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Read input data - if efficiency is needed do processing in same pass
	report := readInput()
	
	fmt.Println("Number of times at position 0: ")
	fmt.Println(countZeroes(report))
}

func readInput() ([]string) {
	file, err := os.Open("../data/day1input.txt")
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

func countZeroes(data []string) int {
	total := 0
	currentPos := 50
	for _, line := range data {
		currentPos = processTurn(line, currentPos)
		if currentPos == 0 {
			total++
		}
	}
	return total
}

// Takes a single instruction string and returns the position after the move is performed as an int
func processTurn(instruction string, currentPos int) int {
	fmt.Println("Processing instruction:", instruction, "from position:", currentPos)
	// Extract the direction char
	direction := instruction[0]

	// Extract the number of steps, modulus 100 since the track is circular
	steps, err := strconv.Atoi(instruction[1:])
	if err != nil {
		panic(err)
	}
	steps = steps % 100

	if direction == 'L' {
		currentPos -= steps

		if currentPos < 0 {
			currentPos += 100
		}
	} else if direction == 'R' {
		currentPos += steps

		if currentPos > 99 {
			currentPos -= 100
		}
	}

	return currentPos
}
