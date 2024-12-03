package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {
	sample := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	
	result := multiplyCorruptedLine(sample)
	fmt.Println("Initial multiplication results: ", result)
}

func multiplyCorruptedLine(s string) int {
	// multPosition: -1 = not multiplying yet, 0 = no numbers detected, 1 = first number stored, looking for closed parens
	multPosition, firstNum, secondNum, total := -1, 0, 0, 0
	var acc strings.Builder
	for _, char := range s {
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