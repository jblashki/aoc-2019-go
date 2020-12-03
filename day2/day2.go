package day2

import (
	"errors"
	"fmt"

	intcode "github.com/jblashki/aoc-intcode-go"
)

const name = "Day 2"
const inputFile = "./day2/program"
const expected2BOutput = 19690720

// RunDay runs Advent of Code Day 2 Puzzle
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
		fmt.Printf("%va: Program Result = %v\n", name, aResult)
	}

	bResult, err = b()
	if err != nil {
		fmt.Printf("%vb: **** Error: %q ****\n", name, err)
	} else {
		fmt.Printf("%vb: Program Result = %v\n", name, bResult)
	}
}

func a() (int, error) {
	ic := intcode.Create()

	err := intcode.Load(ic, inputFile)
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

func b() (int, error) {
	icOrig := intcode.Create()

	err := intcode.Load(icOrig, inputFile)
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

			if value == expected2BOutput {
				return (100 * noun) + verb, nil
			}
		}
	}

	errormsg := fmt.Sprintf("Unable to find value %v", expected2BOutput)
	return 0, errors.New(errormsg)
}
