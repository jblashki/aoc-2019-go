package day1

import (
	"fmt"

	filereader "github.com/jblashki/aoc-filereader-go"
)

const name = "Day 1"
const inputFile = "./day1/mass"

// RunDay runs Advent of Code Day 4 Puzzle
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
		fmt.Printf("%va: Fuel = %v\n", name, aResult)
	}
	bResult, err = b()
	if err != nil {
		fmt.Printf("%vb: **** Error: %q ****\n", name, err)
	} else {
		fmt.Printf("%vb: Fuel = %v\n", name, bResult)
	}
}

func a() (int, error) {

	mass, err := filereader.ReadAllInts(inputFile)
	if err != nil {
		return 0, err
	}

	fuelReq := 0
	for i := 0; i < len(mass); i++ {
		fuelReq += calcFuelReq(mass[i])
	}
	return fuelReq, nil
}

func b() (int, error) {
	mass, err := filereader.ReadAllInts(inputFile)
	if err != nil {
		return 0, err
	}

	fuelReq := 0
	for i := 0; i < len(mass); i++ {
		fuelReq += calcFuelReqRecursive(mass[i])
	}
	return fuelReq, nil
}

func calcFuelReq(mass int) int {
	return ((mass / 3) - 2)
}

func calcFuelReqRecursive(mass int) int {
	retMass := calcFuelReq(mass)

	if retMass <= 0 {
		return 0
	}

	return retMass + calcFuelReqRecursive(retMass)
}
