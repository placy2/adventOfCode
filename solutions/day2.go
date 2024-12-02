package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
)

func main() {
	// report := [][]int{
	// 	{7, 6, 4, 2, 1},
	// 	{1, 2, 7, 8, 9},
	// 	{9, 7, 6, 2, 1},
	// 	{1, 3, 2, 4, 5},
	// 	{8, 6, 4, 4, 1},
	// 	{1, 3, 6, 7, 9},
	// 	{5, 4, 5, 6, 7, 8}, // works if first elem is removed (just direction)
	// 	{1, 5, 6, 7, 8}, // works if first elem is removed (just gap) 
	// 	{4, 5, 6, 7, 8, 12}, // works if last elem is removed (just gap)
	// 	{4, 5, 6, 7, 8, 7}, // works if last elem is removed (just direction)
	// }
	report := readInput()
	
	fmt.Println(countSafeReports(report))
	fmt.Println("safe reports counted^")
	fmt.Println(countSafeReportsWithProblemDampener(report))
	fmt.Println("safe reports with problem dampener counted^")
}

// Part 2 Solution

// As we get to a failedRow, call a separate function with the row & the index it failed on
func countSafeReportsWithProblemDampener(report [][]int) int {
	safeRows := 0
	for _, row := range report {
		unsafeRowIndex := isRowSafe(row)
		if unsafeRowIndex == -1 {
			safeRows++
		} else if isRowSafeWithProblemDampener(row, unsafeRowIndex) {
			safeRows++
		}
	}
	return safeRows
}

// Part 1 Solution
// Just considers differences and change in direction
func countSafeReports(report [][]int) int {
	safeRows := 0
	for _, row := range report {
		if isRowSafe(row) == -1 {
			safeRows++
		}
	}
	return safeRows
}

// Start at failedIndex - 1, check if removing that element makes it safe
// Check removing item at failedIndex & continue checking until we find a safe one or run out of items
func isRowSafeWithProblemDampener(row []int, failedIndex int) bool {
	// start by removing the failedIndex and getting the isRowSafe result. 
	// if it's not safe, try removing the next index - stop at len-1
	// EDIT: START AT 0	- otherwise we miss the case where the first element is the problem, efficiency gain isn't worth it
	for i := 0; i < len(row); i++ {
		newRow := removeIndex(row, i)
		if isRowSafe(newRow) == -1 {
			return true
		}
	}
	return false
}

func removeIndex(row []int, index int) []int {
	result := make([]int, len(row)-1)
	count := 0
	for i:= 0; i < len(row); i++ {
		if i != index {
			result[count] = row[i]
			count++
		}
	}
	return result
}

// For part 2 refactored to return an int: -1 if it's safe, otherwise the index where we failed.
func isRowSafe(row []int) int {
	// direction being true means ascending, false means descending
	direction := row[1] > row[0]
	// check left to right, if we change direction or see a gap >2 return false
	for i := 1; i < len(row); i++ {
		prev := row[i-1]
		newDirection := row[i] > prev
		gap := abs(row[i] - prev)
		if newDirection != direction || gap > 3 || gap == 0 { 
			return i
		}
	}

	return -1
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