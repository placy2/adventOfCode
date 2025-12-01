package main

import (
	"fmt"
	"strings"
	"strconv"
	"bufio"
	"os"
)

func main() {
	file, err := os.Open("../data/day3input.txt")
	if err != nil {
			fmt.Println("Error opening file:", err)
			return
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	// multPosition: -1 = not multiplying yet, 0 = no numbers detected, 1 = first number stored, looking for closed parens
	multPosition, firstNum, secondNum, total := -1, 0, 0, 0
	var acc strings.Builder
	for {
		char, _, err := reader.ReadRune()
		if err != nil {
			break
		}
		// Always check for don't
		switch pos := multPosition; pos {
		case -1:
			// check for ( and if found, see if acc is 'mul' otherwise trim first piece of acc off and continue
			checking := acc.String()
			if len(checking) == 4 {
				checking = checking[1:]
				acc.Reset()
				acc.WriteString(checking)
			}
			if char == '(' {
				if acc.String() == "mul" {
					acc.Reset()
					multPosition++
				} else {
					acc.Reset()
				}
			} else {
				acc.WriteString(string(char))
			}
		case 0:
			// just check for comma to see if first num is over, otherwise accumulate
			if char == ',' && isValidNumber(acc.String()) {
				firstNum, _ = strconv.Atoi(acc.String())
				multPosition++
				acc.Reset()
			} else {
				acc.WriteString(string(char))
				if !isValidNumber(acc.String()) {
					acc.Reset()
					multPosition--
				}
			}
		case 1:
			// check for close parens to see if second num is over, otherwise accumulate
			if char == ')' && isValidNumber(acc.String()) {
				secondNum, _ = strconv.Atoi(acc.String())
				multPosition = -1
				acc.Reset()
				total += firstNum * secondNum
			} else {
				acc.WriteString(string(char))
				if !isValidNumber(acc.String()) {
					acc.Reset()
					multPosition = -1
				}
			}
		}
	}

	sample := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
	
	result := multiplyCorruptedLine(sample)
	fmt.Println("Initial multiplication results: ", result)

	//fmt.Println("Initial multiplication results: ", total)
}

// dumb but maybe wortwhile way - just keep track of don'ts and dos and their positions when reading in. 
// then read through again and only do the switch statement as a whole if we're not in a don't range of positions


func multiplyCorruptedLine(s string) int {
	// multPosition: -1 = not multiplying yet, 0 = no numbers detected, 1 = first number stored, looking for closed parens
	// part 2 added -2 value for 'actively in don't mode'
	multPosition, firstNum, secondNum, total := -1, 0, 0, 0
	var acc strings.Builder
	for _, char := range s {
		switch pos := multPosition; pos {
		case -1:
			// check for ( and if found, see if acc is 'mul' otherwise trim first piece of acc off and continue
			// part 2 - keep 5 chars so we can see 'don't' - change check for "mul" to account for this
			checking := acc.String()
			if len(checking) == 5 {
				if char == '(' {
					if checking == "don't" {
						fmt.Println("Found don't, skipping to next parens")
						multPosition = -2
						acc.Reset()
						break
					} else {
						fmt.Println("Found open parens, checking: ", acc.String()[1:])
						if acc.String()[1:] == "mul" {
							acc.Reset()
							multPosition++
						} else {
							acc.Reset()
						}
					}
				} else {
					checking = checking[1:]
					acc.Reset()
					acc.WriteString(checking)
				}
			} else if char == '(' {
				fmt.Println("Found open parens, checking: ", acc.String())
				if acc.String() == "mul" {
					acc.Reset()
					multPosition++
				} else {
					acc.Reset()
				}
			} else {
				acc.WriteString(string(char))
			}
		case 0:
			// just check for comma to see if first num is over, otherwise accumulate
			if char == ',' && isValidNumber(acc.String()) {
				firstNum, _ = strconv.Atoi(acc.String())
				fmt.Println("First num: ", firstNum)
				multPosition++
				acc.Reset()
			} else {
				acc.WriteString(string(char))
				if !isValidNumber(acc.String()) {
					fmt.Println("Invalid number: ", acc.String())
					acc.Reset()
					multPosition--
				}
			}
		case 1:
			// check for close parens to see if second num is over, otherwise accumulate
			if char == ')' && isValidNumber(acc.String()) {
				secondNum, _ = strconv.Atoi(acc.String())
				fmt.Println("Second num: ", secondNum)
				multPosition = -1
				acc.Reset()
				fmt.Println("Multiplying: ", firstNum, secondNum, firstNum * secondNum)
				total += firstNum * secondNum
			} else {
				acc.WriteString(string(char))
				if !isValidNumber(acc.String()) {
					fmt.Println("Invalid number: ", acc.String())
					acc.Reset()
					multPosition = -1
				}
			}
		}
	}
	return total;
}

func isValidNumber(s string) bool {
	if len(s) <= 3 {
		_, err := strconv.Atoi(s)
		return err == nil
	}
	return false
}