package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

func runDay(day int) (string, error) {
	filename := constructSolutionPath(day)
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return "", fmt.Errorf("solution for day %d does not exist", day)
	}

	return execSolution(filename, day)
}

func main() {
	r := gin.Default()
	r.SetTrustedProxies(nil)

	// Serve static files
	r.Static("/", "./assets")

	// Define a route to handle POST requests and run the runner
	r.POST("/run", func(c *gin.Context) {
		num := c.PostForm("day")
		intDay, err := strconv.Atoi(num)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		result, err := runDay(intDay)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": result,
		})
	})

	r.Run() // Listen and serve on 0.0.0.0:8080 (for Windows, Linux in Docker)
}

func constructSolutionPath(day int) string {
	path := filepath.Join("..", "solutions", fmt.Sprintf("day%d.go", day))
	return path
}

func execSolution(filename string, day int) (string, error) {
	solutionsPath := filepath.Dir(filename)
	binaryPath := filepath.Join(solutionsPath, fmt.Sprintf("day%d", day))
	// Build solution binary
	buildCmd := exec.Command("go", "build", "-o", binaryPath, filename)
	if err := buildCmd.Run(); err != nil {
		return "", fmt.Errorf("failed to build solution for day %d: %v", day, err)
	}

	// Run the built binary
	runCmd := exec.Command(binaryPath)
	output, err := runCmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to run solution for day %d: %v", day, err)
	}

	fmt.Printf("Output from day %d:\n\n%s", day, string(output))

	// Clean up the built binary
	if err := os.Remove(binaryPath); err != nil {
		fmt.Fprintf(os.Stderr, "failed to remove binary for day %d: %v\n", day, err)
	}

	return string(output), nil
}
