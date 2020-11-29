package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/jblashki/aoc-filereader-go"
)

const INPUT_PATH = "./orbits"

type planet struct {
	name  string
	orbit *planet
}

func main() {
	var day6aResult int
	var day6bResult int
	var err error

	fmt.Printf("AoC 2019 Day 6 (GO)\n")
	fmt.Printf("-------------------\n")
	day6aResult, err = day6a()
	if err != nil {
		fmt.Printf("6a: **** Error: %v ****\n", err)
	} else {
		fmt.Printf("6a: Total Orbits = %v\n", day6aResult)
	}
	day6bResult, err = day6b()
	if err != nil {
		fmt.Printf("6b: **** Error: %v ****\n", err)
	} else {
		fmt.Printf("6b: Orbital Steps Required = %v\n", day6bResult)
	}
}

func day6a() (int, error) {
	var planetMap = make(map[string]*planet)

	inputs, err := filereader.ReadLines(INPUT_PATH)
	if err != nil {
		return 0, err
	}

	for i := 0; i < len(inputs); i++ {
		p := strings.Split(inputs[i], ")")
		if len(p) != 2 {
			errormsg := fmt.Sprintf(`Error parsing line %v %q: %v`, i, inputs[i], err)
			return 0, errors.New(errormsg)
		}

		err = addOrbit(planetMap, p[1], p[0])
		if err != nil {
			return 0, err
		}
	}

	directOrbits := 0
	indirectOrbits := 0
	for _, p := range planetMap {
		d, i := countOrbits(p)
		directOrbits += d
		indirectOrbits += i
	}

	return directOrbits + indirectOrbits, nil
}

func day6b() (int, error) {
	var planetMap = make(map[string]*planet)

	inputs, err := filereader.ReadLines(INPUT_PATH)
	if err != nil {
		return 0, err
	}

	for i := 0; i < len(inputs); i++ {
		p := strings.Split(inputs[i], ")")
		if len(p) != 2 {
			errormsg := fmt.Sprintf(`Error parsing line %v %q: %v`, i, inputs[i], err)
			return 0, errors.New(errormsg)
		}

		err = addOrbit(planetMap, p[1], p[0])
		if err != nil {
			return 0, err
		}
	}

	you, exists := planetMap["YOU"]
	if !exists {
		return 0, errors.New("Can't find 'YOU'")
	}

	san, exists := planetMap["SAN"]
	if !exists {
		return 0, errors.New("Can't find 'SAN'")
	}

	youSteps := 0
	for youNode := you.orbit; youNode != nil; youNode = youNode.orbit {

		sanSteps := 0
		for sanNode := san.orbit; sanNode != nil; sanNode = sanNode.orbit {
			if sanNode == youNode {
				return youSteps + sanSteps, nil
			}

			sanSteps++
		}

		youSteps++
	}

	return 0, errors.New("Couldn't find common ancestor")
}

func addOrbit(planetMap map[string]*planet, planetName string, orbitName string) error {
	var o *planet
	var p *planet
	var exists bool

	if planetMap == nil {
		return errors.New("Invalid map: nil")
	}

	o, exists = planetMap[orbitName]
	if !exists {
		o = &planet{orbitName, nil}
		planetMap[o.name] = o
	}

	p, exists = planetMap[planetName]
	if !exists {
		p = &planet{planetName, o}
	} else if p.orbit != nil {
		errormsg := fmt.Sprintf("Planet %q is already in orbit around %q", p.name, p.orbit.name)
		return errors.New(errormsg)
	} else {
		p.orbit = o
	}

	planetMap[p.name] = p

	return nil
}

func countOrbits(p *planet) (direct int, indirect int) {
	if p.orbit == nil {
		return
	} else {
		direct = 1
		indirect = -1
	}

	for nextP := p.orbit; nextP != nil; nextP = nextP.orbit {
		indirect++
	}

	return
}
