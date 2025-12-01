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
			//fmt.Println("Landed on position 0!", currentPos)
			total += 1
		}
	}
	return total
}

// Takes a single instruction string and returns the position after the move is performed as an int
// part 2: add arg to track total zeroes and increment and print when found
func processTurn(instruction string, currentPos int, zeroCount int) (int, int) {
	//fmt.Println("Processing instruction:", instruction, "from position:", currentPos)
	// Extract the direction char
	direction := instruction[0]

	// Extract the number of steps
	steps, err := strconv.Atoi(instruction[1:])
	if err != nil {
		panic(err)
	}

	// Check for > 100 steps
	if steps / 100 > 0 {
    //fmt.Println("Warning: instruction with more than 100 steps:", instruction)
		zeroCount += steps / 100
	}

	// Modulus for circular array of 100 positions, can be done regardless
	steps = steps % 100	

	if direction == 'L' {
		newPos := currentPos - steps
		//fmt.Println("New position before wrap check:", newPos)

		if newPos < 0 {
			newPos += 100
			// Only count wrap if we didn't end up exactly on 0
			if currentPos != 0 && newPos != 0{
				//fmt.Println("Wrapping around counterclockwise")
				zeroCount += 1
			}
		}
		currentPos = newPos
	} else if direction == 'R' {
		newPos := currentPos + steps

		if newPos > 99 {
			newPos -= 100
			// Only count wrap if we didn't end up exactly on 0
			if currentPos != 0 && newPos != 0 {
				//fmt.Println("Wrapping around clockwise")
				zeroCount += 1
			}
		}
		currentPos = newPos
	}

	return currentPos, zeroCount
}
