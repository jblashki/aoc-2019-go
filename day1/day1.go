package main

import (
	"fmt"

	"github.com/jblashki/aoc-filereader-go"
)

const INPUT_FILE = "./mass"

func main() {
	var day1aResult int
	var day1bResult int
	var err error

	fmt.Printf("AoC 2019 Day 1 (GO)\n")
	fmt.Printf("-------------------\n")
	day1aResult, err = day1a()
	if err != nil {
		fmt.Printf("1a: **** Error: %q ****\n", err)
	} else {
		fmt.Printf("1a: Fuel = %v\n", day1aResult)
	}
	day1bResult, err = day1b()
	if err != nil {
		fmt.Printf("1b: **** Error: %q ****\n", err)
	} else {
		fmt.Printf("1b: Fuel = %v\n", day1bResult)
	}
}

func day1a() (int, error) {

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

func day1b() (int, error) {
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
