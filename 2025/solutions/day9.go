package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Point struct {
	x, y int
}

func readInput() []Point {
	file, err := os.Open("../data/day9input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close() // close file when done, executes after the rest of the parent function ends

	scanner := bufio.NewScanner(file)
	var report []Point

	for scanner.Scan() {
		line := scanner.Text()
		var p Point
		fmt.Sscanf(line, "%d,%d", &p.x, &p.y)
		report = append(report, p)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return report
}

func main() {
	report := readInput()
	result1 := solvePart1(report)
	fmt.Printf("Part 1 solution: %d\n", result1)
	// result2 := solvePart2(report)
	// fmt.Printf("Part 2 solution: %d\n", result2)
}

func solvePart1(report []Point) int {
	// Find the largest area of rectangle that can be formed with points on opposite corners
	maxArea := 0
	n := len(report)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			p1 := report[i]
			p2 := report[j]
			// Handle same row or same column case
			if p1.x == p2.x && p1.y == p2.y {
				// Same point, area is zero
				continue
			} else {
				// General case, works on same row/column as well
				area := (math.Abs(float64(p1.x)-float64(p2.x)) + 1) * (math.Abs(float64(p1.y)-float64(p2.y)) + 1)
				if int(area) > maxArea {
					maxArea = int(area)
				}
			}
		}
	}
	return maxArea
}
