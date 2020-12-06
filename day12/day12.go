package day12

import (
	"fmt"

	filereader "github.com/jblashki/aoc-filereader-go"
)

const name = "Day 12"
const inputFile = "./day12/planets"

type coord struct {
	x int
	y int
	z int
}

type dimension int

const (
	dimensionX dimension = iota
	dimensionY
	dimensionZ
)

func (d dimension) String() string {
	return [...]string{"X", "Y", "Z"}[d]
}

type velocity coord

func (c coord) String() string {
	return fmt.Sprintf("(x: %v, y: %v, z: %v)", c.x, c.y, c.z)
}

func (v velocity) String() string {
	return fmt.Sprintf("<x: %v, y: %v, z: %v>", v.x, v.y, v.z)
}

type planet struct {
	pos coord
	vel velocity
}

func (p planet) String() string {
	return fmt.Sprintf("P%v - V%v", p.pos, p.vel)
}

// RunDay runs Advent of Code Day 12 Puzzle
func RunDay(verbose bool) {
	var aResult int
	var bResult int
	var err error

	if verbose {
		fmt.Printf("\n%v Output:\n", name)
	}

	aResult, err = a(verbose)
	if err != nil {
		fmt.Printf("%va: **** Error: %q ****\n", name, err)
	} else {
		fmt.Printf("%va: Program Result = %v\n", name, aResult)
	}

	bResult, err = b(verbose)
	if err != nil {
		fmt.Printf("%vb: **** Error: %q ****\n", name, err)
	} else {
		fmt.Printf("%vb: Program Result = %v\n", name, bResult)
	}
}

func applyTimeStep(planets []planet) {
	for i := 0; i < len(planets); i++ {
		p1 := &planets[i]
		if i < len(planets)-1 {
			for j := i + 1; j < len(planets); j++ {
				p2 := &planets[j]

				// X velocity
				switch {
				case p1.pos.x > p2.pos.x:
					p1.vel.x--
					p2.vel.x++

				case p1.pos.x < p2.pos.x:
					p1.vel.x++
					p2.vel.x--
				}

				// Y velocity
				switch {
				case p1.pos.y > p2.pos.y:
					p1.vel.y--
					p2.vel.y++

				case p1.pos.y < p2.pos.y:
					p1.vel.y++
					p2.vel.y--
				}

				// Z velocity
				switch {
				case p1.pos.z > p2.pos.z:
					p1.vel.z--
					p2.vel.z++

				case p1.pos.z < p2.pos.z:
					p1.vel.z++
					p2.vel.z--
				}
			}
		}
		// Adjust velocity here
		p1.pos.x += p1.vel.x
		p1.pos.y += p1.vel.y
		p1.pos.z += p1.vel.z
	}
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}

	return a
}

func calcEnergy(p planet, verbose bool) int {
	pEnergy := abs(p.pos.x) + abs(p.pos.y) + abs(p.pos.z)
	kEnergy := abs(p.vel.x) + abs(p.vel.y) + abs(p.vel.z)
	tEnergy := pEnergy * kEnergy

	if verbose {
		fmt.Printf("P: %v + %v + %v = %v; K: %v + %v + %v = %v; Total: %v * %v = %v\n\n",
			abs(p.pos.x), abs(p.pos.y), abs(p.pos.z), pEnergy,
			abs(p.vel.x), abs(p.vel.y), abs(p.vel.z), kEnergy,
			pEnergy, kEnergy, tEnergy)
	}

	return tEnergy
}

func coordEqu(c1 coord, c2 coord) bool {
	if c1.x != c2.x {
		return false
	}

	if c1.y != c2.y {
		return false
	}

	if c1.z != c2.z {
		return false
	}
	return true
}

func velocityEqu(v1 velocity, v2 velocity) bool {
	if v1.x != v2.x {
		return false
	}

	if v1.y != v2.y {
		return false
	}

	if v1.z != v2.z {
		return false
	}
	return true
}

func planetEqu(p1 planet, p2 planet) bool {
	if !coordEqu(p1.pos, p2.pos) {
		return false
	}

	if !velocityEqu(p1.vel, p2.vel) {
		return false
	}
	return true
}

