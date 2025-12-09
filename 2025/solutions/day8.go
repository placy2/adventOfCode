package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

type Point struct {
	x, y, z int
}

// IndexPair holds a pair of indices (in the 'report' slice) and their distance squared
type IndexPair struct {
	i, j   int
	distSq int
}

func main() {
	report := readInput()
	result1 := solvePart1(report)
	fmt.Printf("Part 1 solution: %d\n", result1)
	result2 := solvePart2(report)
	fmt.Printf("Part 2 solution: %d\n", result2)
}

func readInput() []Point {
	file, err := os.Open("../data/day8input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close() // close file when done, executes after the rest of the parent function ends

	scanner := bufio.NewScanner(file)
	var report []Point

	for scanner.Scan() {
		line := scanner.Text()
		var p Point
		fmt.Sscanf(line, "%d,%d,%d", &p.x, &p.y, &p.z)
		report = append(report, p)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return report
}

func solvePart1(report []Point) int {
	return makeConnections(report, false)
}

func solvePart2(report []Point) int {
	return makeConnections(report, true)
}

// To avoid copy pasting all this, takes a boolean for part 2 logic
// When false it uses 1000 connections and multiplies sizes of 3 largest circuits
// When true it connects until all points are in a single circuit, then returns X coordinates of those two points multiplied
func makeConnections(report []Point, isPart2 bool) int {
	numConnections := 1000
	if isPart2 {
		numConnections = len(report) * (len(report) - 1) / 2 // max possible connections
	}
	// Create array of sorted index pairs, storing the distance squared from origin
	sortedPairs := make([]IndexPair, 0, len(report))
	uniquePairsMap := make(map[string]bool)
	for i, p1 := range report {
		// Calculate distance from this point to all other points
		for j, p2 := range report {
			if i != j {
				// Check for duplicate pairs before adding
				key1 := fmt.Sprintf("%d,%d", i, j)
				key2 := fmt.Sprintf("%d,%d", j, i)
				_, key1Exists := uniquePairsMap[key1]
				_, key2Exists := uniquePairsMap[key2]
				exists := key1Exists || key2Exists
				if !exists {
					distSq := int(math.Pow(float64(p1.x-p2.x), 2) + math.Pow(float64(p1.y-p2.y), 2) + math.Pow(float64(p1.z-p2.z), 2))
					sortedPairs = append(sortedPairs, IndexPair{i, j, distSq})
					uniquePairsMap[key1] = true
				}
			}
		}
	}
	// Sort pairs by distance squared
	sort.Slice(sortedPairs, func(i, j int) bool {
		return sortedPairs[i].distSq < sortedPairs[j].distSq
	})

	// slice of slices of Points representing connected circuits
	// a single slice at index k represents a single circuit
	circuits := make([][]Point, 0)

	for {
		if numConnections <= 0 {
			break
		}
		// Iterate through sorted pairs, connecting points until numConnections is reached
		for _, pair := range sortedPairs {
			fmt.Printf("Connecting pair: %d,%d,%d to %d,%d,%d\n", report[pair.i].x, report[pair.i].y, report[pair.i].z, report[pair.j].x, report[pair.j].y, report[pair.j].z)
			// Check if points are already connected, and grab indices of their circuits if so
			connectedCircuitI, connectedCircuitJ := -1, -1
			for circuitIndex, circuit := range circuits {
				foundI, foundJ := false, false
				for _, point := range circuit {
					if point == report[pair.i] {
						foundI = true
						connectedCircuitI = circuitIndex
					}
					if point == report[pair.j] {
						foundJ = true
						connectedCircuitJ = circuitIndex
					}
				}
				if foundI && foundJ {
					break
				}
			}

			// Connect points pair.i and pair.j
			if connectedCircuitI == -1 && connectedCircuitJ == -1 {
				// Neither point is in a circuit yet, create new circuit
				newCircuit := []Point{report[pair.i], report[pair.j]}
				circuits = append(circuits, newCircuit)
			} else {
				if connectedCircuitI != -1 && connectedCircuitJ == -1 {
					// Add point j to circuit of point i
					circuits[connectedCircuitI] = append(circuits[connectedCircuitI], report[pair.j])
					if isPart2 && len(circuits[connectedCircuitI]) == len(report) {
						// All points are now connected, return product of X coordinates of first two points
						fmt.Printf("All points connected in circuit of size %d\n", len(circuits[connectedCircuitI]))
						return xCoordsProduct(report[pair.i], report[pair.j])
					}
				}
				if connectedCircuitI == -1 && connectedCircuitJ != -1 {
					// Add point i to circuit of point j
					circuits[connectedCircuitJ] = append(circuits[connectedCircuitJ], report[pair.i])
					if isPart2 && len(circuits[connectedCircuitJ]) == len(report) {
						// All points are now connected, return product of X coordinates of first two points
						fmt.Printf("All points connected in circuit of size %d\n", len(circuits[connectedCircuitJ]))
						return xCoordsProduct(report[pair.i], report[pair.j])
					}
				}
				if connectedCircuitI != -1 && connectedCircuitJ != -1 && connectedCircuitI != connectedCircuitJ {
					// Combine circuits
					circuits[connectedCircuitI] = append(circuits[connectedCircuitI], circuits[connectedCircuitJ]...)
					fmt.Printf("Length of new circuit: %d\n", len(circuits[connectedCircuitI]))
					fmt.Println()
					// Check length of new circuit and return if part 2
					if isPart2 && len(circuits[connectedCircuitI]) == len(report) {
						// All points are now connected, return product of X coordinates of first two points
						fmt.Printf("All points connected in circuit of size %d\n", len(circuits[connectedCircuitI]))
						return xCoordsProduct(report[pair.i], report[pair.j])
					}
					// Remove circuit at connectedCircuitJ
					circuits = append(circuits[:connectedCircuitJ], circuits[connectedCircuitJ+1:]...)
				}
			}

			numConnections--
			if numConnections <= 0 {
				break
			}
		}
	}
	circuitSizes := make([]int, len(circuits))
	for i, circuit := range circuits {
		circuitSizes[i] = len(circuit)
	}
	sort.Slice(circuitSizes, func(i, j int) bool {
		return circuitSizes[i] > circuitSizes[j]
	})
	result := 1
	for i := 0; i < 3 && i < len(circuitSizes); i++ {
		fmt.Printf("Circuit %d size: %d\n", i+1, circuitSizes[i])
		result *= circuitSizes[i]
	}
	return result
}

func xCoordsProduct(p1, p2 Point) int {
	return p1.x * p2.x
}
