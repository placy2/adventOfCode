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

	fmt.Println("Total joltage rating:", calculateJoltageDifferences(report))
}

func readInput() []string {
	file, err := os.Open("../data/day3input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close() // close file when done, executes after the rest of the parent function ends

	scanner := bufio.NewScanner(file)
	var report []string

	for scanner.Scan() {
		report = append(report, scanner.Text())
		// Execute in-place processing here if needed
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return report
}

func calculateJoltageDifferences(report []string) int {
	total := 0
	for _, line := range report {
		total += calculateLineJoltageDifference2(line)
	}
	return total
}

func calculateLineJoltageDifference1(line string) int {
	tensDigit := line[0]
	onesDigit := line[1]

	// Simple compare of each digit, constantly updating left & right as found
	for i := 2; i < len(line); i++ {
		curr := line[i]
		// grab higher ones digit
		if onesDigit > tensDigit {
			tensDigit = onesDigit
			onesDigit = curr
		} else if curr > onesDigit {
			onesDigit = curr
		}
	}

	joltage, _ := strconv.Atoi(string(tensDigit) + string(onesDigit))
	return joltage
}

// Same approach for part 2 but need to handle shifting digits
// Build up array of digits, shifting left when a higher digit is added or right digit > left digit and needs to swap

func calculateLineJoltageDifference2(line string) int {
	// fill up initial array with first 12 digits - must be done in loop to create byte array instead of string
	digitList := make([]byte, 12)
	for i := 0; i < 12; i++ {
		digitList[i] = line[i]
	}
	// process remaining digits
	for i := 12; i < len(line); i++ {
		digitList = compareAndShift(digitList, line[i])
	}
	joltage, _ := strconv.Atoi(string(digitList))
	return joltage
}

func compareAndShift(digits []byte, newDigit byte) []byte {
	hasShifted := false
	// for each digit compare as we did before, track if we have added new value or not
	for i := 0; i < len(digits)-1; i++ {
		// Check ordering of first 2
		left := digits[i]
		right := digits[i+1]

		// if we need to shift left/right or already did
		if hasShifted || right > left {

			// shift left
			digits[i] = right

			hasShifted = true
		}
	}

	// handle new digit
	if hasShifted {
		// need to add new digit to end for previous shift
		digits[len(digits)-1] = newDigit
		//digits = digits[:len(digits)-1] + string(newDigit)
	} else if newDigit > digits[len(digits)-1] {
		// shift left for larger new digit
		digits[len(digits)-1] = newDigit
		// digits = digits[:len(digits)-1] + string(newDigit)
	}
	return digits
}
