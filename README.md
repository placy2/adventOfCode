# Advent of Code 2024
## Solutions by Parker Lacy

This is a repo containing my solution code for recent years of [Advent of Code.](https://adventofcode.com)

Because the owners of Advent of Code ask that input data (which varies by user) is not shared, this code contains paths to files that are not included in the `/data` folder. 

### Running this code

To run this code yourself, sign up on the site yourself to get your own input data (or sample data from the questions). Then, paste it into the relevant file (i.e., for Day 1, `data/day1input.txt`).

After ensuring the data is where the code expects, run the individual Go file with `go run` (i.e. `go run solutions/day1.go`).

#### Using the runner (2025)
Instead of having to run the specific day file for each of them, there is also a simple runner that you can use: 

Use the number of the day you want to run - it will run `solutions/day{n}.go`
Ensure you are in `2025/runner` to run this.
```console
$ go run main.go 1
Running solution for day 1...
Output from day 1:
...
```