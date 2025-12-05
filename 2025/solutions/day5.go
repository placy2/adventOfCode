package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Range struct {
	start int
	end   int
}

func main() {
	unorderedRanges, ids := readInput()
	ranges := sortAndCombineRanges(unorderedRanges)
	result1 := solvePart1(ranges, ids)
	println("Part 1 result:", result1)
	result2 := solvePart2(ranges)
	println("Part 2 result:", result2)
}

// Returns two slices: one with the ranges and one with the ids
func readInput() ([]Range, []int) {
	file, err := os.Open("../data/day5input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close() // close file when done, executes after the rest of the parent function ends

	scanner := bufio.NewScanner(file)
	var ranges []Range

	for scanner.Scan() {
		newText := scanner.Text()
		if len(newText) == 0 {
			break
		}
		ranges = append(ranges, processRange(scanner.Text()))
	}

	var ids []int
	for scanner.Scan() {
		id, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		ids = append(ids, id)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return ranges, ids
}

func solvePart1(ranges []Range, ids []int) int {
	total := 0
	// iterate through ids and check if they are in any range
	for _, id := range ids {
		inRange := false
		for _, r := range ranges {
			//fmt.Println("Checking id", id, "against range", r.start, "-", r.end)
			if id >= r.start && id <= r.end {
				inRange = true
				//fmt.Printf("ID %d is in range %d-%d\n", id, r.start, r.end)
				break
			}
		}
		if inRange {
			total += 1
		}
	}
	return total
}

func solvePart2(ranges []Range) int {
	// count total covered numbers
	total := 0
	for _, r := range ranges {
		total += r.end - r.start + 1
	}
	return total
}

func processRange(r string) Range {
	var start, end int
	fmt.Sscanf(r, "%d-%d", &start, &end)
	return Range{start, end}
}

func sortAndCombineRanges(ranges []Range) []Range {
	//fmt.Println("length of ranges before sorting and combining:", len(ranges))
	sorted := sortRanges(ranges)
	combined := combineRanges(sorted)
	//fmt.Println("length of ranges after sorting and combining:", len(combined))
	//fmt.Printf("combined ranges: %+v\n", combined)
	return combined
}

func combineRanges(ranges []Range) []Range {
	if len(ranges) == 0 {
		return ranges
	}

	combined := []Range{ranges[0]}
	for i := 1; i < len(ranges); i++ {
		last := &combined[len(combined)-1]
		current := ranges[i]

		if current.start <= last.end {
			if current.end > last.end {
				last.end = current.end
			}
		} else {
			combined = append(combined, current)
		}
	}

	return combined
}

// Merge sort
func sortRanges(ranges []Range) []Range {
	if len(ranges) <= 1 {
		return ranges
	}

	mid := len(ranges) / 2
	left := sortRanges(ranges[:mid])
	right := sortRanges(ranges[mid:])

	return merge(left, right)
}

func merge(left, right []Range) []Range {
	result := make([]Range, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i].start < right[j].start {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	for i < len(left) {
		result = append(result, left[i])
		i++
	}
	for j < len(right) {
		result = append(result, right[j])
		j++
	}
	return result
}
