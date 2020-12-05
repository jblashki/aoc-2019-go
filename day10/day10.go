package day10

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"sort"

	filereader "github.com/jblashki/aoc-filereader-go"
)

const name = "Day 10"
const inputFile = "./day10/map"

const asteroidChar = '#'
const emptyChar = '.'
const blockedChar = '*'
const sourceChar = '@'
const destroyedChar = '*'

type coord struct {
	x int
	y int
}

// RunDay runs Advent of Code Day 10 Puzzle
func RunDay(verbose bool) {
	var aMaxCount int
	var aMaxCountCoord coord
	var bResult int
	var err error

	if verbose {
		fmt.Printf("\n%v Output:\n", name)
	}

	aMaxCount, aMaxCountCoord, err = a(verbose)
	if err != nil {
		fmt.Printf("%va: **** Error: %q ****\n", name, err)
	} else {
		fmt.Printf("%va: %v asteroids can be seen from %v\n", name, aMaxCount, aMaxCountCoord)
	}

	bResult, err = b(aMaxCountCoord, verbose)
	if err != nil {
		fmt.Printf("%vb: **** Error: %q ****\n", name, err)
	} else {
		fmt.Printf("%vb: Program Result = %v\n", name, bResult)
	}

}

func shoot(asteroidMap []rune, w int, source coord, dir coord) (coord, error) {
	cur := coordAdd(source, dir)
	for {
		if cur.x < 0 || cur.x >= w || cur.y < 0 || cur.y >= len(asteroidMap)/w {
			break
		}
		idx := coordToIdx(cur, w)

		if asteroidMap[idx] == asteroidChar {
			return cur, nil
		}
		cur = coordAdd(cur, dir)
	}
	return coord{-1, -1}, errors.New("No Hit")
}

func pressEnter() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("<PRESS ENTER TO CONTINUE>")
	reader.ReadString('\n')
}

