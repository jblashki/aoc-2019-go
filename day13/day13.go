package day13

import (
	"bufio"
	"errors"
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

	bResult, err = b()
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

func pressEnter() {
	reader := bufio.NewReader(os.Stdin)
	//fmt.Print("<PRESS ENTER TO CONTINUE>")
	reader.ReadString('\n')

}

// func getChar() (ascii int, keyCode int, err error) {
// 	t, _ := term.Open("/dev/tty")
// 	term.RawMode(t)
// 	bytes := make([]byte, 3)

// 	var numRead int
// 	numRead, err = t.Read(bytes)
// 	if err != nil {
// 		return
// 	}
// 	if numRead == 3 && bytes[0] == 27 && bytes[1] == 91 {
// 		// Three-character control sequence, beginning with "ESC-[".

// 		// Since there are no ASCII codes for arrow keys, we use
// 		// Javascript key codes.
// 		if bytes[2] == 65 {
// 			// Up
// 			keyCode = 38
// 		} else if bytes[2] == 66 {
// 			// Down
// 			keyCode = 40
// 		} else if bytes[2] == 67 {
// 			// Right
// 			keyCode = 39
// 		} else if bytes[2] == 68 {
// 			// Left
// 			keyCode = 37
// 		}
// 	} else if numRead == 1 {
// 		ascii = int(bytes[0])
// 	} else {
// 		// Two characters read??
// 	}
// 	t.Restore()
// 	t.Close()
// 	return
// }

func a(verbose bool) (int, error) {
	var err error

	wg := new(sync.WaitGroup)
	ic, err := intcode.CreateLoad(wg, inputFile, 0, 0)
	if err != nil {
		return 0, err
	}

	intcode.Set(ic, 0, 2)

	defer intcode.Close(ic)

	wg.Add(1)
	go intcode.Run(ic, "")

	screen := make([][]tile, 0)
	score := 0
	//exit := false
	xPos := 0
	yPos := 0
	tileID := 0
	update := 0
	var sig intcode.Signal

	for {
		//control := 0
		exitLoop := false
		for !exitLoop {
			xPos, sig, err = intcode.Read(ic)
			switch sig {
			case intcode.SigNone:
				exitLoop = true
			case intcode.SigHalt:
				return 0, fmt.Errorf("HALT")

			case intcode.SigError:
				return 0, fmt.Errorf("Program error: %v", err)

			case intcode.SigInput:
				drawScreen(screen, score)
				fmt.Printf("? ")
				reader := bufio.NewReader(os.Stdin)
				char, _, err := reader.ReadRune()
				if err != nil {
					return 0, err
				}

				switch char {
				case 'a':
					intcode.Write(ic, -1)
				case 'd':
					intcode.Write(ic, 1)
				default:
					intcode.Write(ic, 0)
				}

				switch sig {
				// case intcode.SigNone:

				case intcode.SigHalt:
					return 0, fmt.Errorf("HALT")

				case intcode.SigError:
					return 0, fmt.Errorf("Program error: %v", err)
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
				return 0, fmt.Errorf("HALT")

			case intcode.SigError:
				return 0, fmt.Errorf("Program error: %v", err)

			case intcode.SigInput:
				drawScreen(screen, score)
				fmt.Printf("? ")
				reader := bufio.NewReader(os.Stdin)
				char, _, err := reader.ReadRune()
				if err != nil {
					return 0, err
				}

				switch char {
				case 'a':
					intcode.Write(ic, -1)
				case 'd':
					intcode.Write(ic, 1)
				default:
					intcode.Write(ic, 0)
				}

				switch sig {
				// case intcode.SigNone:

				case intcode.SigHalt:
					return 0, fmt.Errorf("HALT")

				case intcode.SigError:
					return 0, fmt.Errorf("Program error: %v", err)
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
				return 0, fmt.Errorf("HALT")

			case intcode.SigError:
				return 0, fmt.Errorf("Program error: %v", err)

			case intcode.SigInput:
				drawScreen(screen, score)
				fmt.Printf("? ")
				reader := bufio.NewReader(os.Stdin)
				char, _, err := reader.ReadRune()
				if err != nil {
					return 0, err
				}

				switch char {
				case 'a':
					intcode.Write(ic, -1)
				case 'd':
					intcode.Write(ic, 1)
				default:
					intcode.Write(ic, 0)
				}

				switch sig {
				// case intcode.SigNone:

				case intcode.SigHalt:
					return 0, fmt.Errorf("HALT")

				case intcode.SigError:
					return 0, fmt.Errorf("Program error: %v", err)
				}
			}
		}

		if xPos == -1 {
			score = tileID
		} else {
			t := tile(tileID)
			screen = printTile(screen, t, xPos, yPos)
		}

		update++
	}

	if verbose {
		drawScreen(screen, score)
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

func b() (int, error) {
	return 0, errors.New("Not Complete Yet")
}
