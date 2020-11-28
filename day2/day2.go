package main

import (
	"errors"
	"fmt"

	"github.com/jblashki/aoc-intcode-go"
)

const INPUT_FILE = "./program"
const DAY2B_EXPECTED_OUTPUT = 19690720

func main() {
	var day2aResult int
	var day2bResult int
	var err error

	fmt.Printf("AoC 2019 Day 2 (GO)\n")
	fmt.Printf("-------------------\n")
	day2aResult, err = day2a()
	if err != nil {
		fmt.Printf("2a: **** Error: %q ****\n", err)
	} else {
		fmt.Printf("2a: Program Result = %v\n", day2aResult)
	}
	day2bResult, err = day2b()
	if err != nil {
		fmt.Printf("2b: **** Error: %q ****\n", err)
	} else {
		fmt.Printf("2b: Program Result = %v\n", day2bResult)
	}
}

func day2a() (int, error) {
	ic := intcode.Create()

	err := intcode.Load(ic, INPUT_FILE)
	if err != nil {
		return 0, err
	}

	err = intcode.Set(ic, 1, 12)
	if err != nil {
		return 0, err
	}

	err = intcode.Set(ic, 2, 2)
	if err != nil {
		return 0, err
	}

	var retValue int
	retValue, err = intcode.Run(ic, 0)

	return retValue, err
}

func day2b() (int, error) {
	icOrig := intcode.Create()

	err := intcode.Load(icOrig, INPUT_FILE)
	if err != nil {
		return 0, err
	}

	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			ic := intcode.Copy(icOrig)

			err = intcode.Set(ic, 1, noun)
			if err != nil {
				return 0, err
			}

			err = intcode.Set(ic, 2, verb)
			if err != nil {
				return 0, err
			}

			var value int
			value, err = intcode.Run(ic, 0)
			if err != nil {
				return 0, err
			}

			if value == DAY2B_EXPECTED_OUTPUT {
				return (100 * noun) + verb, nil
			}
		}
	}

	errormsg := fmt.Sprintf("Unable to find value %v", DAY2B_EXPECTED_OUTPUT)
	return 0, errors.New(errormsg)
}
