package main

import (
	"bufio"
	"os"
)

type Cell struct {
	r, c int
}

func main() {
	grid := readInput()
	// result1 := solvePart1(grid)
	// println("Part 1 result:", result1)
	result2 := solvePart2(grid)
	println("Part 2 result:", result2)
}

func readInput() [][]rune {
	file, err := os.Open("../data/day4input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close() // close file when done, executes after the rest of the parent function ends

	scanner := bufio.NewScanner(file)
	var grid [][]rune
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return grid
}

func solvePart1(grid [][]rune) int {
	memo := make(map[Cell]bool)
	rows := len(grid)
	cols := len(grid[0])
	count := 0

	directions := []Cell{
		{-1, 0},  // up
		{-1, -1}, // up-left
		{-1, 1},  // up-right
		{0, -1},  // left
		{0, 1},   // right
		{1, 0},   // down
		{1, -1},  // down-left
		{1, 1},   // down-right
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			cell := Cell{r, c}
			if grid[r][c] == '.' {
				continue
			}
			if isAccessible(cell, grid, memo, directions, rows, cols) {
				count++
			}
		}
	}

	return count
}

func solvePart2(grid [][]rune) int {
	totalRemoved := 0
	// Loop until no more can be removed
	for {
		removed, newGrid := solveAndRemove(grid)
		if removed == 0 {
			break
		}
		totalRemoved += removed
		grid = newGrid
	}

	return totalRemoved
}

func solveAndRemove(grid [][]rune) (int, [][]rune) {
	// Lazy copying of same logic to avoid passing around memo/new grid
	memo := make(map[Cell]bool)
	rows := len(grid)
	cols := len(grid[0])
	count := 0

	directions := []Cell{
		{-1, 0},  // up
		{-1, -1}, // up-left
		{-1, 1},  // up-right
		{0, -1},  // left
		{0, 1},   // right
		{1, 0},   // down
		{1, -1},  // down-left
		{1, 1},   // down-right
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			cell := Cell{r, c}
			if grid[r][c] == '.' {
				continue
			}
			if isAccessible(cell, grid, memo, directions, rows, cols) {
				count++
			}
		}
	}

	// can't remove while still in grid, so just loop through again to remove papers
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			cell := Cell{r, c}
			if grid[r][c] == '@' {
				if memo[cell] {
					grid[r][c] = '.'
				}
			}
		}
	}

	return count, grid
}

func isAccessible(cell Cell, grid [][]rune, memo map[Cell]bool, directions []Cell, rows, cols int) bool {
	// check memo first
	if val, exists := memo[cell]; exists {
		return val
	}

	paperCount := 0
	for _, dir := range directions {
		nr, nc := cell.r+dir.r, cell.c+dir.c
		// Prevent out of bound cells from being checked
		if nr >= 0 && nr < rows && nc >= 0 && nc < cols {
			if grid[nr][nc] == '@' {
				paperCount++
				if paperCount >= 4 {
					memo[cell] = false
					return false
				}
			}
		}
	}

	memo[cell] = paperCount < 4
	//fmt.Printf("Cell (%d,%d) accessible: %v\n", cell.r, cell.c, memo[cell])
	return memo[cell]
}
