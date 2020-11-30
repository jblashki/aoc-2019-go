package day4

import (
	"fmt"
)

const name = "Day 4"

func RunDay(verbose bool) {
	var aResult int
	var bResult int
	var err error

	if verbose {
		fmt.Printf("\n%v Output:\n", name)
	}

	aResult, err = a()
	if err != nil {
		fmt.Printf("%va: **** Error: %q ****\n", name, err)
	} else {
		fmt.Printf("%va: %v passwords match in range\n", name, aResult)
	}

	bResult, err = b()
	if err != nil {
		fmt.Printf("%vb: **** Error: %q ****\n", name, err)
	} else {
		fmt.Printf("%vb: %v passwords match in range\n", name, bResult)
	}
}

func a() (int, error) {
	count := 0
	for i := 123257; i <= 647015; i++ {
		if aMatch(i) {
			count++
		}
	}
	return count, nil
}

func b() (int, error) {
	count := 0
	for i := 123257; i <= 647015; i++ {
		if bMatch(i) {
			count++
		}
	}
	return count, nil
}

func aMatch(num int) bool {
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

func bMatch(num int) bool {
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
