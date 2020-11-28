package main

import (
	"fmt"

	"github.com/jblashki/aoc-intcode-go"
)

const INPUT_FILE = "./program"

func main() {
	var err error

	fmt.Printf("AoC 2019 Day 5 (GO)\n")
	fmt.Printf("-------------------\n")
	fmt.Printf("5a: (Enter 1 below)\n")
	err = runProgram()
	if err != nil {
		fmt.Printf("5a: **** Error: %q ****\n", err)
	} else {
		fmt.Printf("5a: See Debug Above\n")
	}
	fmt.Printf("5b: (Enter 5 below)\n")
	err = runProgram()
	if err != nil {
		fmt.Printf("5b: **** Error: %q ****\n", err)
	} else {
		fmt.Printf("5b: See Debug Above\n")
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
