package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: go run main.go <num>")
		os.Exit(1)
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid number: %v\n", err)
		os.Exit(1)
	}

	runDay(n)
}

func runDay(day int) {
	filename := constructSolutionPath(day)
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "solution for day %d does not exist\n", day)
		fmt.Fprintf(os.Stderr, "ensure you are running this from `2025/runner/`\n")
		os.Exit(1)
	}

	fmt.Printf("Running solution for day %d...\n", day)
	execSolution(filename, day)
}

func constructSolutionPath(day int) string {
	path := filepath.Join("..", "solutions", fmt.Sprintf("day%d.go", day))
	return path
}

func execSolution(filename string, day int) {
	solutionsPath := filepath.Dir(filename)
	binaryPath := filepath.Join(solutionsPath, fmt.Sprintf("day%d", day))
	// Build solution binary
	buildCmd := exec.Command("go", "build", "-o", binaryPath, filename)
	if err := buildCmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to build solution for day %d: %v\n", day, err)
		os.Exit(1)
	}

	// Run the built binary
	runCmd := exec.Command(binaryPath)
	output, err := runCmd.CombinedOutput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to run solution for day %d: %v\n", day, err)
		os.Exit(1)
	}

	fmt.Printf("Output from day %d:\n\n%s", day, string(output))

	// Clean up the built binary
	if err := os.Remove(binaryPath); err != nil {
		fmt.Fprintf(os.Stderr, "failed to remove binary for day %d: %v\n", day, err)
	}
}
