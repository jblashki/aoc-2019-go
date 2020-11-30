package day3

import (
	"errors"
	"fmt"

	"aoc2019/day3/wirell"
	"github.com/jblashki/aoc-filereader-go"
)

const name = "Day 3"
const INPUT_FILE = "./day3/wire_map"

func RunDay(verbose bool) {
	var aDist int
	var aPoint wirell.WirePos
	var bSteps int
	var bPoint wirell.WirePos
	var err error

	if verbose {
		fmt.Printf("%v Output:\n", name)
	}
	aDist, aPoint, bSteps, bPoint, err = run(verbose)
	if err != nil {
		fmt.Printf("%va: **** Error: %q ****\n", name, err)
		fmt.Printf("%vb: **** Error: %q ****\n", name, err)
	} else {
		fmt.Printf("%va: Closest crossover = %v @ %v\n", name, aDist, aPoint)
		fmt.Printf("%vb: Least Steps       = %v @ %v\n", name, bSteps, bPoint)
	}
}
func run(verbose bool) (int, wirell.WirePos, int, wirell.WirePos, error) {
	ll1, err := wirell.CreateWireLL()
	if err != nil {
		errormsg := fmt.Sprintf("Failed to create Wire1 linked list: %q", err)
		return 0, wirell.WirePos{0, 0}, 0, wirell.WirePos{0, 0}, errors.New(errormsg)
	}

	ll2, err := wirell.CreateWireLL()
	if err != nil {
		errormsg := fmt.Sprintf("Failed to create Wire1 linked list: %q", err)
		return 0, wirell.WirePos{0, 0}, 0, wirell.WirePos{0, 0}, errors.New(errormsg)
	}

	wires, err := filereader.ReadCSVStringsPerLine(INPUT_FILE)
	if err != nil {
		return 0, wirell.WirePos{0, 0}, 0, wirell.WirePos{0, 0}, err
	}

	wire1 := wires[0]
	wire2 := wires[1]

	if verbose {
		fmt.Printf("Loading Wire1...")
	}
	for i := 0; i < len(wire1); i++ {
		move, err := wirell.ParseMove(wire1[i])
		if err != nil {
			errormsg := fmt.Sprintf("Invalid move %q @ line 1 field %v: %v\n", wire1[i], i, err)
			return 0, wirell.WirePos{0, 0}, 0, wirell.WirePos{0, 0}, errors.New(errormsg)
		}

		err = wirell.AddMove(ll1, move)
		if err != nil {
			errormsg := fmt.Sprintf("Failed to make move %q @ line 1 field %v: %v\n", move, i, err)
			return 0, wirell.WirePos{0, 0}, 0, wirell.WirePos{0, 0}, errors.New(errormsg)
		}
	}
	if verbose {
		fmt.Printf("DONE\n")
		fmt.Printf("Loading Wire2...")
	}
	for i := 0; i < len(wire2); i++ {
		move, err := wirell.ParseMove(wire2[i])
		if err != nil {
			errormsg := fmt.Sprintf("Invalid move %q @ line 2 field %v: %v\n", wire2[i], i, err)
			return 0, wirell.WirePos{0, 0}, 0, wirell.WirePos{0, 0}, errors.New(errormsg)
		}
		err = wirell.AddMove(ll2, move)
		if err != nil {
			errormsg := fmt.Sprintf("Failed to make move %q @ line 2 field %v: %v\n", move, i, err)
			return 0, wirell.WirePos{0, 0}, 0, wirell.WirePos{0, 0}, errors.New(errormsg)
		}
	}
	if verbose {
		fmt.Printf("DONE\n")
	}

	if verbose {
		fmt.Printf("Finding intersections...")
	}
	intersections := wirell.FindCrossovers(ll1, ll2)
	if verbose {
		fmt.Printf("DONE\n")
	}

	minDistance := -1
	minPoint := wirell.WirePos{-1, -1}
	minSteps := -1
	minStepsPoint := wirell.WirePos{-1, -1}
	for i := 0; i < len(intersections); i++ {
		newDistance := 0

		if intersections[i].XPos < 0 {
			newDistance += intersections[i].XPos * -1
		} else {
			newDistance += intersections[i].XPos
		}

		if intersections[i].YPos < 0 {
			newDistance += intersections[i].YPos * -1
		} else {
			newDistance += intersections[i].YPos
		}

		if minDistance == -1 || newDistance < minDistance {
			minDistance = newDistance
			minPoint = intersections[i]
		}

		steps1, _ := wirell.FindStepsToPos(ll1, intersections[i])
		steps2, _ := wirell.FindStepsToPos(ll2, intersections[i])

		steps := steps1 + steps2

		if minSteps == -1 || steps < minSteps {
			minSteps = steps
			minStepsPoint = intersections[i]
		}

	}

	return minDistance, minPoint, minSteps, minStepsPoint, nil
}