func planetsSliceEqu(p1 []planet, p2 []planet) bool {
	if len(p1) != len(p2) {
		return false
	}

	for i := 0; i < len(p1); i++ {
		if !planetEqu(p1[i], p2[i]) {
			return false
		}
	}

	return true
}

func planetsSliceEquDimension(p1 []planet, p2 []planet, d dimension) bool {
	if len(p1) != len(p2) {
		return false
	}

	for i := 0; i < len(p1); i++ {
		planet1 := p1[i]
		planet2 := p2[i]

		switch d {
		case dimensionX:
			if planet1.pos.x != planet2.pos.x || planet1.vel.x != planet2.vel.x {
				return false
			}
		case dimensionY:
			if planet1.pos.y != planet2.pos.y || planet1.vel.y != planet2.vel.y {
				return false
			}

		case dimensionZ:
			if planet1.pos.z != planet2.pos.z || planet1.vel.z != planet2.vel.z {
				return false
			}
		}
	}

	return true
}

func getMaxSharedFactor(a int, b int) int {
	if a < 0 {
		a *= -1
	}
	if b < 0 {
		b *= -1
	}

	min := a
	if b < a {
		min = b
	}

	if a == 0 {
		return b
	} else if b == 0 {
		return a
	}

	for i := min; i > 1; i-- {
		if a%i == 0 && b%i == 0 {
			return i
		}
	}

	return 1
}

func getCycleNumber(a int, b int) int {
	gcd := getMaxSharedFactor(a, b)
	returnValue := abs(a) * abs(b)
	returnValue /= gcd
	return returnValue
}

func readPlanets(file string) ([]planet, error) {
	returnPlanets := make([]planet, 0)

	lines, err := filereader.ReadLines(file)
	if err != nil {
		return nil, err
	}

	for _, line := range lines {
		xVal := 0
		yVal := 0
		zVal := 0
		fmt.Sscanf(line, "<x=%d, y=%d, z=%d>", &xVal, &yVal, &zVal)
		p := planet{
			pos: coord{x: xVal, y: yVal, z: zVal},
			vel: velocity{x: 0, y: 0, z: 0},
		}
		returnPlanets = append(returnPlanets, p)
	}

	return returnPlanets, nil
}

func a(verbose bool) (int, error) {
	planets, err := readPlanets(inputFile)
	if err != nil {
		return 0, err
	}

	if verbose {
		for i := 0; i < len(planets); i++ {
			fmt.Printf("%v\n", planets[i])
		}
		fmt.Println()
	}

	for i := 0; i < 1000; i++ {
		applyTimeStep(planets)
	}

	if verbose {
		for i := 0; i < len(planets); i++ {
			fmt.Printf("%v\n", planets[i])
		}
		fmt.Println()
	}

	totalEnergy := 0
	for i := 0; i < len(planets); i++ {
		totalEnergy += calcEnergy(planets[i], verbose)
	}

	return totalEnergy, nil
}

func b(verbose bool) (int, error) {
	planets, err := readPlanets(inputFile)
	if err != nil {
		return 0, err
	}

	origPlanets := make([]planet, len(planets))
	copy(origPlanets, planets)

	xCycle := 0
	yCycle := 0
	zCycle := 0
	steps := 0
	if verbose {
		fmt.Println()
	}
	for {
		applyTimeStep(planets)
		steps++
		if planetsSliceEquDimension(planets, origPlanets, dimensionX) {
			if xCycle == 0 {
				xCycle = steps
				if verbose {
					fmt.Printf("Found x Cycle after %v steps\n", xCycle)
				}
			}
		}

		if planetsSliceEquDimension(planets, origPlanets, dimensionY) {
			if yCycle == 0 {
				yCycle = steps
				if verbose {
					fmt.Printf("Found y Cycle after %v steps\n", yCycle)
				}
			}
		}

		if planetsSliceEquDimension(planets, origPlanets, dimensionZ) {
			if zCycle == 0 {
				zCycle = steps
				if verbose {
					fmt.Printf("Found Z Cycle after %v steps\n", zCycle)
				}
			}
		}

		if xCycle != 0 && yCycle != 0 && zCycle != 0 {
			break
		}
	}

	if verbose {
		fmt.Println()
	}

	cycle := getCycleNumber(xCycle, yCycle)
	cycle = getCycleNumber(cycle, zCycle)

	return cycle, nil
}