func b(source coord, verbose bool) (int, error) {
	asteroidMap, w, err := readAsteroidMap(inputFile)
	if err != nil {
		return -1, err
	}

	idx := coordToIdx(source, w)
	asteroidMap[idx] = sourceChar

	seenAngles := map[float64]bool{}
	angleToCoord := map[float64]coord{}
	angles := make([]float64, 0)
	for i := 0; i < len(asteroidMap); i++ {
		target := idxToCoord(i, w)
		dist := coord{x: target.x - source.x, y: target.y - source.y}
		theta := float64(0)
		if dist.x == 0 {
			switch {
			case dist.y > 0:
				theta = math.Pi
			case dist.y < 0:
				theta = 0
			}
		} else {
			ror := float64(dist.y) / float64(dist.x)
			theta = math.Atan(ror)
			switch {
			case dist.x < 0:
				theta += (3 * math.Pi) / 2
			case dist.x > 0:
				theta += math.Pi / 2
			}
		}
		if !seenAngles[theta] {
			seenAngles[theta] = true
			factor := getMaxSharedFactor(dist.x, dist.y)
			dist := coordDiv(dist, factor)
			angleToCoord[theta] = dist
			angles = append(angles, theta)
		}

	}

	sort.Float64s(angles)

	done := false
	asteroidsDestroyed := 0
	cycle := 0
	for !done {
		destroyedInCycle := 0
		for _, theta := range angles {
			c := angleToCoord[theta]
			shot, err := shoot(asteroidMap, w, source, c)
			if err == nil {
				if verbose {
					fmt.Printf("%v) Shot @ %v\n", asteroidsDestroyed+1, shot)
				}
				idx := coordToIdx(shot, w)
				asteroidMap[idx] = destroyedChar
				if verbose {
					drawAsteroidMap(asteroidMap, w)
					pressEnter()
				}
				destroyedInCycle++
				asteroidsDestroyed++
				if asteroidsDestroyed == 200 {
					retValue := (shot.x * 100) + shot.y
					return retValue, nil
				}
			}
		}
		if destroyedInCycle == 0 {
			done = true
		}
		cycle++
	}

	errormsg := fmt.Sprintf("Only destroyed %v asteroids\n", asteroidsDestroyed)
	return -1, errors.New(errormsg)
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

func blocked(asteroidMap []rune, w int, source coord, target coord) (coord, error) {
	dist := coord{x: target.x - source.x, y: target.y - source.y}

	maxFactor := getMaxSharedFactor(dist.x, dist.y)

	inc := coordDiv(dist, maxFactor)

	pos := coordAdd(source, inc)

	for !coordEqual(pos, target) {
		idx := coordToIdx(pos, w)

		if asteroidMap[idx] == asteroidChar {
			return pos, nil
		}

		pos = coordAdd(pos, inc)
	}
	return coord{x: -1, y: -1}, errors.New("Not Blocked")
}

func blockAsteroids(asteroidMap []rune, w int, c coord) {
	idx := coordToIdx(c, w)

	if asteroidMap[idx] != '@' {
		return
	}

	for i := 0; i < len(asteroidMap); i++ {
		pos := idxToCoord(i, w)

		if !coordEqual(pos, c) && asteroidMap[i] == asteroidChar {
			_, err := blocked(asteroidMap, w, c, pos)
			if err == nil {
				asteroidMap[i] = blockedChar
			}
		}
	}
}

func countAsteroids(asteroidMap []rune) (count int) {
	for i := 0; i < len(asteroidMap); i++ {
		if asteroidMap[i] == asteroidChar {
			count++
		}
	}
	return
}

func readAsteroidMap(file string) ([]rune, int, error) {
	mapString := ""
	rowLength := -1
	lines, err := filereader.ReadLines(file)
	if err != nil {
		return []rune(""), 0, err
	}
	for i := 0; i < len(lines); i++ {
		rowLength = len(lines[i])
		mapString += lines[i]
	}

	return []rune(mapString), rowLength, nil
}

func a(verbose bool) (int, coord, error) {
	asteroidMap, w, err := readAsteroidMap(inputFile)
	if err != nil {
		return 0, idxToCoord(0, w), err
	}

	maxCount := -1
	maxCountIdx := -1
	for i := 0; i < len(asteroidMap); i++ {
		c := idxToCoord(i, w)

		if asteroidMap[i] == asteroidChar {
			var tmpMap = make([]rune, len(asteroidMap))
			copy(tmpMap, asteroidMap)

			tmpMap[i] = '@'

			if verbose {
				fmt.Printf("Postion %v:", c)
				for j := 0; j < len(tmpMap); j++ {
					if j%w == 0 {
						fmt.Println()
					}
					fmt.Printf("%c", tmpMap[j])
				}

				fmt.Printf("\n\n")
			}
			blockAsteroids(tmpMap, w, c)

			count := countAsteroids(tmpMap)
			if maxCount == -1 || count > maxCount {
				maxCount = count
				maxCountIdx = i
			}

			if verbose {

				fmt.Printf("After:")
				for j := 0; j < len(tmpMap); j++ {
					if j%w == 0 {
						fmt.Println()
					}
					fmt.Printf("%c", tmpMap[j])
				}

				fmt.Printf("\n\n")
			}
		}
	}

	return maxCount, idxToCoord(maxCountIdx, w), nil
}

func (c coord) String() string {
	return fmt.Sprintf("(%v,%v)", c.x, c.y)
}

func coordToIdx(c coord, maxX int) (idx int) {
	idx = c.x + (c.y * maxX)
	return
}

func idxToCoord(idx int, maxX int) (c coord) {
	c = coord{x: idx % maxX, y: idx / maxX}
	return
}

func coordEqual(a coord, b coord) bool {
	if a.x == b.x && a.y == b.y {
		return true
	}
	return false
}

func coordAdd(a coord, b coord) (ret coord) {
	ret = a
	ret.x += b.x
	ret.y += b.y
	return
}

func coordSub(a coord, b coord) (ret coord) {
	ret = a
	ret.x -= b.x
	ret.y -= b.y
	return
}

func coordDiv(a coord, f int) (ret coord) {
	ret = a
	ret.x /= f
	ret.y /= f
	return
}

func coordMul(a coord, f int) (ret coord) {
	ret = a
	ret.x *= f
	ret.y *= f
	return
}

func drawAsteroidMap(asteroidMap []rune, w int) {
	for i := 0; i < len(asteroidMap); i++ {
		if i%w == 0 {
			fmt.Println()
		}
		fmt.Printf("%c", asteroidMap[i])
	}
	fmt.Printf("\n")
}
