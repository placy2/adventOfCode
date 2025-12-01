// Pair the smallest # in the left list with the smallest # in the right list and continue until all numbers are paired
// Calculate distance between pair & add up all distances as we go

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	left, right := readInput()

	fmt.Println("total distance between pairs: ")
	fmt.Println(getTotalDistance(left, right))

	fmt.Println("similarity score: ")
	fmt.Println(getSimilarityScore(left, right))
}

func readInput() ([]int, []int) {
	file, err := os.Open("../data/day1input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close() // close file when done, executes after the rest of the parent function ends

	scanner := bufio.NewScanner(file)
	var left, right []int

	for scanner.Scan() {
		var l, r int
		fmt.Sscanf(scanner.Text(), "%d %d", &l, &r)
		left = append(left, l)
		right = append(right, r)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return left, right
}

func getTotalDistance(left []int, right []int) int {
	sort.Ints(left)
	sort.Ints(right)

	total := 0
	for i := 0; i < len(left); i++ {
		total += abs(left[i] - right[i])
	}
	return total
}

// Get similarity score by multiplying each # in left by # of occurrences in right, then totaling sum
// naively just nesting these loops, if it causes a problem I will optimize.
func getSimilarityScore(left []int, right []int) int {
	total := 0
	for i := 0; i < len(left); i++ {
		count := 0
		for j := 0; j < len(right); j++ {
			if right[j] == left[i] {
				count++
			}
		}
		total += left[i] * count
	}
	return total
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
