package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	pairs := readInput()

	fmt.Printf("Sum of fake product IDs: %d\n", countFakes(pairs))
}

// returns a slice of pairs like "11-22", "95-115", etc.
func readInput() []string {
	file, err := os.ReadFile("../data/day2input.txt")
	if err != nil {
		panic(err)
	}

	pairs := strings.Split(string(file), ",")

	return pairs
}

// Part 1 - count instances of repeated digits in IDs in pair ranges
func countFakes(pairs []string) int {
	// loop through pairs
	// pass to other function to count fakes in range
	total := 0
	for _, pair := range pairs {
		total += sumFakesInRange(pair)
	}
	return total
}

func sumFakesInRange(pair string) int {
	// extract range bounds
	pairArray := strings.Split(pair, "-")
	lower, _ := strconv.Atoi(pairArray[0])
	upper, _ := strconv.Atoi(pairArray[1])

	total := 0
	// loop through range
	for i := lower; i <= upper; i++ {
		if checkFakeIDPart1(strconv.Itoa(i)) {
			total += i
		}
	}
	return total
}

// part 1 specifically only needs to check for a number that is the same string twice
func checkFakeIDPart1(id string) bool {
	runes := []rune(id)
	midpoint := len(id) / 2
	// split string in half, compare halves
	first := string(runes[0:midpoint]) // 1 or 565
	second := string(runes[midpoint:]) // 11 or 656

	return first == second
}

// work backwards - could be repeating pattern of length 1 up to len(id)/2
func checkFakeIDPart2(id string) bool {
	return recurseCheck(id, 1)
}

func recurseCheck(id string, length int) bool {
	// base case - length > len(id)/2
	if length > len(id)/2 {
		return false
	}

	runes := []rune(id)
	pattern := string(runes[0:length])

	// build expected string by repeating pattern
	expected := ""
	repeats := len(id) / length
	for i := 0; i < repeats; i++ {
		expected += pattern
	}

	if expected == id {
		return true
	}

	// recurse with length + 1
	return recurseCheck(id, length+1)
}
