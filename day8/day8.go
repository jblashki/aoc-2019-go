package day8

import (
	"fmt"

	"github.com/jblashki/aoc-filereader-go"
)

const name = "Day 8"
const input_file = "./day8/image"

const imageWidth = 25
const imageHeight = 6

func RunDay(verbose bool) {
	var aResult int
	var err error

	if verbose {
		fmt.Printf("\n%v Output:\n", name)
	}

	aResult, err = a()
	if err != nil {
		fmt.Printf("%va: **** Error: %q ****\n", name, err)
	} else {
		fmt.Printf("%va: Program Result = %v\n", name, aResult)
	}

	err = b()
	if err != nil {
		fmt.Printf("%vb: **** Error: %q ****\n", name, err)
	} else {
		fmt.Printf("%vb: See display above\n", name)
	}
}

func a() (int, error) {
	image, err := readImage(input_file, imageWidth, imageHeight)
	if err != nil {
		return 0, err
	}

	minZeros := -1
	output := -1
	for i := 0; i < len(image); i++ {
		count := make([]int, 10)
		for j := 0; j < len(image[i]); j++ {
			for k := 0; k < len(image[i][j]); k++ {
				count[image[i][j][k]] += 1
			}
			// End of Row
		}
		if minZeros == -1 || count[0] < minZeros {
			minZeros = count[0]
			output = count[1] * count[2]
		}
		// End of Layer
	}
	return output, nil
}

func b() error {
	image, err := readImage(input_file, imageWidth, imageHeight)
	if err != nil {
		return err
	}

	actualImage := make([][]int, 0)
	for j := 0; j < imageHeight; j++ {
		row := make([]int, 0)
		for k := 0; k < imageWidth; k++ {
			pixel := 2
			for i := 0; i < len(image); i++ {
				if image[i][j][k] != 2 {
					pixel = image[i][j][k]
					break
				}
			}
			row = append(row, pixel)
		}
		actualImage = append(actualImage, row)
		// End of Row
	}

	for i := 0; i < len(actualImage); i++ {
		for j := 0; j < len(actualImage[i]); j++ {
			if actualImage[i][j] == 1 {
				fmt.Printf("██")
			} else {
				fmt.Printf("  ")
			}
		}
		fmt.Println()
	}
	return nil
}

func readImage(file string, w int, h int) ([][][]int, error) {
	imageString, err := filereader.ReadIntoString(file)
	if err != nil {
		return nil, err
	}

	returnImage := make([][][]int, 0)

	layer := make([][]int, 0)
	row := make([]int, 0)
	rowCount := 0
	for i := 0; i < len(imageString); i++ {
		digit := int(imageString[i]) - int('0')
		row = append(row, digit)
		if i%w == w-1 {
			layer = append(layer, row)
			row = make([]int, 0)
			rowCount++
			if rowCount%h == 0 {
				returnImage = append(returnImage, layer)
				layer = make([][]int, 0)
			}
		}
	}

	return returnImage, nil
}
