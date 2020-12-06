package day11

import (
	"fmt"
	"sync"

	intcode "github.com/jblashki/aoc-intcode-go/v4"
)

const name = "Day 11"
const inputFile = "./day11/program"

// RunDay runs Advent of Code Day 11 Puzzle
func RunDay(verbose bool) {
	var aResult int
	//var bResult int
	var err error

	if verbose {
		fmt.Printf("\n%v Output:\n", name)
	}

	aResult, err = runRobot()
	if err != nil {
		fmt.Printf("%va: **** Error: %q ****\n", name, err)
	} else {
		fmt.Printf("%va: Program Result = %v\n", name, aResult)
	}

	err = runRobot2(verbose)
	if err != nil {
		fmt.Printf("%vb: **** Error: %q ****\n", name, err)
	} else {
		fmt.Printf("%vb: See Above\n", name)
	}
}

type coord struct {
	x int
	y int
}

type direction int

const (
	dirUp direction = iota
	dirDown
	dirLeft
	dirRight
)

func (d direction) String() string {
	return [...]string{"Up", "Down", "Left", "Right"}[d]
}

func (c coord) String() string {
	return fmt.Sprintf("(%v,%v)", c.x, c.y)
}

type robot struct {
	pos coord
	dir direction
}

func runRobot() (int, error) {
	wg := new(sync.WaitGroup)
	ic, err := intcode.CreateLoad(wg, inputFile, 0, 0)
	if err != nil {
		return 0, err
	}

	defer intcode.Close(ic)

	wg.Add(1)
	go intcode.Run(ic, "")

	// Setup robot here
	r := robot{pos: coord{x: 0, y: 0}, dir: dirUp}
	hullMapVisited := make(map[coord]bool)
	hullMapPaint := make(map[coord]int)
	paintCount := 0
	for {
		halted, err := intcode.Write(ic, hullMapPaint[r.pos])
		if err != nil {
			return 0, err
		} else if halted {
			break
		}

		value, halted, err := intcode.Read(ic)
		if err != nil {
			return 0, err
		} else if halted {
			break
		}
		// value equals colour to paint
		hullMapPaint[r.pos] = value
		if !hullMapVisited[r.pos] {
			hullMapVisited[r.pos] = true
			paintCount++
		}

		value, halted, err = intcode.Read(ic)
		if err != nil {
			// Error
			return 0, err
		} else if halted {
			return 0, fmt.Errorf("Recieved unexpected program halt")
		}
		// value equals instruction to go
		switch value {
		case 0:
			// Left 90 degrees
			switch r.dir {
			case dirUp:
				r.dir = dirLeft
			case dirDown:
				r.dir = dirRight
			case dirLeft:
				r.dir = dirDown
			case dirRight:
				r.dir = dirUp
			}
		case 1:
			// Right 90 degress
			switch r.dir {
			case dirUp:
				r.dir = dirRight
			case dirDown:
				r.dir = dirLeft
			case dirLeft:
				r.dir = dirUp
			case dirRight:
				r.dir = dirDown
			}
		}

		switch r.dir {
		case dirUp:
			r.pos.y++
		case dirDown:
			r.pos.y--
		case dirLeft:
			r.pos.x--
		case dirRight:
			r.pos.x++
		}
	}

	wg.Wait()

	return paintCount, nil
}

func increaseHullMapY(hullMap [][]int, r robot, translate bool) ([][]int, robot) {
	newHullMap := make([][]int, len(hullMap)+1)
	for i := 0; i < len(newHullMap); i++ {
		newHullMap[i] = make([]int, len(hullMap[0]))
		if !translate {
			if i < len(hullMap) {
				copy(newHullMap[i], hullMap[i])
			}
		} else if i > 0 {
			copy(newHullMap[i], hullMap[i-1])
		}
	}

	if translate {
		r.pos.y++
	}

	return newHullMap, r
}

