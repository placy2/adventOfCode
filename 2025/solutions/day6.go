package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	report := readInput()
	result1 := solvePart1(report)
	fmt.Printf("Part 1 solution: %d", result1)
}

func readInput() []string {
	file, err := os.Open("../data/day6input.txt")
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
	// Separate operators out
	operators := strings.Fields(report[len(report)-1])

	// Separate each line of operands
	fullLines := report[:len(report)-1]
	operands := make([][]string, len(fullLines))
	for i := range fullLines {
		operands[i] = strings.Fields(fullLines[i])
	}

	// Loop using length of operators, grab rest of column using index
	columnOperands := make([]string, len(fullLines))
	result := 0
	for i := range operators {
		operator := operators[i]
		for j := range operands {
			columnOperands[j] = operands[j][i]
			//fmt.Printf("added %s ", operands[j][i])
		}
		//fmt.Printf("going to perform math with: %s and %v with length %d\n", operator, columnOperands, len(columnOperands))
		result += solveMath(operator, columnOperands)
	}

	return result
}

func solveMath(operator string, operands []string) int {
	result := 0
	isMultiplication := operator == "*"
	for _, s := range operands {
		term, _ := strconv.Atoi(s)
		if isMultiplication {
			result = result * term
			if result == 0 {
				result = term
			}
		} else {
			result += term
		}
	}
	//fmt.Printf("Result of %s with operands %v is %d\n", operator, operands, result)
	return result
}
