package day13

import (
	"fmt"
	"os"
	"os/exec"
	"sync"

	intcode "github.com/jblashki/aoc-intcode-go/v5"
)

const name = "Day 13"
const inputFile = "./day13/program"

type tile int

const (
	tileEmpty tile = iota
	tileWall
	tileBlock
	tilePaddle
	tileBall
)

func (t tile) String() string {
	return [...]string{" ", "#", "█", "-", "*"}[t]
}

// RunDay runs Advent of Code Day 13 Puzzle
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

func drawScreen(screen [][]tile, score int) {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
	for _, line := range screen {
		for _, tile := range line {
			fmt.Printf("%v", tile)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("Score: %v\n", score)
}

func printTile(screen [][]tile, t tile, xPos int, yPos int) [][]tile {
	rowLen := 0
	if len(screen) != 0 {
		rowLen = len(screen[0])
	} else {
		// add at least 1 row
		newRow := make([]tile, xPos+1)
		screen = append(screen, newRow)
		rowLen = xPos + 1
	}

	if xPos >= rowLen {
		//Expand all rows
		for i, row := range screen {
			newRow := make([]tile, xPos+1)
			copy(newRow, row)
			screen[i] = newRow
		}
	}

	if yPos >= len(screen) {
		for i := len(screen); i <= yPos; i++ {
			newRow := make([]tile, rowLen)
			screen = append(screen, newRow)
		}
	}

	screen[yPos][xPos] = t

	return screen
}

func findBallPaddleXPos(screen [][]tile) (ball int, paddle int) {
	for _, line := range screen {
		for xPos, tile := range line {
			if tile == tileBall {
				ball = xPos
			} else if tile == tilePaddle {
				paddle = xPos
			}
		}
	}

	return
}

func playGame(setQuarters bool, verbose bool) ([][]tile, int, error) {
	var err error

	wg := new(sync.WaitGroup)
	ic, err := intcode.CreateLoad(wg, inputFile, 0, 0)
	if err != nil {
		return nil, 0, err
	}

	if setQuarters {
		intcode.Set(ic, 0, 2)
	}

	defer intcode.Close(ic)

	wg.Add(1)
	go intcode.Run(ic, "")

	screen := make([][]tile, 0)
	score := 0
	xPos := 0
	yPos := 0
	tileID := 0
	var sig intcode.Signal

	for {
		exitLoop := false
		for !exitLoop {
			xPos, sig, err = intcode.Read(ic)
			switch sig {
			case intcode.SigNone:
				exitLoop = true
			case intcode.SigHalt:
				return screen, score, nil

			case intcode.SigError:
				return screen, score, fmt.Errorf("Program error: %v", err)

			case intcode.SigInput:
				ballPos, paddlePos := findBallPaddleXPos(screen)

				if ballPos < paddlePos {
					intcode.Write(ic, -1)
				} else if ballPos > paddlePos {
					intcode.Write(ic, 1)
				} else {
					intcode.Write(ic, 0)
				}
			}
		}

		exitLoop = false
		for !exitLoop {
			yPos, sig, err = intcode.Read(ic)
			switch sig {
			case intcode.SigNone:
				exitLoop = true

			case intcode.SigHalt:
				return screen, score, nil

			case intcode.SigError:
				return screen, score, fmt.Errorf("Program error: %v", err)

			case intcode.SigInput:
				ballPos, paddlePos := findBallPaddleXPos(screen)

				if ballPos < paddlePos {
					intcode.Write(ic, -1)
				} else if ballPos > paddlePos {
					intcode.Write(ic, 1)
				} else {
					intcode.Write(ic, 0)
				}
			}
		}

		exitLoop = false
		for !exitLoop {
			tileID, sig, err = intcode.Read(ic)
			switch sig {
			case intcode.SigNone:
				exitLoop = true

			case intcode.SigHalt:
				return screen, score, nil

			case intcode.SigError:
				return screen, score, fmt.Errorf("Program error: %v", err)

			case intcode.SigInput:
				ballPos, paddlePos := findBallPaddleXPos(screen)

				if ballPos < paddlePos {
					intcode.Write(ic, -1)
				} else if ballPos > paddlePos {
					intcode.Write(ic, 1)
				} else {
					intcode.Write(ic, 0)
				}
			}
		}

		if xPos == -1 {
			score = tileID
			if verbose {
				fmt.Printf("Score: %v\n", score)
			}
		} else {
			t := tile(tileID)
			screen = printTile(screen, t, xPos, yPos)
		}

		// update++
	}
}

func a(verbose bool) (int, error) {
	screen, _, err := playGame(false, verbose)
	if err != nil {
		return 0, err
	}

	blockCount := 0
	for _, line := range screen {
		for _, tile := range line {
			if tile == tileBlock {
				blockCount++
			}
		}
	}

	return blockCount, nil
}

func b(verbose bool) (int, error) {
	screen, score, err := playGame(true, verbose)
	if err != nil {
		return 0, err
	}

	blockCount := 0
	for _, line := range screen {
		for _, tile := range line {
			if tile == tileBlock {
				blockCount++
			}
		}
	}

	if blockCount != 0 {
		return 0, fmt.Errorf("Failed to finish game. %v blocks remaining", blockCount)
	}

	return score, nil
}
