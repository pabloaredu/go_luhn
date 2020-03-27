package main

import (
	"fmt"
	"strconv"
)

func main() {
	// TEST EXAMPLES
	fmt.Println("true = ", verifySinNumber("046454286"))
	fmt.Println("false = ", verifySinNumber("000000000"))
	fmt.Println("false = ", verifySinNumber("046A54286"))
	fmt.Println("false = ", verifySinNumber("046474286"))
}

// Receive a string and return a bool
func verifySinNumber(sin string) bool {
	if len(sin) != 9 {
		return false
	}

	sinNumber, _ := strconv.Atoi(sin)
	if sumDigits(sinNumber) == 0 {
		return false
	}

	var digit int
	var sumOfDigits int

	for i := 0; i < len(sin); i++ {
		if i%2 != 0 {
			digit, _ = strconv.Atoi(sin[i : i+1])
			addValue := digit * 2
			if s := strconv.Itoa(addValue); len(s) > 1 {
				v1, _ := strconv.Atoi(s[0:1])
				v2, _ := strconv.Atoi(s[1:2])
				addValue = v1 + v2
			}
			sumOfDigits += addValue

		}
		if i%2 == 0 {
			digit, _ = strconv.Atoi(sin[i : i+1])
			sumOfDigits += digit
		}
	}

	return sumOfDigits%10 == 0
}

// HELPER FUNCTIONS
func sumDigits(number int) int {
	var remainder int
	var sum int
	for number != 0 {
		remainder = number % 10
		sum += remainder
		number = number / 10
	}
	return sum
}
