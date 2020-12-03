package day3

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	filereader "github.com/jblashki/aoc-filereader-go"
)

const name = "Day 3"
const inputFile = "./day3/wire_map"

type node struct {
	wire1         bool
	wire1MinSteps int
	wire2         bool
	wire2MinSteps int
}

type direction int

const (
	wireUp = iota
	wireDown
	wireLeft
	wireRight
)

type move struct {
	dir  direction
	dist int
}

// RunDay runs Advent of Code Day 3 Puzzle
func RunDay(verbose bool) {
	var aDist int
	var bSteps int
	var err error

	if verbose {
		fmt.Printf("\n%v Output:\n", name)
	}
	aDist, bSteps, err = calc(verbose)
	if err != nil {
		fmt.Printf("%va: **** Error: %q ****\n", name, err)
		fmt.Printf("%vb: **** Error: %q ****\n", name, err)
	} else {
		fmt.Printf("%va: Closest crossover = %v\n", name, aDist)
		fmt.Printf("%vb: Least Steps       = %v\n", name, bSteps)
	}
}

func calc(verbose bool) (int, int, error) {
	wires, err := filereader.ReadCSVStringsPerLine(inputFile)
	if err != nil {
		return 0, 0, err
	}
	wire1 := wires[0]
	wire2 := wires[1]

	xMax, yMax, sourceX, sourceY, err := calcWireMapDimensions(wires)
	if err != nil {
		return 0, 0, err
	}

	minDist := -1
	minDistX := -1
	minDistY := -1

	minSteps := -1
	minStepsX := -1
	minStepsY := -1

	wireMap := make([][]node, xMax)
	for i := 0; i < xMax; i++ {
		wireMap[i] = make([]node, yMax)
	}

	expectedWire1Steps := 0
	expectedWire2Steps := 0

	curWireX := sourceX
	curWireY := sourceY
	wireMap[curWireX][curWireY].wire1 = true
	for i := 0; i < len(wire1); i++ {
		m, err := parseMove(wire1[i])
		if err != nil {
			return 0, 0, err
		}

		incX := 0
		incY := 0
		switch m.dir {
		case wireUp:
			incY = 1

		case wireDown:
			incY = -1

		case wireLeft:
			incX = -1

		case wireRight:
			incX = 1
		}

		for j := 0; j < m.dist; j++ {
			expectedWire1Steps++

			curWireX += incX
			curWireY += incY
			if wireMap[curWireX][curWireY].wire1MinSteps == 0 {
				wireMap[curWireX][curWireY].wire1MinSteps = expectedWire1Steps
			}
			wireMap[curWireX][curWireY].wire1 = true
		}
	}

	curWireX = sourceX
	curWireY = sourceY
	wireMap[curWireX][curWireY].wire2 = true
	for i := 0; i < len(wire2); i++ {
		m, err := parseMove(wire2[i])
		if err != nil {
			return 0, 0, err
		}

		incX := 0
		incY := 0
		switch m.dir {
		case wireUp:
			incY = 1

		case wireDown:
			incY = -1

		case wireLeft:
			incX = -1

		case wireRight:
			incX = 1
		}

		for j := 0; j < m.dist; j++ {
			expectedWire2Steps++
			curWireX += incX
			curWireY += incY
			wireMap[curWireX][curWireY].wire2 = true
			if wireMap[curWireX][curWireY].wire2MinSteps == 0 {
				wireMap[curWireX][curWireY].wire2MinSteps = expectedWire2Steps
			}
			if wireMap[curWireX][curWireY].wire1 == true {
				// Intersection
				distX := curWireX - sourceX
				distY := curWireY - sourceY
				dist := 0
				if distX < 0 {
					dist += distX * -1
				} else {
					dist += distX
				}
				if distY < 0 {
					dist += distY * -1
				} else {
					dist += distY
				}

				if dist < minDist || minDist == -1 {
					minDist = dist
					minDistX = curWireX
					minDistY = curWireY
				}

				totalSteps := wireMap[curWireX][curWireY].wire1MinSteps + wireMap[curWireX][curWireY].wire2MinSteps

				if totalSteps < minSteps || minSteps == -1 {
					minSteps = totalSteps
					minStepsX = curWireX
					minStepsY = curWireY

				}
			}
		}
	}

	if verbose {
		fmt.Printf("Found Closest Intersection @ (%v, %v) = %v\n", minDistX, minDistY, minDist)
		fmt.Printf("Found Closest Steps        @ (%v, %v) = %v\n", minStepsX, minStepsY, minSteps)
	}

	return minDist, minSteps, nil
}

func getMinMax(wireInst []string, minX *int, maxX *int, minY *int, maxY *int) error {
	curX := 0
	curY := 0
	for i := 0; i < len(wireInst); i++ {
		m, err := parseMove(wireInst[i])
		if err != nil {
			return err
		}

		switch m.dir {
		case wireUp:
			curY += m.dist

		case wireDown:
			curY -= m.dist

		case wireLeft:
			curX -= m.dist

		case wireRight:
			curX += m.dist
		}

		if curX < *minX {
			*minX = curX
		}

		if curX > *maxX {
			*maxX = curX
		}

		if curY < *minY {
			*minY = curY
		}

		if curY > *maxY {
			*maxY = curY
		}
	}

	return nil
}

func calcWireMapDimensions(wireInstructions [][]string) (xMax int, yMax int, sourceX int, sourceY int, err error) {
	err = nil
	wire1 := wireInstructions[0]
	wire2 := wireInstructions[1]

	minX := 0
	maxX := 0
	minY := 0
	maxY := 0

	err = getMinMax(wire1, &minX, &maxX, &minY, &maxY)
	if err != nil {
		return
	}

	err = getMinMax(wire2, &minX, &maxX, &minY, &maxY)
	if err != nil {
		return
	}

	xMax = maxX - minX
	yMax = maxY - minY

	if minX < 0 {
		sourceX = minX * -1
	}

	if minY < 0 {
		sourceY = minY * -1
	}

	xMax++
	yMax++

	return
}

func parseMove(s string) (move, error) {
	m := move{wireUp, 0}
	if s == "" {
		errormsg := fmt.Sprintf("Parsing error. Invalid value %q", s)
		return m, errors.New(errormsg)
	}

	switch s[0] {
	case 'u':
		fallthrough
	case 'U':
		m.dir = wireUp

	case 'd':
		fallthrough
	case 'D':
		m.dir = wireDown

	case 'l':
		fallthrough
	case 'L':
		m.dir = wireLeft

	case 'r':
		fallthrough
	case 'R':
		m.dir = wireRight

	default:
		errormsg := fmt.Sprintf("Invalid direction '%c' in %q", s[0], s)
		return m, errors.New(errormsg)
	}

	var err error
	m.dist, err = strconv.Atoi(strings.TrimSpace(s[1:]))
	if err != nil {
		return m, err
	}

	return m, nil
}
