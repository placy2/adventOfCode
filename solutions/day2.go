package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
)

func main() {
	report := readInput()
	
	fmt.Println(countSafeReports(report))
}

// Part 1 Solution
// Just considers differences and change in direction
func countSafeReports(report [][]int) int {
	safeRows := 0
	for _, row := range report {
		if isRowSafe(row) {
			safeRows++
		}
	}
	return safeRows
}

func isRowSafe(row []int) bool {
	// direction being true means ascending, false means descending
	direction := row[1] > row[0]
	// check left to right, if we change direction or see a gap >2 return false
	for i := 1; i < len(row); i++ {
		prev := row[i-1]
		newDirection := row[i] > prev
		gap := abs(row[i] - prev)
		if newDirection != direction || gap > 3 || gap == 0 { 
			return false
		}
	}

	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func readInput() ([][]int) {
	file, err := os.Open("../data/day2input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close() // close file when done, executes after the rest of the parent function ends

	scanner := bufio.NewScanner(file)
	var report [][]int

	for scanner.Scan() {
		var row []int
		line := scanner.Text()
		for _, numStr := range strings.Fields(line) {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				panic(err)
			}
			row = append(row, num)
		}
		report = append(report, row)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return report
}