# Advent of Code 2025 - Runner UI

A simple web UI for running Advent of Code solutions.

## Prerequisites

- Go 1.16 or later
- Gin web framework: `go get -u github.com/gin-gonic/gin`

## Project Structure

```
adventOfCode/2025/
├── runner/
│   ├── runner-ui.go
│   ├── assets/        # Static files (HTML, CSS, JS)
│   └── README.md
└── solutions/
    ├── day1.go
    ├── day2.go
    └── ...
```

## Usage

1. Navigate to the runner directory:
   ```bash
   cd adventOfCode/2025/runner
   ```

2. Run the server:
   ```bash
   go run runner-ui.go
   ```

3. Open your browser and navigate to `http://localhost:8080`

4. Enter a day number and click run to execute the corresponding solution

## How It Works

- The UI sends a POST request to `/run` with the day number
- The server builds and executes the corresponding `dayX.go` file from the `solutions` directory
- The output is returned and displayed in the browser
- Built binaries are automatically cleaned up after execution