func increaseHullMapX(hullMap [][]int, r robot, translate bool) ([][]int, robot) {
	newHullMap := make([][]int, len(hullMap))
	for i := 0; i < len(newHullMap); i++ {
		newHullMap[i] = make([]int, len(hullMap[i])+1)
		if !translate {
			copy(newHullMap[i], hullMap[i])
		} else {
			for j := 1; j < len(newHullMap[i]); j++ {
				newHullMap[i][j] = hullMap[i][j-1]
			}
		}
	}

	if translate {
		r.pos.x++
	}

	return newHullMap, r
}

func moveRobot(hullMap [][]int, r robot, move int) ([][]int, robot) {
	switch move {
	case 0:
		// Left 90 degrees
		switch r.dir {
		case dirUp:
			r.dir = dirLeft
		case dirDown:
			r.dir = dirRight
		case dirLeft:
			r.dir = dirDown
		case dirRight:
			r.dir = dirUp
		}
	case 1:
		// Right 90 degress
		switch r.dir {
		case dirUp:
			r.dir = dirRight
		case dirDown:
			r.dir = dirLeft
		case dirLeft:
			r.dir = dirUp
		case dirRight:
			r.dir = dirDown
		}
	}

	switch r.dir {
	case dirUp:
		if r.pos.y == 0 {
			// need to increase size of map
			// translate 1 down
			hullMap, r = increaseHullMapY(hullMap, r, true)
		}
		r.pos.y--
	case dirDown:
		if r.pos.y == len(hullMap)-1 {
			// need to increase size of map
			// NO TRANSLATION NEEDED
			hullMap, r = increaseHullMapY(hullMap, r, false)
		}
		r.pos.y++
	case dirLeft:
		if r.pos.x == 0 {
			// need to increase size of map
			// translate 1 right
			hullMap, r = increaseHullMapX(hullMap, r, true)
		}
		r.pos.x--
	case dirRight:
		if r.pos.x == len(hullMap[0])-1 {
			// need to increase size of map
			// NO TRANSLATION NEEDED
			hullMap, r = increaseHullMapX(hullMap, r, false)
		}
		r.pos.x++
	}

	return hullMap, r
}

func drawHullMap(hullMap [][]int, r robot) {
	fmt.Printf("--\n")
	for i := 0; i < len(hullMap); i++ {
		for j := 0; j < len(hullMap[i]); j++ {
			if r.pos.y == i && r.pos.x == j {
				switch r.dir {
				case dirUp:
					fmt.Printf("^")
				case dirDown:
					fmt.Printf("v")
				case dirLeft:
					fmt.Printf("<")
				case dirRight:
					fmt.Printf(">")
				}
			} else {
				if hullMap[i][j] == 1 {
					fmt.Printf("█")
				} else {
					fmt.Printf(" ")
				}
			}
		}
		fmt.Println()
	}
	fmt.Printf("--\n")
}

func runRobot2(verbose bool) error {
	wg := new(sync.WaitGroup)
	ic, err := intcode.CreateLoad(wg, inputFile, 0, 0)
	if err != nil {
		return err
	}

	defer intcode.Close(ic)

	wg.Add(1)
	go intcode.Run(ic, "")

	// Setup robot here
	r := robot{pos: coord{x: 0, y: 0}, dir: dirUp}
	hullMap := make([][]int, 1)
	hullMap[0] = make([]int, 1)
	hullMap[0][0] = 1

	for {
		if verbose {
			drawHullMap(hullMap, r)
		}
		halted, err := intcode.Write(ic, hullMap[r.pos.y][r.pos.x])
		if err != nil {
			return err
		} else if halted {
			break
		}

		value, halted, err := intcode.Read(ic)
		if err != nil {
			return err
		} else if halted {
			break
		}
		// value equals colour to paint
		hullMap[r.pos.y][r.pos.x] = value

		value, halted, err = intcode.Read(ic)
		if err != nil {
			// Error
			return err
		} else if halted {
			return fmt.Errorf("Recieved unexpected program halt")
		}
		// value equals instruction to go
		hullMap, r = moveRobot(hullMap, r, value)
	}

	wg.Wait()

	drawHullMap(hullMap, r)

	return nil
}
