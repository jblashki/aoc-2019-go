package day5

import (
	"fmt"

	"github.com/jblashki/aoc-intcode-go"
)

const name = "Day 5"
const INPUT_FILE = "./day5/program"

func RunDay(verbose bool) {
	var err error

	if verbose {
		fmt.Printf("\n%v Output:\n", name)
	}
	fmt.Printf("%va: (Enter 1 below)\n", name)
	err = runProgram()
	if err != nil {
		fmt.Printf("%va: **** Error: %q ****\n", name, err)
	} else {
		fmt.Printf("%va: See Debug Above\n", name)
	}
	fmt.Printf("%vb: (Enter 5 below)\n", name)
	err = runProgram()
	if err != nil {
		fmt.Printf("%vb: **** Error: %q ****\n", name, err)
	} else {
		fmt.Printf("%vb: See Debug Above\n", name)
	}
}

func runProgram() error {
	ic := intcode.Create()

	err := intcode.Load(ic, INPUT_FILE)
	if err != nil {
		return err
	}

	_, err = intcode.Run(ic, 0)

	return err
}
