package main

import (
	"fmt"
)

func main() {
	var day4aResult int
	var day4bResult int
	var err error

	fmt.Printf("AoC 2019 Day 4 (GO)\n")
	fmt.Printf("-------------------\n")
	day4aResult, err = day4a()
	if err != nil {
		fmt.Printf("4a: **** Error: %q ****\n", err)
	} else {
		fmt.Printf("4a: %v passwords match in range\n", day4aResult)
	}
	day4bResult, err = day4b()
	if err != nil {
		fmt.Printf("4b: **** Error: %q ****\n", err)
	} else {
		fmt.Printf("4b: %v passwords match in range\n", day4bResult)
	}
}

func day4a() (int, error) {
	count := 0
	for i := 123257; i <= 647015; i++ {
		if day4aMatch(i) {
			count++
		}
	}
	return count, nil
}

func day4b() (int, error) {
	count := 0
	for i := 123257; i <= 647015; i++ {
		if day4bMatch(i) {
			count++
		}
	}
	return count, nil
}

func day4aMatch(num int) bool {
	pair := false
	lastDigit := -1
	for num > 0 {
		digit := num % 10

		if lastDigit == digit && lastDigit != -1 {
			pair = true
		}

		if lastDigit != -1 && lastDigit < digit {
			return false
		}

		lastDigit = digit
		num /= 10
	}
	return pair
}

func day4bMatch(num int) bool {
	lastDigit := -1
	runCount := 1
	pair := false
	for num > 0 {
		digit := num % 10

		if lastDigit == digit && lastDigit != -1 {
			runCount++
		} else {
			if runCount == 2 {
				pair = true
			}
			runCount = 1
		}

		if lastDigit != -1 && lastDigit < digit {
			return false
		}

		lastDigit = digit
		num /= 10
	}

	if runCount == 2 {
		pair = true
	}

	return pair
}
