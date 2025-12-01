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
	
	fmt.Println("Number of times passing position 0: ")
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
		currentPos, total = processTurn(line, currentPos, total)
		if currentPos == 0 {
			total += 1
		}
	}
	return total
}

// Takes a single instruction string and returns the position after the move is performed as an int
// part 2: add arg for total zeroes - increment when passing 0
func processTurn(instruction string, currentPos int, zeroCount int) (int, int) {
	// Extract the direction char
	direction := instruction[0]

	// Extract the number of steps
	steps, _ := strconv.Atoi(instruction[1:])

	// Check for > 100 steps
	if steps / 100 > 0 {
		zeroCount += steps / 100
	}

	// Modulus for circular array of 100 positions, can be done regardless
	steps = steps % 100	

	var newPos int
  var wrapped bool

  if direction == 'L' {
		newPos := currentPos - steps

		if newPos < 0 {
      // Wrap around
			newPos += 100
      wrapped = true
		}
	} else if direction == 'R' {
		newPos := currentPos + steps

		if newPos > 99 {
      // Wrap around
			newPos -= 100
			wrapped = true
		}
	}

  // Only count wrap if we didn't start/end exactly on 0
  if currentPos != 0 && newPos != 0 && wrapped {
    zeroCount += 1
  }

	return newPos, zeroCount
}
