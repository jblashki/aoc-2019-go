package day1

import (
	"fmt"

	"github.com/jblashki/aoc-filereader-go"
)

const name = "Day 1"
const INPUT_FILE = "./day1/mass"

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

	mass, err := filereader.ReadAllInts(INPUT_FILE)
	if err != nil {
		return 0, err
	}

	fuel_req := 0
	for i := 0; i < len(mass); i++ {
		fuel_req += calc_fuel_req(mass[i])
	}
	return fuel_req, nil
}

func b() (int, error) {
	mass, err := filereader.ReadAllInts(INPUT_FILE)
	if err != nil {
		return 0, err
	}

	fuel_req := 0
	for i := 0; i < len(mass); i++ {
		fuel_req += calc_fuel_req_recursive(mass[i])
	}
	return fuel_req, nil
}

func calc_fuel_req(mass int) int {
	return ((mass / 3) - 2)
}

func calc_fuel_req_recursive(mass int) int {
	retMass := calc_fuel_req(mass)

	if retMass <= 0 {
		return 0
	} else {
		return retMass + calc_fuel_req_recursive(retMass)
	}
}
